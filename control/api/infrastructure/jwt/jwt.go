// infrastructure/jwt/jwt.go
package jwt

import (
	"fmt"
	"src/infrastructure/jwt/jose"

	"github.com/leandroluk/gox/env"
)

func Provide() {
	provider := env.Get("API_JWT_PROVIDER", "jose")
	switch provider {
	case "jose":
		jose.Provide()
	default:
		panic(fmt.Errorf("invalid 'API_JWT_PROVIDER': %s", provider))
	}
}
