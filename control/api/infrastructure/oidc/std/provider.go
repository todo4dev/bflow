// infrastructure/oidc/std/provider.go
package std

import (
	"errors"
	"src/infrastructure/oidc/std/adapter"
	"src/port/oidc"
)

type Provider struct {
	config    *Config
	microsoft oidc.Adapter
	google    oidc.Adapter
}

var _ oidc.Provider = (*Provider)(nil)

func New(config *Config) (*Provider, error) {
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
