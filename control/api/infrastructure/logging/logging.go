package logging

import (
	"fmt"
	"src/infrastructure/logging/rs_zerolog"
	"src/port/logging"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

func Provide() {
	di.SingletonAs[logging.Logger](func() logging.Logger {
		provider := env.Get("LOGGING_PROVIDER", "rs_zerolog")
		switch provider {
		case "rs_zerolog":
			config := rs_zerolog.Config{
				Level:       env.Get("LOGGING_LEVEL", "info"),
				ServiceName: env.Get("APP_NAME", "bflow-control"),
			}
			if _, err := rs_zerolog.ConfigSchema.Validate(&config); err != nil {
				panic(fmt.Errorf("logging config validation failed: %w", err))
			}
			instance, err := rs_zerolog.NewLogger(config)
			if err != nil {
				panic(fmt.Errorf("failed to create logger: %w", err))
			}
			return instance
		default:
			panic(fmt.Errorf("invalid 'LOGGING_PROVIDER': %s", provider))
		}
	})
}
