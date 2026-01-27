// infrastructure/interpolate/std/config.go
package std

import "github.com/leandroluk/gox/validate"

var configSchema = validate.Object(func(t *Config, s *validate.ObjectSchema[Config]) {
	s.Field(&t.Path).Text().Required().Default("./template")
})

type Config struct {
	Path string
}

func (c *Config) Validate() (err error) {
	_, err = configSchema.Validate(c)
	return err
}
