// infrastructure/mailing/mailing.go
package mailing

import (
	"fmt"
	"src/infrastructure/mailing/smtp"

	"github.com/leandroluk/gox/env"
)

func Provide() {
	provider := env.Get("API_MAILING_PROVIDER", "smtp")
	switch provider {
	case "smtp":
		smtp.Provide()
	default:
		panic(fmt.Errorf("invalid 'API_MAILING_PROVIDER': %s", provider))
	}
}
