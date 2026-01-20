// infrastructure/oidc/config.go
package oidc

import v "github.com/leandroluk/gox/validate"

type Config struct {
	BaseURI               string
	MicrosoftClientID     string
	MicrosoftClientSecret string
	MicrosoftCallbackURI  string
	GoogleClientID        string
	GoogleClientSecret    string
	GoogleCallbackURI     string
}

var ConfigSchema = v.Object(func(t *Config, s *v.ObjectSchema[Config]) {
	s.Field(&t.BaseURI).Text().Default("http://localhost:4000")
	s.Field(&t.MicrosoftClientID).Text().Required()
	s.Field(&t.MicrosoftClientSecret).Text().Required()
	s.Field(&t.MicrosoftCallbackURI).Text().Required()
	s.Field(&t.GoogleClientID).Text().Required()
	s.Field(&t.GoogleClientSecret).Text().Required()
	s.Field(&t.GoogleCallbackURI).Text().Required()
})
