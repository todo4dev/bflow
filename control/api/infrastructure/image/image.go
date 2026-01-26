// infrastructure/image/image.go
package image

import (
	"fmt"
	"src/infrastructure/image/resizer"

	"github.com/leandroluk/gox/env"
)

func Provide() {
	provider := env.Get("API_IMAGE_PROVIDER", "resizer")
	switch provider {
	case "resizer":
		resizer.Provide()
	default:
		panic(fmt.Errorf("invalid 'API_IMAGE_PROVIDER': %s", provider))
	}

}
