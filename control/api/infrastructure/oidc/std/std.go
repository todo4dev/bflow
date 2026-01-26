// infrastructure/oidc/std/std.go
package std

import (
	"fmt"
	"src/port/oidc"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

func Provide() {
	config := Config{
		BaseURI:               env.Get("API_OIDC_STD_BASE_URI", "http://localhost:3000"),
		MicrosoftClientID:     env.Get("API_OIDC_STD_MICROSOFT_CLIENT_ID", ""),
		MicrosoftClientSecret: env.Get("API_OIDC_STD_MICROSOFT_CLIENT_SECRET", ""),
		MicrosoftCallbackURI:  env.Get("API_OIDC_STD_MICROSOFT_CALLBACK", "/auth/microsoft/callback"),
		GoogleClientID:        env.Get("API_OIDC_STD_GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret:    env.Get("API_OIDC_STD_GOOGLE_CLIENT_SECRET", ""),
		GoogleCallbackURI:     env.Get("API_OIDC_STD_GOOGLE_CALLBACK", "/auth/google/callback"),
	}
	if err := config.Validate(); err != nil {
		panic(fmt.Errorf("oidc config validation failed: %w", err))
	}

	instance, err := New(&config)
	if err != nil {
		panic(fmt.Errorf("failed to create oidc provider: %w", err))
	}

	di.SingletonAs[oidc.Provider](func() oidc.Provider { return instance })
}
