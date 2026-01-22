// infrastructure/logging/logging.go
package logging

import (
	"fmt"
	"src/infrastructure/logging/rs_zerolog"
	"src/port/logging"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

func Provide() {
	provider := env.Get("LOGGING_PROVIDER", "rs_zerolog")
	switch provider {
	case "rs_zerolog":
		config := rs_zerolog.Config{
			Level:       env.Get("LOGGING_LEVEL", "info"),
			ServiceName: env.Get("APP_NAME", "bflow-control"),
		}
		if err := config.Validate(); err != nil {
			panic(fmt.Errorf("logging config validation failed: %w", err))
		}

		instance, err := rs_zerolog.NewLogger(&config)
		if err != nil {
			panic(fmt.Errorf("failed to create logger: %w", err))
		}

		di.SingletonAs[logging.Logger](func() logging.Logger { return instance })
	default:
		panic(fmt.Errorf("invalid 'LOGGING_PROVIDER': %s", provider))
	}

}
