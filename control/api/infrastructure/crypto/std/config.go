// infrastructure/cripto/std/config.go
package std

import "github.com/leandroluk/gox/validate"

var configSchema = validate.Object(func(t *Config, s *validate.ObjectSchema[Config]) {
	s.Field(&t.Key).Text().Required().Min(32).Max(32)
})

type Config struct {
	Key string
}

func (c *Config) Validate() error {
	_, err := configSchema.Validate(c)
	return err
}
