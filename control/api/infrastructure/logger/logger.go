// infrastructure/logging/logger.go
package logger

import (
	"fmt"
	"src/infrastructure/logger/json"

	"github.com/leandroluk/gox/env"
)

func Provide() {
	provider := env.Get("API_LOGGER_PROVIDER", "json")
	switch provider {
	case "json":
		json.Provide()
	default:
		panic(fmt.Errorf("invalid 'API_LOGGER_PROVIDER': %s", provider))
	}

}
