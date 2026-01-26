// infrastructure/mailing/smtp/config.go
package smtp

import "github.com/leandroluk/gox/validate"

var configSchema = validate.Object(func(t *Config, s *validate.ObjectSchema[Config]) {
	s.Field(&t.Host).Text().Required()
	s.Field(&t.Port).Number().Min(1).Max(65535).Required()
	s.Field(&t.Username).Text().Required()
	s.Field(&t.Password).Text().Required()
	s.Field(&t.FromAddress).Text().Email().Required()
})

type Config struct {
	Host         string
	Port         int
	Username     string
	Password     string
	FromAddress  string
	TemplatePath string
}

func (c *Config) Validate() (err error) {
	_, err = configSchema.Validate(c)
	return err
}
