// application/auth/register_account/register_account.go
package register_account

import (
	"context"
	"fmt"
	"time"

	"src/domain"
	"src/domain/constant"
	"src/domain/entity"
	"src/domain/enum"
	"src/domain/event"
	"src/domain/issue"
	"src/port/broker"
	"src/port/cache"
	"src/port/logging"
	"src/port/mailing"

	"github.com/google/uuid"
	"github.com/leandroluk/gox/meta"
	"github.com/leandroluk/gox/util"
	"github.com/leandroluk/gox/validate"
	"golang.org/x/crypto/bcrypt"
)

type Data struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var dataSchema = validate.Object(func(t *Data, s *validate.ObjectSchema[Data]) {
	s.Field(&t.Email).Text().Required().Email()
	s.Field(&t.Password).Text().Required().Pattern(constant.ACCOUNT_CREDENTIAL_PASSWORD_REGEX)
})

type Handler struct {
	domainUow              domain.Uow
	mailingMailer          mailing.Mailer
	loggingLogger          logging.Logger
	brokerPublisherAccount broker.Client
	cacheClient            cache.Client
}

func New(
	domainUow domain.Uow,
	mailingMailer mailing.Mailer,
	loggingLogger logging.Logger,
	brokerPublisher broker.Client,
	cacheClient cache.Client,
) *Handler {
	return &Handler{
		domainUow:              domainUow,
		mailingMailer:          mailingMailer,
		loggingLogger:          loggingLogger,
		brokerPublisherAccount: brokerPublisher,
		cacheClient:            cacheClient,
	}
}

func (h *Handler) checkEmailInUse(email string) error {
	exists, err := h.domainUow.Account().ExistsByEmail(email)
	if err != nil {
		return err
	}
	if exists {
		return &issue.AccountEmailInUse{}
	}
	return nil
}

func (h *Handler) createAccount(email string, password string) (*entity.Account, error) {
	now := time.Now().UTC()
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	account := entity.Account{
		ID:              util.Must(uuid.NewV7()),
		TS:              now,
		CreatedAt:       now,
		DeletedAt:       nil,
		Email:           email,
		EmailVerifiedAt: nil,
		Role:            enum.AccountRole_MEMBER}
	credential := entity.AccountCredential{
		ID:           util.Must(uuid.NewV7()),
		TS:           now,
		CreatedAt:    now,
		PasswordHash: string(passwordHash),
		AccountID:    account.ID}
	profile := entity.AccountProfile{
		ID:         util.Must(uuid.NewV7()),
		TS:         now,
		GivenName:  nil,
		FamilyName: nil,
		Language:   constant.ACCOUNT_PROFILE_LANGUAGE_DEFAULT,
		Timezone:   constant.ACCOUNT_PROFILE_TIMEZONE_DEFAULT,
		AccountID:  account.ID}
	preference := entity.AccountPreference{
		ID:                      util.Must(uuid.NewV7()),
		TS:                      now,
		Theme:                   constant.ACCOUNT_PREFERENCE_THEME_DEFAULT,
		NotifyOnPipelineSuccess: true,
		NotifyOnInfraAlerts:     true,
		AccountID:               account.ID}

	err = h.domainUow.Do(func(t domain.Repository) error {
		if err := t.Account().Create(&account); err != nil {
			return err
		}
		if err := t.AccountCredential().Create(&credential); err != nil {
			return err
		}
		if err := t.AccountProfile().Create(&profile); err != nil {
			return err
		}
		if err := t.AccountPreference().Create(&preference); err != nil {
			return err
		}
		return nil
	})

	return &account, err
}

func (h *Handler) createOTP(ctx context.Context, accountID uuid.UUID) (string, error) {
	otp := uuid.New().String()[:6]
	key := fmt.Sprintf("account:%s:otp:%s", accountID.String(), otp)
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
			h.loggingLogger.Error(ctx, msg, err)
		}
	}()
}

func (h *Handler) publishAccountRegistered(email string) {
	go func() {
		ctx := context.Background()
		event := event.AccountRegistered(email)
		if err := h.brokerPublisherAccount.Publish(ctx, event); err != nil {
			msg := fmt.Sprintf("failed to publish AccountRegistered event to %s", email)
			h.loggingLogger.Error(ctx, msg, err)
		}
	}()
}

func (h *Handler) Handle(ctx context.Context, data *Data) error {
	if _, err := dataSchema.Validate(*data); err != nil {
		return err
	}

	if err := h.checkEmailInUse(data.Email); err != nil {
		return err
	}

	account, err := h.createAccount(data.Email, data.Password)
	if err != nil {
		return err
	}

	otp, err := h.createOTP(ctx, account.ID)
	if err != nil {
		return err
	}

	h.sendActivationEmail(otp, data.Email)

	h.publishAccountRegistered(data.Email)

	return nil
}

func init() {
	data := Data{
		Email:    "john.doe@email.com",
		Password: "Test@123"}
	meta.Describe(&data, meta.Description("Data for registering a new account"),
		meta.Field(&data.Email, meta.Description("Email address to create the account")),
		meta.Field(&data.Password, meta.Description("Account password")),
		meta.Example(data))

	meta.Describe(&Handler{}, meta.Description("Handler for registering a new account"),
		meta.Throws[*issue.AccountEmailInUse]())
}
