package sso_provider_redirect

import (
	"context"

	"src/port/oidc"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/validate"
)

var dataSchema = validate.Object(func(t *Data, s *validate.ObjectSchema[Data]) {
	s.Field(&t.Provider).Text().Required().OneOf("google", "microsoft")
	s.Field(&t.CallbackURL).Text().Required().URI()
})

type Data struct {
	Provider    string `json:"provider"`
	CallbackURL string `json:"callback_url"`
}

func (d *Data) Validate() error {
	_, err := dataSchema.Validate(d)
	return err
}

type Result struct {
	RedirectURL string `json:"redirect_url"`
}

type Handler struct {
	oidcProvider oidc.Provider
}

func New(
	oidcProvider oidc.Provider,
) *Handler {
	return &Handler{
		oidcProvider: oidcProvider,
	}
}

func (h *Handler) Handle(ctx context.Context, data *Data) (*Result, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	adapter, err := h.oidcProvider.GetAdapter(oidc.ProviderName(data.Provider))
	if err != nil {
		return nil, err
	}

	redirectURL := adapter.GetAuthURL(map[string]any{"callback_url": data.CallbackURL})

	return &Result{RedirectURL: redirectURL}, nil
}

func Provide() {
	di.SingletonAs[*Handler](New)
}
