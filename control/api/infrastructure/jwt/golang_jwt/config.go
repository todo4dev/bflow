// infrastructure/jwt/golang_jwt/config.go
package golang_jwt

import (
	"time"

	"github.com/leandroluk/go/v"
)

type GolangJWTConfig struct {
	Algorithm  string
	Issuer     string
	Audience   string
	PrivateKey string
	PublicKey  string
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}

var GolangJWTConfigSchema = v.Object(func(t *GolangJWTConfig, s *v.ObjectSchema[GolangJWTConfig]) {
	s.Field(&t.Algorithm).Text().Required().OneOf("HS256", "RS256").Default("HS256")
	s.Field(&t.Issuer).Text().Required().Default("issuer")
	s.Field(&t.Audience).Text().Required().Default("audience")
	s.Field(&t.PrivateKey).Text().Required().Default("key")
	s.Field(&t.PublicKey).Text().Required().Default("key")
	s.Field(&t.AccessTTL).Duration().Required().Min(time.Second).Default(time.Minute * 15)
	s.Field(&t.RefreshTTL).Duration().Required().Min(time.Second).Default(time.Hour * 24)
})
