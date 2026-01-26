// infrastructure/logging/json/json.go
package json

import (
	"fmt"
	"src/port/logger"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

func Provide() {
	config := Config{
		Level:       env.Get("API_LOGGER_LEVEL", "info"),
		ServiceName: env.Get("API_NAME", "bflow-control"),
	}
	if err := config.Validate(); err != nil {
		panic(fmt.Errorf("logging config validation failed: %w", err))
	}

	instance, err := New(&config)
	if err != nil {
		panic(fmt.Errorf("failed to create logger: %w", err))
	}

	di.SingletonAs[logger.Client](func() logger.Client { return instance })
}
