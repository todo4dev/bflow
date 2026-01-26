// infrastructure/mailing/smtp/smtp.go
package smtp

import (
	"fmt"
	"src/port/mailing"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

func Provide() {
	config := Config{
		Host:         env.Get("API_MAILING_SMTP_HOST", "localhost"),
		Port:         env.Get("API_MAILING_SMTP_PORT", 587),
		Username:     env.Get("API_MAILING_SMTP_USERNAME", ""),
		Password:     env.Get("API_MAILING_SMTP_PASSWORD", ""),
		FromAddress:  env.Get("API_MAILING_SMTP_FROM", "noreply@bflow.dev"),
		TemplatePath: env.Get("API_MAILING_SMTP_TEMPLATE_PATH", "./template"),
	}
	if err := config.Validate(); err != nil {
		panic(fmt.Errorf("mailer config validation failed: %w", err))
	}

	instance, err := New(&config)
	if err != nil {
		panic(fmt.Errorf("failed to create mailer: %w", err))
	}

	di.SingletonAs[mailing.Mailer](func() mailing.Mailer { return instance })
}
