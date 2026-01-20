package jwt

import (
	"fmt"
	"src/infrastructure/jwt/golang_jwt"
	"src/port/jwt"
	"time"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

func Provide() {
	provider := env.Get("JWT_PROVIDER", "golang_jwt")
	switch provider {
	case "golang_jwt":
		config := golang_jwt.Config{
			Algorithm:  env.Get("JWT_ALGORITHM", "HS256"),
			Issuer:     env.Get("JWT_ISSUER", "bflow"),
			Audience:   env.Get("JWT_AUDIENCE", "bflow-api"),
			PrivateKey: env.Get("JWT_PRIVATE_KEY", "secret-key"),
			PublicKey:  env.Get("JWT_PUBLIC_KEY", "secret-key"),
			AccessTTL:  time.Duration(env.Get("JWT_ACCESS_TTL_SECONDS", 900)) * time.Second,
			RefreshTTL: time.Duration(env.Get("JWT_REFRESH_TTL_SECONDS", 86400)) * time.Second,
		}
		if _, err := golang_jwt.ConfigSchema.Validate(&config); err != nil {
			panic(err)
		}
		instance, err := golang_jwt.NewProvider(config)
		if err != nil {
			panic(fmt.Errorf("failed to create jwt provider: %w", err))
		}
		di.SingletonAs[jwt.Provider](func() jwt.Provider { return instance })
	default:
		panic(fmt.Errorf("invalid 'JWT_PROVIDER': %s", provider))
	}
}
