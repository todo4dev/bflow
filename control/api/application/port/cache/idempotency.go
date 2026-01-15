// control/api/application/port/cache/idempotency.go
package cache

import "time"

type Idempotency interface {
	Acquire(key string, ttl time.Duration) (acquired bool, err error)
	Release(key string) error
}
