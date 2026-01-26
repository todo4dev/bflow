// application/auth/login_using_credential/login_using_credential.go
package login_using_credential

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"src/domain/entity"
	"src/domain/issue"
	"src/domain/repository"
	"src/port/cache"
	"src/port/jwt"
	"src/port/logger"
	"src/port/mailing"

	"github.com/google/uuid"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/meta"
	"github.com/leandroluk/gox/util"
	"github.com/leandroluk/gox/validate"
	"golang.org/x/crypto/bcrypt"
)

var dataSchema = validate.Object(func(t *Data, s *validate.ObjectSchema[Data]) {
	s.Field(&t.Email).Text().Required().Email()
	s.Field(&t.Password).Text().Required()
})

type Data struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (d *Data) Validate() error {
	_, err := dataSchema.Validate(d)
	return err
}

type Result struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type Handler struct {
	repositoryAccount           repository.Account
	repositoryAccountCredential repository.AccountCredential
	repositoryAccountProfile    repository.AccountProfile
	jwtProvider                 jwt.Provider
	cacheClient                 cache.Client
	mailingMailer               mailing.Mailer
	loggerClient                logger.Client
}

func New(
	repositoryAccount repository.Account,
	repositoryAccountCredential repository.AccountCredential,
	repositoryAccountProfile repository.AccountProfile,
	jwtProvider jwt.Provider,
	cacheClient cache.Client,
	mailingMailer mailing.Mailer,
	loggerClient logger.Client,
) *Handler {
	return &Handler{
		repositoryAccount:           repositoryAccount,
		repositoryAccountCredential: repositoryAccountCredential,
		repositoryAccountProfile:    repositoryAccountProfile,
		jwtProvider:                 jwtProvider,
		cacheClient:                 cacheClient,
		mailingMailer:               mailingMailer,
		loggerClient:                loggerClient,
	}
}

func (h *Handler) findAccountByEmail(email string) (*entity.Account, error) {
	account, err := h.repositoryAccount.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, &issue.AccountInvalidCredentials{}
	}
	if account.DeletedAt != nil {
		return nil, &issue.AccountDeactivated{}
	}
	return account, nil
}

func (h *Handler) findAccountCredentialByAccountId(accountId uuid.UUID) (*entity.AccountCredential, error) {
	accountCredential, err := h.repositoryAccountCredential.FindByAccountId(accountId)
	if err != nil {
		return nil, err
	}
	if accountCredential == nil {
		return nil, &issue.AccountInvalidCredentials{}
	}
	return accountCredential, nil
}

func (h *Handler) findAccountProfileByAccountId(accountId uuid.UUID) (*entity.AccountProfile, error) {
	accountProfile, err := h.repositoryAccountProfile.FindByAccountId(accountId)
	if err != nil {
		return nil, err
	}
	return accountProfile, nil
}

func (h *Handler) verifyPassword(passwordHash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return &issue.AccountInvalidCredentials{}
	}
	return nil
}

func (h *Handler) setActivationKey(ctx context.Context, account *entity.Account) (string, error) {
	otp := uuid.New().String()[:6]
	key := fmt.Sprintf("account:%s:otp:%s:activate", account.ID.String(), otp)

	err := h.cacheClient.Set(ctx, key, "", 15*time.Minute)
	if err != nil {
		return "", err
	}
	return otp, nil
}

func (h *Handler) sendActivationEmail(account *entity.Account, otp string) {
	go func() {
		ctx := context.Background()
		err := h.mailingMailer.Send(ctx, mailing.Email{
			To:        []string{account.Email},
			Subject:   "Activate your account",
			Template:  "activate.html",
			Variables: map[string]any{"otp": otp},
		})
		if err != nil {
			msg := fmt.Sprintf("failed to send activation email to %s", account.Email)
			h.loggerClient.Error(ctx, msg, err)
		}
	}()
}

func (h *Handler) createAccountSession(ctx context.Context, account *entity.Account) (string, error) {
	sessionID := uuid.New().String()
	key := fmt.Sprintf("account:%s:session:%s", account.ID.String(), sessionID)

	valueMap := map[string]any{"maxAge": time.Now().Add(h.jwtProvider.GetRefreshTokenTTL())}
	value := string(util.Must(json.Marshal(valueMap)))

	if err := h.cacheClient.Set(ctx, key, value, h.jwtProvider.GetAccessTokenTTL()); err != nil {
		return "", err
	}

	return sessionID, nil
}

func (h *Handler) createJwt(sessionID string, account *entity.Account, profile *entity.AccountProfile) (*jwt.Token, error) {
	claims := jwt.Claims{Subject: account.ID.String(), Email: account.Email}

	if profile != nil {
		claims.GivenName = profile.GivenName
		claims.FamilyName = profile.FamilyName
		claims.Language = util.Ptr(profile.Language)
		claims.Timezone = util.Ptr(profile.Timezone)
	}

	token, err := h.jwtProvider.Create(sessionID, claims, true)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (h *Handler) Handle(ctx context.Context, data *Data) (*Result, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	account, err := h.findAccountByEmail(data.Email)
	if err != nil {
		return nil, err
	}

	accountCredential, err := h.findAccountCredentialByAccountId(account.ID)
	if err != nil {
		return nil, err
	}

	if err := h.verifyPassword(accountCredential.PasswordHash, data.Password); err != nil {
		return nil, err
	}

	if account.EmailVerifiedAt == nil {
		otp, err := h.setActivationKey(ctx, account)
		if err != nil {
			return nil, err
		}
		h.sendActivationEmail(account, otp)
		return nil, &issue.AccountNotVerified{}
	}

	accountProfile, err := h.findAccountProfileByAccountId(account.ID)
	if err != nil {
		return nil, err
	}

	sessionID, err := h.createAccountSession(ctx, account)
	if err != nil {
		return nil, err
	}

	token, err := h.createJwt(sessionID, account, accountProfile)
	if err != nil {
		return nil, err
	}

	return &Result{
		TokenType:    token.TokenType,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresIn:    token.ExpiresIn,
	}, nil
}

func init() {
	data := Data{
		Email:    "john.doe@email.com",
		Password: "Test@123"}
	meta.Describe(&data, meta.Description("Login Data"),
		meta.Field(&data.Email, meta.Description("User email")),
		meta.Field(&data.Password, meta.Description("User password")),
		meta.Example(data))

	meta.Describe(&Handler{}, meta.Description("Handler for login using credential"),
		meta.Throws[*issue.AccountInvalidCredentials](),
		meta.Throws[*issue.AccountNotVerified](),
		meta.Throws[*issue.AccountDeactivated]())

	result := Result{
		TokenType:    "Bearer",
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
		ExpiresIn:    3600}
	meta.Describe(&result, meta.Description("Login Result"),
		meta.Field(&result.TokenType, meta.Description("Token type")),
		meta.Field(&result.AccessToken, meta.Description("Access token")),
		meta.Field(&result.RefreshToken, meta.Description("Refresh token")),
		meta.Field(&result.ExpiresIn, meta.Description("Expires in")),
		meta.Example(result))
}

func Provide() {
	di.SingletonAs[*Handler](New)
}
