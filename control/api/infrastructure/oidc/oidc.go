// infrastructure/oidc/oidc.go
package oidc

import (
	"fmt"
	"src/infrastructure/oidc/std"

	"github.com/leandroluk/gox/env"
)

func Provide() {
	provider := env.Get("API_OIDC_PROVIDER", "std")
	switch provider {
	case "std":
		std.Provide()
	default:
		panic(fmt.Errorf("invalid 'API_OIDC_PROVIDER': %s", provider))
	}

}
