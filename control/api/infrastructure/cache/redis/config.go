// infrastructure/cache/redis/config.go
package redis

import "github.com/leandroluk/go/v"

type Config struct {
	URL string
}

var ConfigSchema = v.Object(func(c *Config, s *v.ObjectSchema[Config]) {
	s.Field(&c.URL, func(ctx *v.Context, value any) (any, error) {
		return v.Text().Required().ValidateAny(value, ctx.Options)
	})
})
