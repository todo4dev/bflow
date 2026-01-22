package resend_activation_code

import (
	"context"
	"fmt"
	"time"

	"src/domain"
	"src/domain/entity"
	"src/domain/event"
	"src/domain/issue"
	"src/port/broker"
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
	domainUow              domain.Uow
	mailingMailer          mailing.Mailer
	loggingLogger          logging.Logger
	brokerPublisherAccount broker.Publisher[event.Account]
	cacheClient            cache.Client
}

func New(
	domainUow domain.Uow,
	mailingMailer mailing.Mailer,
	loggingLogger logging.Logger,
	brokerPublisherAccount broker.Publisher[event.Account],
	cacheClient cache.Client,
) *Handler {
	return &Handler{
		domainUow:              domainUow,
		mailingMailer:          mailingMailer,
		loggingLogger:          loggingLogger,
		brokerPublisherAccount: brokerPublisherAccount,
		cacheClient:            cacheClient,
	}
}

func (h *Handler) findAccountByEmail(email string) (*entity.Account, error) {
	account, err := h.domainUow.Account().FindByEmail(email)
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

func (h *Handler) createOTP(ctx context.Context, accountID uuid.UUID) (string, error) {
	otp := uuid.New().String()[:6]
	key := fmt.Sprintf("account:%s:otp:%s", accountID.String(), otp)
	h.cacheClient.Set(ctx, key, "", 10*time.Minute)
	return otp, nil
}

func (h *Handler) sendActivationEmail(ctx context.Context, otp string, email string) error {
	return h.mailingMailer.Send(ctx, mailing.Email{
		To:        []string{email},
		Subject:   "Activate your account",
		Template:  "activate.html",
		Variables: map[string]any{"otp": otp},
	})
}

func (h *Handler) Handle(ctx context.Context, data *Data) (any, error) {
	if _, err := dataSchema.Validate(*data); err != nil {
		return nil, err
	}

	account, err := h.findAccountByEmail(data.Email)
	if err != nil {
		return nil, err
	}

	otp, err := h.createOTP(ctx, account.ID)
	if err != nil {
		return nil, err
	}

	if err := h.sendActivationEmail(ctx, otp, data.Email); err != nil {
		h.loggingLogger.Error(ctx, "failed to send email", err)
	}

	return nil, nil
}

func init() {
	data := Data{
		Email: "john.doe@email.com",
	}
	meta.Describe(&data, meta.Description("Resend activation code Data"),
		meta.Field(&data.Email, meta.Description("Account email")),
		meta.Example(data))

	meta.Describe(&Handler{}, meta.Description("Handler for resending activation code"),
		meta.Throws[*issue.AccountNotFound](),
		meta.Throws[*issue.AccountAlreadyActivated]())
}
