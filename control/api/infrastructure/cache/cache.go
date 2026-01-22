// infrastructure/cache/cache.go
package cache

import (
	"fmt"
	"src/infrastructure/cache/go_redis"
	"src/port/cache"
	"time"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

func Provide() {
	provider := env.Get("CACHE_PROVIDER", "go_redis")
	switch provider {
	case "go_redis":
		config := go_redis.Config{
			URL: env.Get("CACHE_URL", "redis://localhost:6379"),
			TTL: time.Duration(env.Get("CACHE_TTL_SECONDS", 900)) * time.Second,
		}
		if err := config.Validate(); err != nil {
			panic(fmt.Errorf("cache config validation failed: %w", err))
		}

		instance, err := go_redis.NewClient(&config)
		if err != nil {
			panic(fmt.Errorf("failed to create cache client: %w", err))
		}

		di.SingletonAs[cache.Client](func() cache.Client { return instance })
	default:
		panic(fmt.Errorf("invalid 'CACHE_PROVIDER': %s", provider))
	}
}
