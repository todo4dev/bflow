// infrastructure/logging/rs_zerolog/config.go
package rs_zerolog

import v "github.com/leandroluk/gox/validate"

type Config struct {
	Level       string
	ServiceName string
}

var ConfigSchema = v.Object(func(t *Config, s *v.ObjectSchema[Config]) {
	s.Field(&t.Level).Text().OneOf("debug", "info", "warn", "error").Default("info")
	s.Field(&t.ServiceName).Text().Required()
})
