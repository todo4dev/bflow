// infrastructure/oidc/impl.go
package oidc

import (
	"errors"
	"src/infrastructure/oidc/http/adapter"
	"src/port/oidc"
)

type Provider struct {
	config    Config
	microsoft oidc.Adapter
	google    oidc.Adapter
}

var _ oidc.Provider = (*Provider)(nil)

func NewProvider(rawConfig Config) (*Provider, error) {
	config, err := ConfigSchema.Validate(rawConfig)
	if err != nil {
		return nil, err
	}

	f := &Provider{config: config}
	f.microsoft = adapter.NewHttpMicrosoft(
		config.BaseURI,
		config.MicrosoftClientID,
		config.MicrosoftClientSecret,
		config.MicrosoftCallbackURI,
	)
	f.google = adapter.NewHttpGoogle(
		config.BaseURI,
		config.GoogleClientID,
		config.GoogleClientSecret,
		config.GoogleCallbackURI,
	)
	return f, nil
}

func (f *Provider) GetAdapter(provider oidc.ProviderName) (oidc.Adapter, error) {
	switch provider {
	case oidc.ProviderName_MICROSOFT:
		return f.microsoft, nil
	case oidc.ProviderName_GOOGLE:
		return f.google, nil
	default:
		return nil, errors.New("unsupported provider")
	}
}
