package sso_provider_callback

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"time"

	"src/domain"
	"src/domain/constant"
	"src/domain/entity"
	"src/domain/enum"
	"src/domain/event"
	"src/domain/issue"
	"src/port/broker"
	"src/port/crypto"
	"src/port/image"
	"src/port/logger"
	"src/port/oidc"
	"src/port/storage"

	"github.com/google/uuid"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/meta"
	"github.com/leandroluk/gox/util"
	"github.com/leandroluk/gox/validate"
)

var dataSchema = validate.Object(func(t *Data, s *validate.ObjectSchema[Data]) {
	s.Field(&t.Provider).Text().Required().OneOf("google", "microsoft")
	s.Field(&t.Code).Text()
	s.Field(&t.State).Text().Required()
	s.Field(&t.Error).Text()
})

type Data struct {
	Provider string  `json:"provider"`
	Code     string  `json:"code"`
	State    string  `json:"state"`
	Error    *string `json:"error"`
}

func (d *Data) Validate() error {
	_, err := dataSchema.Validate(d)
	return err
}

type Result struct {
	RedirectURL string `json:"redirect_url"`
}

type Handler struct {
	domainUow              domain.Uow
	oidcProvider           oidc.Provider
	cryptoClient           crypto.Client
	storageClient          storage.Client
	brokerPublisherAccount broker.Client
	loggerClient           logger.Client
	imageProcessor         image.Processor
}

func New(
	domainUow domain.Uow,
	oidcProvider oidc.Provider,
	cryptoClient crypto.Client,
	storageClient storage.Client,
	brokerPublisher broker.Client,
	loggerClient logger.Client,
	imageProcessor image.Processor,
) *Handler {
	return &Handler{
		domainUow:              domainUow,
		oidcProvider:           oidcProvider,
		cryptoClient:           cryptoClient,
		storageClient:          storageClient,
		brokerPublisherAccount: brokerPublisher,
		loggerClient:           loggerClient,
		imageProcessor:         imageProcessor,
	}
}

func (h *Handler) getCallbackURL(adapter oidc.Adapter, state string) (string, error) {
	stateMap, err := adapter.DecodeState(state)
	if err == nil {
		callbackURL, ok := stateMap["callback_url"].(string)
		if ok && callbackURL != "" {
			return callbackURL, nil
		}
	}
	return "", &issue.AccountInvalidSSOCallback{}
}

func (h *Handler) setQuery(callbackURL string, key, value string) string {
	u, _ := url.Parse(callbackURL)
	q := u.Query()
	q.Set(key, value)
	u.RawQuery = q.Encode()
	return u.String()
}

func (h *Handler) processProfilePicture(ctx context.Context, accountID uuid.UUID, adapter oidc.Adapter, accessToken string) {
	path := fmt.Sprintf("account/%s/picture.png", accountID.String())

	exists, err := h.storageClient.Exists(ctx, path)
	if err == nil && exists {
		return // Already has picture
	}

	// Download from provider
	reader, err := adapter.GetPicture(ctx, accessToken)
	if err != nil {
		h.loggerClient.Warn(ctx, fmt.Sprintf("failed to get picture for account %s", accountID), logger.Error(err))
		return
	}
	defer reader.Close()

	// Resize using Port
	resized, err := h.imageProcessor.Resize(reader, 192, 192, image.PNG)
	if err != nil {
		h.loggerClient.Warn(ctx, fmt.Sprintf("failed to resize picture for account %s", accountID), logger.Error(err))
		return
	}

	// Prepare buffer for upload
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resized); err != nil {
		h.loggerClient.Warn(ctx, fmt.Sprintf("failed to read resized picture for account %s", accountID), logger.Error(err))
		return
	}

	// Upload
	if err := h.storageClient.Upload(ctx, path, &buf, map[string]string{"Content-Type": "image/png"}); err != nil {
		h.loggerClient.Warn(ctx, fmt.Sprintf("failed to upload picture for account %s", accountID), logger.Error(err))
	}
}

func (h *Handler) createOrUpdateAccount(claims *oidc.Claims) (*entity.Account, bool, error) {
	now := time.Now().UTC()
	var account *entity.Account
	var created bool

	err := h.domainUow.Do(func(repo domain.Repository) error {
		existing, err := repo.Account().FindByEmail(claims.Email)
		if err == nil && existing != nil {
			// Update existing
			account = existing
			profile, err := repo.AccountProfile().FindByAccountId(existing.ID)
			if err != nil {
				return err // Should exist
			}

			updated := false
			if profile.GivenName == nil && claims.Custom["given_name"] != nil {
				gn := fmt.Sprintf("%v", claims.Custom["given_name"])
				profile.GivenName = &gn
				updated = true
			}
			if profile.FamilyName == nil && claims.Custom["family_name"] != nil {
				fn := fmt.Sprintf("%v", claims.Custom["family_name"])
				profile.FamilyName = &fn
				updated = true
			}

			if updated {
				return repo.AccountProfile().Save(profile)
			}
			return nil
		}

		// Create new
		newAccount := entity.Account{
			ID:              util.Must(uuid.NewV7()),
			TS:              now,
			CreatedAt:       now,
			DeletedAt:       nil,
			Email:           claims.Email,
			EmailVerifiedAt: &now, // Verified by Provider
			Role:            enum.AccountRole_MEMBER,
		}

		credential := entity.AccountCredential{
			ID:           util.Must(uuid.NewV7()),
			TS:           now,
			CreatedAt:    now,
			PasswordHash: "", // Empty for SSO users
			AccountID:    newAccount.ID,
		}

		// Prepare names
		var givenName *string
		var familyName *string
		if gn, ok := claims.Custom["given_name"].(string); ok {
			givenName = &gn
		}
		if fn, ok := claims.Custom["family_name"].(string); ok {
			familyName = &fn
		}

		profile := entity.AccountProfile{
			ID:         util.Must(uuid.NewV7()),
			TS:         now,
			GivenName:  givenName,
			FamilyName: familyName,
			Language:   constant.ACCOUNT_PROFILE_LANGUAGE_DEFAULT,
			Timezone:   constant.ACCOUNT_PROFILE_TIMEZONE_DEFAULT,
			AccountID:  newAccount.ID,
		}

		preference := entity.AccountPreference{
			ID:                      util.Must(uuid.NewV7()),
			TS:                      now,
			Theme:                   constant.ACCOUNT_PREFERENCE_THEME_DEFAULT,
			NotifyOnPipelineSuccess: true,
			NotifyOnInfraAlerts:     true,
			AccountID:               newAccount.ID,
		}

		// Execute Creations
		if err := repo.Account().Create(&newAccount); err != nil {
			return err
		}
		if err := repo.AccountCredential().Create(&credential); err != nil {
			return err
		}
		if err := repo.AccountProfile().Create(&profile); err != nil {
			return err
		}
		if err := repo.AccountPreference().Create(&preference); err != nil {
			return err
		}

		account = &newAccount
		created = true
		return nil
	})

	return account, created, err
}

func (h *Handler) Handle(ctx context.Context, data *Data) (*Result, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	adapter := util.Must(h.oidcProvider.GetAdapter(oidc.ProviderName(data.Provider)))

	callbackURL, err := h.getCallbackURL(adapter, data.State)
	if err != nil {
		h.loggerClient.Error(ctx, "failed to get callback url", err)
		return nil, err
	}

	if data.Error != nil {
		return &Result{RedirectURL: h.setQuery(callbackURL, "error", *data.Error)}, nil
	}

	tokens, err := adapter.Exchange(ctx, data.Code)
	if err != nil {
		h.loggerClient.Error(ctx, "failed to exchange code", err)
		return &Result{RedirectURL: h.setQuery(callbackURL, "error", "exchange_failed")}, nil
	}

	claims, err := adapter.GetInfo(ctx, tokens.AccessToken)
	if err != nil {
		h.loggerClient.Error(ctx, "failed to get user info", err)
		return &Result{RedirectURL: h.setQuery(callbackURL, "error", "user_info_failed")}, nil
	}

	account, created, err := h.createOrUpdateAccount(claims)
	if err != nil {
		h.loggerClient.Error(ctx, "failed to create/update account", err)
		return &Result{RedirectURL: h.setQuery(callbackURL, "error", "account_error")}, nil
	}

	if created {
		if err := h.brokerPublisherAccount.Publish(ctx, event.AccountRegistered(account.Email)); err != nil {
			h.loggerClient.Warn(ctx, "failed to publish account registered event", logger.Error(err))
		}
	}

	h.processProfilePicture(ctx, account.ID, adapter, tokens.AccessToken)

	codePayload := map[string]any{
		"account_id": account.ID.String(),
		"ttl":        time.Now().Add(1 * time.Minute).Format(time.RFC3339),
	}
	codeJson, _ := json.Marshal(codePayload)
	encryptedCode, err := h.cryptoClient.Encrypt(string(codeJson))
	if err != nil {
		h.loggerClient.Error(ctx, "failed to encrypt login code", err)
		return &Result{RedirectURL: h.setQuery(callbackURL, "error", "server_error")}, nil
	}

	return &Result{RedirectURL: h.setQuery(callbackURL, "code", encryptedCode)}, nil
}

func init() {
	data := Data{}
	meta.Describe(&data, meta.Description("Data for SSO callback"),
		meta.Field(&data.Provider, meta.Description("Identity Provider (google/microsoft)")),
		meta.Field(&data.Code, meta.Description("Auth code from provider")),
		meta.Field(&data.State, meta.Description("State containing callback URL")),
		meta.Field(&data.Error, meta.Description("Provider error code")))

	meta.Describe(&Handler{}, meta.Description("Handler for SSO callback"),
		meta.Throws[*issue.AccountInvalidSSOCallback]())

	result := Result{
		RedirectURL: "http://localhost:3000?code=code&error=error"}
	meta.Describe(&result, meta.Description("Result of SSO callback"),
		meta.Field(&result.RedirectURL, meta.Description("Redirect URL")))
}

func Provide() {
	di.SingletonAs[*Handler](New)
}
