package activate_account

import (
	"context"
	"fmt"
	"strings"
	"time"

	"src/domain"
	"src/domain/entity"
	"src/domain/event"
	"src/domain/issue"
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
	domainUow              domain.Uow
	loggingLogger          logging.Logger
	brokerPublisherAccount broker.Publisher[event.Account]
	cacheClient            cache.Client
}

func New(
	domainUow domain.Uow,
	cacheClient cache.Client,
	loggingLogger logging.Logger,
	brokerPublisherAccount broker.Publisher[event.Account],
) *Handler {
	return &Handler{
		domainUow:              domainUow,
		cacheClient:            cacheClient,
		loggingLogger:          loggingLogger,
		brokerPublisherAccount: brokerPublisherAccount,
	}
}

func (h *Handler) findCachedKeyByOTP(ctx context.Context, otp string) (accountID uuid.UUID, key string, err error) {
	match := fmt.Sprintf("account:*:otp:%s", otp)
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
	account, err := h.domainUow.Account().FindById(accountID)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, &issue.AccountNotFound{}
	}
	return account, nil
}

func (h *Handler) activateAccount(account *entity.Account) error {
	now := time.Now().UTC()
	account.TS = now
	account.EmailVerifiedAt = &now
	if err := h.domainUow.Account().Save(account); err != nil {
		return err
	}
	return nil
}

func (h *Handler) Handle(ctx context.Context, data *Data) (any, error) {
	if _, err := dataSchema.Validate(*data); err != nil {
		return nil, err
	}

	accountID, key, err := h.findCachedKeyByOTP(ctx, data.OTP)
	if err != nil {
		return nil, err
	}

	account, err := h.findAccountById(accountID)
	if err != nil {
		return nil, err
	}

	if err := h.activateAccount(account); err != nil {
		return nil, err
	}

	if err := h.cacheClient.Delete(ctx, key); err != nil {
		h.loggingLogger.Error(ctx, "failed to delete key", err)
	}

	if err := h.brokerPublisherAccount.Publish(ctx, event.AccountActivated(accountID)); err != nil {
		h.loggingLogger.Error(ctx, "failed to publish message", err)
	}

	return nil, nil
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
		meta.Throws[issue.AccountInvalidOTP](issue.AccountInvalidOTP_MESSAGE),
		meta.Throws[issue.AccountNotFound](issue.AccountNotFound_MESSAGE))
}
