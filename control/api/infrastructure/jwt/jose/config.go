// infrastructure/jwt/jose/config.go
package jose

import (
	"time"

	"github.com/leandroluk/gox/validate"
)

var configSchema = validate.Object(func(t *Config, s *validate.ObjectSchema[Config]) {
	s.Field(&t.Algorithm).Text().Required().OneOf("HS256", "RS256").Default("HS256")
	s.Field(&t.Issuer).Text().Required().Default("issuer")
	s.Field(&t.Audience).Text().Required().Default("audience")
	s.Field(&t.PrivateKey).Text().Required().Default("key")
	s.Field(&t.PublicKey).Text().Required().Default("key")
	s.Field(&t.AccessTTL).Duration().Required().Min(time.Second).Default(time.Minute * 15)
	s.Field(&t.RefreshTTL).Duration().Required().Min(time.Second).Default(time.Hour * 24)
})

type Config struct {
	Algorithm  string
	Issuer     string
	Audience   string
	PrivateKey string
	PublicKey  string
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}

func (c *Config) Validate() (err error) {
	_, err = configSchema.Validate(c)
	return err
}
