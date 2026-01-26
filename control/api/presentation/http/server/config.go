// presentation/http/server/config.go
package server

import (
	"path"

	"github.com/leandroluk/gox/validate"
)

type Config struct {
	AppName             string
	AppPort             string
	AppPath             string
	AppOrigin           string
	OpenAPIEnable       bool
	OpenAPIUIPath       string
	OpenAPIJSONPath     string
	OpenAPITitle        string
	OpenAPIDescription  string
	OpenAPIContactName  string
	OpenAPIContactURL   string
	OpenAPIContactEmail string
	OpenAPILicenseName  string
	OpenAPILicenseURL   string
	OpenAPIVersion      string
}

func (c *Config) GetOpenAPIPath() string {
	return path.Clean(c.AppPath + "/" + c.OpenAPIUIPath)
}

func (c *Config) GetOpenAPIJsonPath() string {
	return path.Clean(c.AppPath + "/" + c.OpenAPIUIPath + "/" + c.OpenAPIJSONPath)
}

var configSchema = validate.Object(func(t *Config, s *validate.ObjectSchema[Config]) {
	s.Field(&t.AppName).Text().Required()
	s.Field(&t.AppPort).Text().Required()
	s.Field(&t.AppPath).Text().Required().Default("")
	s.Field(&t.AppOrigin).Text().Required().Default("*")
	s.Field(&t.OpenAPIEnable).Boolean().Default(true)
	s.Field(&t.OpenAPIUIPath).Text().Required().Default("/")
	s.Field(&t.OpenAPIJSONPath).Text().Required().Default("/openapi.json")
	s.Field(&t.OpenAPITitle).Text().Required().Default("Bflow - Control Plane API")
	s.Field(&t.OpenAPIDescription).Text().Required().Default("API for Control Plane of Bflow solution")
	s.Field(&t.OpenAPIContactName).Text().Required().Default("Leandro Santiago Gomes")
	s.Field(&t.OpenAPIContactURL).Text().Required().Default("https://github.com/leandroluk")
	s.Field(&t.OpenAPIContactEmail).Text().Required().Default("leandroluk@gmail.com")
	s.Field(&t.OpenAPILicenseName).Text().Required().Default("MIT")
	s.Field(&t.OpenAPILicenseURL).Text().Required().Default("https://opensource.org/licenses/MIT")
	s.Field(&t.OpenAPIVersion).Text().Required().Default("1.0.0")
})
