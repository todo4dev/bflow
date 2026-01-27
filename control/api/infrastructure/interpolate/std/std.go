// infrastructure/interpolate/std/std.go
package std

import (
	"fmt"
	"src/port/interpolate"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

func Provide() {
	config := Config{
		Path: env.Get("API_INTERPOLATE_STD_PATH", "./template"),
	}
	if err := config.Validate(); err != nil {
		panic(fmt.Errorf("interpolate config validation failed: %w", err))
	}

	instance, err := New(&config)
	if err != nil {
		panic(fmt.Errorf("failed to create interpolate: %w", err))
	}

	di.SingletonAs[interpolate.Renderer](func() interpolate.Renderer { return instance })
}
