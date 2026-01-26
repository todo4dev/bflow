package reset_password

import (
	"context"
	"fmt"
	"time"

	"src/domain"
	"src/domain/constant"
	"src/domain/entity"
	"src/domain/event"
	"src/domain/issue"
	"src/domain/repository"
	"src/port/broker"
	"src/port/cache"
	"src/port/logger"

	"github.com/google/uuid"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/meta"
	"github.com/leandroluk/gox/validate"
	"golang.org/x/crypto/bcrypt"
)

var dataSchema = validate.Object(func(t *Data, s *validate.ObjectSchema[Data]) {
	s.Field(&t.OTP).Text().Required().Min(6).Max(6)
	s.Field(&t.Email).Text().Required().Email()
	s.Field(&t.NewPassword).Text().Required().Pattern(constant.ACCOUNT_CREDENTIAL_PASSWORD_REGEX)
})

type Data struct {
	OTP         string `json:"otp"`
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
}

func (d *Data) Validate() error {
	_, err := dataSchema.Validate(d)
	return err
}

type Handler struct {
	repositoryAccount           repository.Account
	repositoryAccountCredential repository.AccountCredential
	domainUow                   domain.Uow
	loggerClient                logger.Client
	brokerPublisher             broker.Client
	cacheClient                 cache.Client
}

func New(
	repositoryAccount repository.Account,
	repositoryAccountCredential repository.AccountCredential,
	loggerClient logger.Client,
	brokerPublisher broker.Client,
	cacheClient cache.Client,
) *Handler {
	return &Handler{
		repositoryAccount:           repositoryAccount,
		repositoryAccountCredential: repositoryAccountCredential,
		loggerClient:                loggerClient,
		brokerPublisher:             brokerPublisher,
		cacheClient:                 cacheClient,
	}
}

func (h *Handler) findActiveAccountByEmail(email string) (*entity.Account, error) {
	account, err := h.repositoryAccount.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, &issue.AccountNotFound{}
	}
	if account.EmailVerifiedAt == nil {
		return nil, &issue.AccountNotFound{}
	}
	if account.DeletedAt != nil {
		return nil, &issue.AccountNotFound{}
	}
	return account, nil
}

func (h *Handler) verifyOTP(ctx context.Context, accountID string, inputOTP string) error {
	match := fmt.Sprintf("account:%s:recover_otp", accountID)
	key, storedOTP, err := h.cacheClient.Get(ctx, match)
	if err != nil {
		return &issue.AccountInvalidOTP{}
	}
	if key == "" || storedOTP != inputOTP {
		return &issue.AccountInvalidOTP{}
	}
	_ = h.cacheClient.Delete(ctx, match)
	return nil
}

func (h *Handler) upsertPassword(accountID uuid.UUID, newPassword string) error {
	accountCredential, err := h.repositoryAccountCredential.FindByAccountId(accountID)
	if err != nil {
		return err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	if accountCredential != nil {
		accountCredential.TS = now
		accountCredential.PasswordHash = string(passwordHash)
		err = h.repositoryAccountCredential.Save(accountCredential)
	} else {
		accountCredential = &entity.AccountCredential{
			ID:           uuid.Must(uuid.NewV7()),
			TS:           now,
			CreatedAt:    now,
			PasswordHash: string(passwordHash),
			AccountID:    accountID,
		}
		err = h.repositoryAccountCredential.Create(accountCredential)
	}

	return err
}

func (h *Handler) publishAccountCredentialChanged(accountID uuid.UUID) {
	go func() {
		ctx := context.Background()
		event := event.AccountCredentialChanged(accountID)
		if err := h.brokerPublisher.Publish(ctx, event); err != nil {
			msg := fmt.Sprintf("failed to publish AccountCredentialChanged event to %s", accountID.String())
			h.loggerClient.Error(ctx, msg, err)
		}
	}()
}

func (h *Handler) Handle(ctx context.Context, data *Data) error {
	if err := data.Validate(); err != nil {
		return err
	}

	account, err := h.findActiveAccountByEmail(data.Email)
	if err != nil {
		return err
	}

	err = h.verifyOTP(ctx, account.ID.String(), data.OTP)
	if err != nil {
		return err
	}

	err = h.upsertPassword(account.ID, data.NewPassword)
	if err != nil {
		return err
	}

	h.publishAccountCredentialChanged(account.ID)

	return nil
}

func init() {
	data := Data{
		OTP:         "123456",
		Email:       "john.doe@email.com",
		NewPassword: "NewTest@123"}
	meta.Describe(&data, meta.Description("Data for resetting password"),
		meta.Field(&data.OTP, meta.Description("6-digit recovery code")),
		meta.Field(&data.Email, meta.Description("Account email address")),
		meta.Field(&data.NewPassword, meta.Description("New password complying with security rules")),
		meta.Example(data))

	meta.Describe(&Handler{}, meta.Description("Handler for resetting user password"),
		meta.Throws[*issue.AccountNotFound](),
		meta.Throws[*issue.AccountInvalidOTP]())
}

func Provide() {
	di.SingletonAs[*Handler](New)
}
