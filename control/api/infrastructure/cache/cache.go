// infrastructure/cache/cache.go
package cache

import (
	"fmt"
	"src/infrastructure/cache/redis"

	"github.com/leandroluk/gox/env"
)

func Provide() {
	provider := env.Get("API_CACHE_PROVIDER", "redis")
	switch provider {
	case "redis":
		redis.Provide()
	default:
		panic(fmt.Errorf("invalid 'API_CACHE_PROVIDER': %s", provider))
	}
}
