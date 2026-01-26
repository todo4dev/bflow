package resend_activation_code

import (
	"context"
	"fmt"
	"time"

	"src/domain/entity"
	"src/domain/event"
	"src/domain/issue"
	"src/domain/repository"
	"src/port/broker"
	"src/port/cache"
	"src/port/logger"
	"src/port/mailing"

	"github.com/google/uuid"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/meta"
	"github.com/leandroluk/gox/validate"
)

var dataSchema = validate.Object(func(t *Data, s *validate.ObjectSchema[Data]) {
	s.Field(&t.Email).Text().Required().Email()
})

type Data struct {
	Email string `json:"email"`
}

func (d *Data) Validate() error {
	_, err := dataSchema.Validate(d)
	return err
}

type Handler struct {
	repositoryAccount repository.Account
	mailingMailer     mailing.Mailer
	loggerClient      logger.Client
	brokerPublisher   broker.Client
	cacheClient       cache.Client
}

func New(
	repositoryAccount repository.Account,
	mailingMailer mailing.Mailer,
	loggerClient logger.Client,
	brokerPublisher broker.Client,
	cacheClient cache.Client,
) *Handler {
	return &Handler{
		repositoryAccount: repositoryAccount,
		mailingMailer:     mailingMailer,
		loggerClient:      loggerClient,
		brokerPublisher:   brokerPublisher,
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
	if account.EmailVerifiedAt != nil {
		return nil, &issue.AccountAlreadyActivated{}
	}
	return account, nil
}

func (h *Handler) recreateOTP(ctx context.Context, accountID uuid.UUID) (string, error) {
	oldKey := fmt.Sprintf("account:%s:otp:*:activate", accountID.String())
	h.cacheClient.Delete(ctx, oldKey)
	otp := uuid.New().String()[:6]
	key := fmt.Sprintf("account:%s:otp:%s:activate", accountID.String(), otp)
	h.cacheClient.Set(ctx, key, "", 10*time.Minute)
	return otp, nil
}

func (h *Handler) sendActivationEmail(otp string, email string) {
	go func() {
		ctx := context.Background()
		err := h.mailingMailer.Send(ctx, mailing.Email{
			To:        []string{email},
			Subject:   "Activate your account",
			Template:  "activate.html",
			Variables: map[string]any{"otp": otp}})
		if err != nil {
			msg := fmt.Sprintf("failed to send activation email to %s", email)
			h.loggerClient.Error(ctx, msg, err)
		}
	}()
}

func (h *Handler) publishAccountActivated(accountID uuid.UUID) {
	go func() {
		event := event.AccountActivated(accountID)
		if err := h.brokerPublisher.Publish(context.Background(), event); err != nil {
			msg := fmt.Sprintf("failed to publish AccountActivated event to %s", accountID.String())
			h.loggerClient.Error(context.Background(), msg, err)
		}
	}()
}

func (h *Handler) Handle(ctx context.Context, data *Data) error {
	if err := data.Validate(); err != nil {
		return err
	}

	account, err := h.findAccountByEmail(data.Email)
	if err != nil {
		return err
	}

	otp, err := h.recreateOTP(ctx, account.ID)
	if err != nil {
		return err
	}

	h.sendActivationEmail(otp, data.Email)

	h.publishAccountActivated(account.ID)

	return nil
}

func init() {
	data := Data{
		Email: "john.doe@email.com"}
	meta.Describe(&data, meta.Description("Resend activation code Data"),
		meta.Field(&data.Email, meta.Description("Account email")),
		meta.Example(data))

	meta.Describe(&Handler{}, meta.Description("Handler for resending activation code"),
		meta.Throws[*issue.AccountNotFound](),
		meta.Throws[*issue.AccountAlreadyActivated]())
}

func Provide() {
	di.SingletonAs[*Handler](New)
}
