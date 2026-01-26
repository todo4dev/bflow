// infrastructure/storage/storage.go
package storage

import (
	"fmt"
	"src/infrastructure/storage/s3"

	"github.com/leandroluk/gox/env"
)

func Provide() {
	provider := env.Get("API_STORAGE_PROVIDER", "s3")
	switch provider {
	case "s3":
		s3.Provide()
	default:
		panic(fmt.Errorf("invalid 'API_STORAGE_PROVIDER': %s", provider))
	}
}
