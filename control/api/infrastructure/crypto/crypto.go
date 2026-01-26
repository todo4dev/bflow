// infrastructure/crypto/crypto.go
package crypto

import (
	"fmt"
	"src/infrastructure/crypto/std"

	"github.com/leandroluk/gox/env"
)

func Provide() {
	provider := env.Get("API_CRIPTO_PROVIDER", "std")
	switch provider {
	case "std":
		std.Provide()
	default:
		panic(fmt.Errorf("invalid 'API_CRIPTO_PROVIDER': %s", provider))
	}
}
