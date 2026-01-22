// infrastructure/database/jackc_pgx/config.go
package jackc_pgx

import "github.com/leandroluk/gox/validate"

var configSchema = validate.Object(func(t *Config, s *validate.ObjectSchema[Config]) {
	s.Field(&t.DSN).Text().FieldContains("postgres://").Required()
})

type Config struct {
	DSN string
}

func (c *Config) Validate() (err error) {
	_, err = configSchema.Validate(c)
	return err
}
