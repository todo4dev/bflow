// infrastructure/mailer/wneessen_go_mail/config.go
package wneessen_go_mail

import v "github.com/leandroluk/gox/validate"

type Config struct {
	Host        string
	Port        int
	Username    string
	Password    string
	FromAddress string
}

var ConfigSchema = v.Object(func(t *Config, s *v.ObjectSchema[Config]) {
	s.Field(&t.Host).Text().Required()
	s.Field(&t.Port).Number().Min(1).Max(65535).Required()
	s.Field(&t.Username).Text().Required()
	s.Field(&t.Password).Text().Required()
	s.Field(&t.FromAddress).Text().Email().Required()
})
