// application/auth/refresh_token/refresh_token.go
package refresh_token

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

	"github.com/google/uuid"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/meta"
	"github.com/leandroluk/gox/util"
	"github.com/leandroluk/gox/validate"
)

var dataSchema = validate.Object(func(t *Data, s *validate.ObjectSchema[Data]) {
	s.Field(&t.RefreshToken).Text().Required()
})

type Data struct {
	RefreshToken string `json:"refresh_token"`
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
	jwtProvider              jwt.Provider
	cacheClient              cache.Client
	loggerClient             logger.Client
}

func New(
	repositoryAccount repository.Account,
	repositoryAccountProfile repository.AccountProfile,
	jwtProvider jwt.Provider,
	cacheClient cache.Client,
	loggerClient logger.Client,
) *Handler {
	return &Handler{
		repositoryAccount:        repositoryAccount,
		repositoryAccountProfile: repositoryAccountProfile,
		jwtProvider:              jwtProvider,
		cacheClient:              cacheClient,
		loggerClient:             loggerClient,
	}
}

func (h *Handler) decodeRefreshToken(refreshToken string) (*jwt.Decoded, error) {
	decoded, err := h.jwtProvider.Decode(refreshToken)
	if err != nil || decoded.Kind != jwt.Kind_REFRESH {
		return nil, &issue.AccountInvalidToken{}
	}
	return decoded, nil
}

func (h *Handler) getSession(ctx context.Context, sessionID string) (*string, map[string]any, error) {
	genericSessionKey := fmt.Sprintf("account:*:session:%s", sessionID)
	key, value, err := h.cacheClient.Get(ctx, genericSessionKey)
	if err != nil || value == "" {
		return nil, nil, &issue.AccountSessionExpired{}
	}

	var session map[string]any
	if err := json.Unmarshal([]byte(value), &session); err != nil {
		return nil, nil, &issue.AccountSessionExpired{}
	}

	return &key, session, nil
}

func (h *Handler) verifySessionExpiration(ctx context.Context, key string, value map[string]any) error {
	maxAgeStr, ok := value["maxAge"]
	if !ok {
		return &issue.AccountSessionExpired{}
	}

	maxAge, err := time.Parse(time.RFC3339, maxAgeStr.(string))
	if err != nil {
		return &issue.AccountSessionExpired{}
	}

	if time.Now().After(maxAge) {
		return &issue.AccountSessionExpired{}
	}

	valueStr := string(util.Must(json.Marshal(value)))
	if err := h.cacheClient.Set(ctx, key, valueStr, h.jwtProvider.GetAccessTokenTTL()); err != nil {
		return err
	}

	return nil
}

func (h *Handler) deleteAccountSession(key string) {
	go func() {
		ctx := context.Background()
		if err := h.cacheClient.Delete(ctx, key); err != nil {
			h.loggerClient.Error(ctx, "Failed to delete expired session", err)
		}
	}()
}

func (h *Handler) createJwt(sessionID string, account *entity.Account, accountProfile *entity.AccountProfile) (*jwt.Token, error) {
	claims := jwt.Claims{
		Subject:    account.ID.String(),
		Email:      account.Email,
		GivenName:  accountProfile.GivenName,
		FamilyName: accountProfile.FamilyName,
		Language:   util.Ptr(accountProfile.Language),
		Timezone:   util.Ptr(accountProfile.Timezone),
	}

	newToken, err := h.jwtProvider.Create(sessionID, claims, true)
	if err != nil {
		return nil, err
	}
	return newToken, nil
}

func (h *Handler) getAccountWithProfile(accountID uuid.UUID) (*entity.Account, *entity.AccountProfile, error) {
	account, err := h.repositoryAccount.FindById(accountID)
	if err != nil {
		return nil, nil, err
	}
	if account == nil {
		return nil, nil, &issue.AccountInvalidToken{}
	}
	if account.DeletedAt != nil {
		return nil, nil, &issue.AccountDeactivated{}
	}

	accountProfile, err := h.repositoryAccountProfile.FindByAccountId(account.ID)
	if err != nil {
		return nil, nil, err
	}

	return account, accountProfile, nil
}

func (h *Handler) Handle(ctx context.Context, data *Data) (*Result, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	decoded, err := h.decodeRefreshToken(data.RefreshToken)
	if err != nil {
		return nil, err
	}

	sessionKey, sessionValue, err := h.getSession(ctx, decoded.SessionID)
	if err != nil {
		return nil, err
	}

	if err := h.verifySessionExpiration(ctx, *sessionKey, sessionValue); err != nil {
		h.deleteAccountSession(*sessionKey)
		return nil, err
	}

	account, accountProfile, err := h.getAccountWithProfile(uuid.Must(uuid.Parse(decoded.Claims.Subject)))
	if err != nil {
		return nil, err
	}

	newToken, err := h.createJwt(decoded.SessionID, account, accountProfile)
	if err != nil {
		return nil, err
	}

	return &Result{
		TokenType:    newToken.TokenType,
		AccessToken:  newToken.AccessToken,
		RefreshToken: newToken.RefreshToken,
		ExpiresIn:    newToken.ExpiresIn,
	}, nil
}

func init() {
	data := Data{
		RefreshToken: "header.payload.signature"}
	meta.Describe(&data, meta.Description("Refresh Token Data"),
		meta.Field(&data.RefreshToken, meta.Description("Refresh token")),
		meta.Example(data))

	meta.Describe(&Handler{}, meta.Description("Handler for refresh authorization token"),
		meta.Throws[*issue.AccountInvalidToken](),
		meta.Throws[*issue.AccountSessionExpired](),
		meta.Throws[*issue.AccountDeactivated]())

	result := Result{
		TokenType:    "Bearer",
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
		ExpiresIn:    3600}
	meta.Describe(&result, meta.Description("Refresh Token Result"),
		meta.Field(&result.TokenType, meta.Description("Token type")),
		meta.Field(&result.AccessToken, meta.Description("Access token")),
		meta.Field(&result.RefreshToken, meta.Description("Refresh token")),
		meta.Field(&result.ExpiresIn, meta.Description("Expires in")),
		meta.Example(result))
}

func Provide() {
	di.SingletonAs[*Handler](New)
}
