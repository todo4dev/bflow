// infrastructure/cache/go_redis/config.go
package go_redis

import (
	"time"

	v "github.com/leandroluk/gox/validate"
)

type GoRedisConfig struct {
	URL string
	TTL time.Duration
}

var GoRedisConfigSchema = v.Object(func(t *GoRedisConfig, s *v.ObjectSchema[GoRedisConfig]) {
	s.Field(&t.URL).Text().URL().Required()
	s.Field(&t.TTL).Duration().Min(time.Second).Default(time.Minute * 15)
})
