// application/auth/login_using_code/login_using_code.go
package login_using_code

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"src/domain/entity"
	"src/domain/issue"
	"src/domain/repository"
	"src/port/cache"
	"src/port/crypto"
	"src/port/jwt"
	"src/port/logger"

	"github.com/google/uuid"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/meta"
	"github.com/leandroluk/gox/util"
	"github.com/leandroluk/gox/validate"
)

var dataSchema = validate.Object(func(t *Data, s *validate.ObjectSchema[Data]) {
	s.Field(&t.Code).Text().Required()
})

type Data struct {
	Code string `json:"code"`
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
	repositoryAccount        repository.Account
	repositoryAccountProfile repository.AccountProfile
	cryptoClient             crypto.Client
	jwtProvider              jwt.Provider
	cacheClient              cache.Client
	loggerClient             logger.Client
}

func New(
	repositoryAccount repository.Account,
	repositoryAccountProfile repository.AccountProfile,
	cryptoClient crypto.Client,
	jwtProvider jwt.Provider,
	cacheClient cache.Client,
	loggerClient logger.Client,
) *Handler {
	return &Handler{
		repositoryAccount:        repositoryAccount,
		repositoryAccountProfile: repositoryAccountProfile,
		cryptoClient:             cryptoClient,
		jwtProvider:              jwtProvider,
		cacheClient:              cacheClient,
		loggerClient:             loggerClient,
	}
}

type codePayload struct {
	AccountID string `json:"account_id"`
	TTL       string `json:"ttl"`
}

func (h *Handler) validateCode(code string) (*codePayload, error) {
	decrypted, err := h.cryptoClient.Decrypt(code)
	if err != nil {
		return nil, &issue.AccountInvalidCredentials{}
	}

	var payload codePayload
	if err := json.Unmarshal([]byte(decrypted), &payload); err != nil {
		return nil, &issue.AccountInvalidCredentials{}
	}

	ttl, err := time.Parse(time.RFC3339, payload.TTL)
	if err != nil {
		return nil, &issue.AccountInvalidCredentials{}
	}

	if time.Now().UTC().After(ttl) {
		return nil, &issue.AccountInvalidCredentials{}
	}

	return &payload, nil
}

func (h *Handler) findAccount(accountID string) (*entity.Account, error) {
	id, err := uuid.Parse(accountID)
	if err != nil {
		return nil, &issue.AccountInvalidCredentials{}
	}

	account, err := h.repositoryAccount.FindById(id)
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

func (h *Handler) findAccountProfile(accountID uuid.UUID) (*entity.AccountProfile, error) {
	profile, err := h.repositoryAccountProfile.FindByAccountId(accountID)
	if err != nil {
		return nil, err
	}
	return profile, nil
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

	payload, err := h.validateCode(data.Code)
	if err != nil {
		return nil, err
	}

	account, err := h.findAccount(payload.AccountID)
	if err != nil {
		return nil, err
	}

	profile, err := h.findAccountProfile(account.ID)
	if err != nil {
		return nil, err
	}

	sessionID, err := h.createAccountSession(ctx, account)
	if err != nil {
		return nil, err
	}

	token, err := h.createJwt(sessionID, account, profile)
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
	data := Data{Code: "encrypted_code"}
	meta.Describe(&data, meta.Description("Login via Code Data"),
		meta.Field(&data.Code, meta.Description("Encrypted Code from SSO Callback")),
		meta.Example(data))

	meta.Describe(&Handler{}, meta.Description("Handler for login using internal code"),
		meta.Throws[*issue.AccountInvalidCredentials](),
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
