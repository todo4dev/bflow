package send_reset_password

import (
	"context"
	"fmt"
	"time"

	"src/domain/entity"
	"src/domain/issue"
	"src/domain/repository"
	"src/port/cache"
	"src/port/logging"
	"src/port/mailing"

	"github.com/google/uuid"
	"github.com/leandroluk/gox/meta"
	"github.com/leandroluk/gox/validate"
)

type Data struct {
	Email string `json:"email"`
}

var dataSchema = validate.Object(func(t *Data, s *validate.ObjectSchema[Data]) {
	s.Field(&t.Email).Text().Required().Email()
})

type Handler struct {
	repositoryAccount repository.Account
	mailingMailer     mailing.Mailer
	loggingLogger     logging.Logger
	cacheClient       cache.Client
}

func New(
	repositoryAccount repository.Account,
	mailingMailer mailing.Mailer,
	loggingLogger logging.Logger,
	cacheClient cache.Client,
) *Handler {
	return &Handler{
		repositoryAccount: repositoryAccount,
		mailingMailer:     mailingMailer,
		loggingLogger:     loggingLogger,
		cacheClient:       cacheClient,
	}
}

func (h *Handler) findAccountByEmail(email string) (*entity.Account, error) {
	account, err := h.repositoryAccount.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, &issue.AccountNotFound{}
	}
	return account, nil
}

func (h *Handler) createOTP(ctx context.Context, accountID uuid.UUID) (string, error) {
	otp := uuid.New().String()[:6]
	key := fmt.Sprintf("account:%s:recover_otp", accountID.String())
	if err := h.cacheClient.Set(ctx, key, otp, 10*time.Minute); err != nil {
		return "", err
	}
	return otp, nil
}

func (h *Handler) sendRecoveryEmail(otp string, email string) {
	go func() {
		ctx := context.Background()
		err := h.mailingMailer.Send(ctx, mailing.Email{
			To:        []string{email},
			Subject:   "Reset your password",
			Template:  "recover.html",
			Variables: map[string]any{"otp": otp}})
		if err != nil {
			msg := fmt.Sprintf("failed to send recovery email to %s", email)
			h.loggingLogger.Error(ctx, msg, err)
		}
	}()
}

func (h *Handler) Handle(ctx context.Context, data *Data) error {
	if _, err := dataSchema.Validate(*data); err != nil {
		return err
	}

	account, err := h.findAccountByEmail(data.Email)
	if err != nil {
		return err
	}

	otp, err := h.createOTP(ctx, account.ID)
	if err != nil {
		return err
	}

	h.sendRecoveryEmail(otp, data.Email)

	return nil
}

func init() {
	data := Data{
		Email: "john.doe@email.com"}
	meta.Describe(&data, meta.Description("Data for sending reset password email"),
		meta.Field(&data.Email, meta.Description("Account email address")),
		meta.Example(data))

	meta.Describe(&Handler{}, meta.Description("Handler for sending reset password email"),
		meta.Throws[*issue.AccountNotFound]())
}
