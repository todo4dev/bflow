package activate_account

import (
	"context"
	"fmt"
	"strings"
	"time"

	"src/domain/entity"
	"src/domain/event"
	"src/domain/issue"
	"src/domain/repository"
	"src/port/broker"
	"src/port/cache"
	"src/port/logging"

	"github.com/google/uuid"
	"github.com/leandroluk/gox/meta"
	"github.com/leandroluk/gox/validate"
)

type Data struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

var dataSchema = validate.Object(func(t *Data, s *validate.ObjectSchema[Data]) {
	s.Field(&t.Email).Text().Required()
	s.Field(&t.OTP).Text().Required()
})

type Handler struct {
	repositoryAccount repository.Account
	loggingLogger     logging.Logger
	brokerPublisher   broker.Client
	cacheClient       cache.Client
}

func New(
	repositoryAccount repository.Account,
	cacheClient cache.Client,
	loggingLogger logging.Logger,
	brokerPublisher broker.Client,
) *Handler {
	return &Handler{
		repositoryAccount: repositoryAccount,
		cacheClient:       cacheClient,
		loggingLogger:     loggingLogger,
		brokerPublisher:   brokerPublisher,
	}
}

func (h *Handler) findOTP(ctx context.Context, otp string) (accountID uuid.UUID, key string, err error) {
	match := fmt.Sprintf("account:*:otp:%s:activate", otp)
	key, _, err = h.cacheClient.Get(ctx, match)
	if err != nil {
		return uuid.Nil, "", &issue.AccountInvalidOTP{}
	}
	accountID, err = uuid.Parse(strings.Split(key, ":")[1])
	if err != nil {
		return uuid.Nil, "", &issue.AccountInvalidOTP{}
	}

	return accountID, key, nil
}

func (h *Handler) findAccountById(accountID uuid.UUID) (*entity.Account, error) {
	account, err := h.repositoryAccount.FindById(accountID)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, &issue.AccountNotFound{}
	}
	if account.EmailVerifiedAt != nil {
		return nil, &issue.AccountAlreadyActivated{}
	}
	return account, nil
}

func (h *Handler) activateAccount(account *entity.Account) error {
	now := time.Now().UTC()
	account.TS = now
	account.EmailVerifiedAt = &now
	if err := h.repositoryAccount.Save(account); err != nil {
		return err
	}
	return nil
}

func (h *Handler) deleteKey(key string) {
	go func() {
		ctx := context.Background()
		if err := h.cacheClient.Delete(ctx, key); err != nil {
			h.loggingLogger.Error(ctx, "failed to delete key", err)
		}
	}()
}

func (h *Handler) publishAccountActivated(accountID uuid.UUID) {
	go func() {
		ctx := context.Background()
		if err := h.brokerPublisher.Publish(ctx, event.AccountActivated(accountID)); err != nil {
			h.loggingLogger.Error(ctx, "failed to publish message", err)
		}
	}()
}

func (h *Handler) Handle(ctx context.Context, data *Data) error {
	if _, err := dataSchema.Validate(*data); err != nil {
		return err
	}

	accountID, key, err := h.findOTP(ctx, data.OTP)
	if err != nil {
		return err
	}

	account, err := h.findAccountById(accountID)
	if err != nil {
		return err
	}

	if err := h.activateAccount(account); err != nil {
		return err
	}

	h.deleteKey(key)

	h.publishAccountActivated(accountID)

	return nil
}

func init() {
	data := Data{
		Email: "john.doe@email.com",
		OTP:   "123456"}
	meta.Describe(&data, meta.Description("Activate account data"),
		meta.Field(&data.Email, meta.Description("Email address")),
		meta.Field(&data.OTP, meta.Description("One-Time Password")),
		meta.Example(data))

	meta.Describe(&Handler{}, meta.Description("Activating account usecase"),
		meta.Throws[*issue.AccountInvalidOTP](),
		meta.Throws[*issue.AccountNotFound]())
}
