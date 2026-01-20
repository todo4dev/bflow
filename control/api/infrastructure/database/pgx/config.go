// infrastructure/sqldb/pgx/config.go
package pgx

import v "github.com/leandroluk/gox/validate"

type Config struct {
	DSN string
}

var ConfigSchema = v.Object(func(t *Config, s *v.ObjectSchema[Config]) {
	s.Field(&t.DSN).Text().FieldContains("postgres://").Required()
})
