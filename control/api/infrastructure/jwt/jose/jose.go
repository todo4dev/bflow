// infrastructure/jwt/jose/jose.go
package jose

import (
	"fmt"
	"src/port/jwt"
	"time"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

func Provide() {
	config := Config{
		Algorithm:  env.Get("API_JWT_JOSE_ALGORITHM", "HS256"),
		Issuer:     env.Get("API_JWT_JOSE_ISSUER", "bflow"),
		Audience:   env.Get("API_JWT_JOSE_AUDIENCE", "bflow-api"),
		PrivateKey: env.Get("API_JWT_JOSE_PRIVATE_KEY", "secret-key"),
		PublicKey:  env.Get("API_JWT_JOSE_PUBLIC_KEY", "secret-key"),
		AccessTTL:  time.Duration(env.Get("API_JWT_JOSE_ACCESS_TTL", 900)) * time.Second,
		RefreshTTL: time.Duration(env.Get("API_JWT_JOSE_REFRESH_TTL", 86400)) * time.Second,
	}
	if err := config.Validate(); err != nil {
		panic(fmt.Errorf("jwt config validation failed: %w", err))
	}

	instance, err := New(&config)
	if err != nil {
		panic(fmt.Errorf("failed to create jwt provider: %w", err))
	}

	di.SingletonAs[jwt.Provider](func() jwt.Provider { return instance })
}
