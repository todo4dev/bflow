// presentation/http/router/config.go
package router

import "github.com/leandroluk/gox/validate"

var configSchema = validate.Object(func(t *Config, s *validate.ObjectSchema[Config]) {
	s.Field(&t.Port).Number().Required().Integer().Min(1)
	s.Field(&t.BasePath).Text().Required().Default("")
	s.Field(&t.EnableSwagger).Boolean().Default(true)
	s.Field(&t.SwaggerPath).Text().Required().Default("/")
	s.Field(&t.SwaggerTitle).Text().Required().Default("Bflow - Control Plane API")
	s.Field(&t.SwaggerDescription).Text().Required().Default("API for Control Plane of Bflow solution")
	s.Field(&t.SwaggerContactName).Text().Required().Default("Leandro Santiago Gomes")
	s.Field(&t.SwaggerContactURL).Text().Required().Default("https://github.com/leandroluk")
	s.Field(&t.SwaggerContactEmail).Text().Required().Default("leandroluk@gmail.com")
	s.Field(&t.SwaggerLicenseName).Text().Required().Default("MIT")
	s.Field(&t.SwaggerLicenseURL).Text().Required().Default("https://opensource.org/licenses/MIT")
	s.Field(&t.SwaggerVersion).Text().Required().Default("1.0.0")
})

type Config struct {
	Port                int
	BasePath            string
	EnableSwagger       bool
	SwaggerPath         string
	SwaggerTitle        string
	SwaggerDescription  string
	SwaggerContactName  string
	SwaggerContactURL   string
	SwaggerContactEmail string
	SwaggerLicenseName  string
	SwaggerLicenseURL   string
	SwaggerVersion      string
}

func (c *Config) Validate() (err error) {
	_, err = configSchema.Validate(c)
	return err
}
