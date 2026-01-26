// infrastructure/cache/redis/redis.go
package redis

import (
	"fmt"
	"src/port/cache"
	"time"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

func Provide() {
	config := Config{
		URL: env.Get("API_CACHE_REDIS_URL", "redis://localhost:6379"),
		TTL: time.Duration(env.Get("API_CACHE_REDIS_TTL_SECONDS", 900)) * time.Second,
	}
	if err := config.Validate(); err != nil {
		panic(fmt.Errorf("cache config validation failed: %w", err))
	}

	instance, err := New(&config)
	if err != nil {
		panic(fmt.Errorf("failed to create cache client: %w", err))
	}

	di.SingletonAs[cache.Client](func() cache.Client { return instance })
}
