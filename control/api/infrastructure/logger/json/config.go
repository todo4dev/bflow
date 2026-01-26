// infrastructure/logging/json/config.go
package json

import "github.com/leandroluk/gox/validate"

var configSchema = validate.Object(func(t *Config, s *validate.ObjectSchema[Config]) {
	s.Field(&t.Level).Text().OneOf("debug", "info", "warn", "error").Default("info")
	s.Field(&t.ServiceName).Text().Required()
})

type Config struct {
	Level       string
	ServiceName string
}

func (c *Config) Validate() (err error) {
	_, err = configSchema.Validate(c)
	return err
}
