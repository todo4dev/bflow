// application/port/cache/store.go
package cache

import "time"

type Store interface {
	Get(key string) (value []byte, found bool, err error)
	Set(key string, value []byte, ttl time.Duration) error
	Delete(key string) error
	Expire(key string, ttl time.Duration) (updated bool, err error)
}
