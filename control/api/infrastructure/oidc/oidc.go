// infrastructure/oidc/oidc.go
package oidc

import (
	"fmt"
	"src/infrastructure/oidc/std_http"
	"src/port/oidc"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

func Provide() {
	provider := env.Get("OIDC_PROVIDER", "http")
	switch provider {
	case "http":
		config := std_http.Config{
			BaseURI:               env.Get("OIDC_BASE_URI", "http://localhost:3000"),
			MicrosoftClientID:     env.Get("OIDC_MICROSOFT_CLIENT_ID", ""),
			MicrosoftClientSecret: env.Get("OIDC_MICROSOFT_CLIENT_SECRET", ""),
			MicrosoftCallbackURI:  env.Get("OIDC_MICROSOFT_CALLBACK", "/auth/microsoft/callback"),
			GoogleClientID:        env.Get("OIDC_GOOGLE_CLIENT_ID", ""),
			GoogleClientSecret:    env.Get("OIDC_GOOGLE_CLIENT_SECRET", ""),
			GoogleCallbackURI:     env.Get("OIDC_GOOGLE_CALLBACK", "/auth/google/callback"),
		}
		if err := config.Validate(); err != nil {
			panic(fmt.Errorf("oidc config validation failed: %w", err))
		}

		instance, err := std_http.NewProvider(&config)
		if err != nil {
			panic(fmt.Errorf("failed to create oidc provider: %w", err))
		}

		di.SingletonAs[oidc.Provider](func() oidc.Provider { return instance })
	default:
		panic(fmt.Errorf("invalid 'OIDC_PROVIDER': %s", provider))
	}

}
