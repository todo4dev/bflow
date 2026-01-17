// infrastructure/oidc/config.go
package oidc

import "github.com/leandroluk/go/v"

type HttpConfig struct {
	BaseURI               string
	MicrosoftClientID     string
	MicrosoftClientSecret string
	MicrosoftCallbackURI  string
	GoogleClientID        string
	GoogleClientSecret    string
	GoogleCallbackURI     string
}

var HttpConfigSchema = v.Object(func(t *HttpConfig, s *v.ObjectSchema[HttpConfig]) {
	s.Field(&t.BaseURI).Text().Default("http://localhost:4000")
	s.Field(&t.MicrosoftClientID).Text().Required()
	s.Field(&t.MicrosoftClientSecret).Text().Required()
	s.Field(&t.MicrosoftCallbackURI).Text().Required()
	s.Field(&t.GoogleClientID).Text().Required()
	s.Field(&t.GoogleClientSecret).Text().Required()
	s.Field(&t.GoogleCallbackURI).Text().Required()
})
