// application/port/cache/rate_limit.go
package cache

import "time"

type RateLimit interface {
	Allow(key string, limit int, window time.Duration) (allowed bool, remaining int, resetAt time.Time, err error)
}
