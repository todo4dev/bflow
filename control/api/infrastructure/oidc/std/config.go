// infrastructure/oidc/std/config.go
package std

import "github.com/leandroluk/gox/validate"

var configSchema = validate.Object(func(t *Config, s *validate.ObjectSchema[Config]) {
	s.Field(&t.BaseURI).Text().Default("http://localhost:4000")
	s.Field(&t.MicrosoftClientID).Text().Required()
	s.Field(&t.MicrosoftClientSecret).Text().Required()
	s.Field(&t.MicrosoftCallbackURI).Text().Required()
	s.Field(&t.GoogleClientID).Text().Required()
	s.Field(&t.GoogleClientSecret).Text().Required()
	s.Field(&t.GoogleCallbackURI).Text().Required()
})

type Config struct {
	BaseURI               string
	MicrosoftClientID     string
	MicrosoftClientSecret string
	MicrosoftCallbackURI  string
	GoogleClientID        string
	GoogleClientSecret    string
	GoogleCallbackURI     string
}

func (c *Config) Validate() (err error) {
	_, err = configSchema.Validate(c)
	return err
}
