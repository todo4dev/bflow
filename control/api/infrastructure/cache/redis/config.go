// infrastructure/cache/redis/config.go
package redis

import (
	"time"

	"github.com/leandroluk/gox/validate"
)

var configSchema = validate.Object(func(t *Config, s *validate.ObjectSchema[Config]) {
	s.Field(&t.URL).Text().URL().Required()
	s.Field(&t.TTL).Duration().Min(time.Second).Default(time.Minute * 15)
})

type Config struct {
	URL string
	TTL time.Duration
}

func (c *Config) Validate() (err error) {
	_, err = configSchema.Validate(c)
	return err
}
