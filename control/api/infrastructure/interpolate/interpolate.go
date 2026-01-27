// infrastructure/interpolate/interpolate.go
package interpolate

import (
	"fmt"
	"src/infrastructure/interpolate/std"

	"github.com/leandroluk/gox/env"
)

func Provide() {
	provider := env.Get("API_INTERPOLATE_PROVIDER", "std")
	switch provider {
	case "std":
		std.Provide()
	default:
		panic(fmt.Errorf("invalid 'API_INTERPOLATE_PROVIDER': %s", provider))
	}
}
