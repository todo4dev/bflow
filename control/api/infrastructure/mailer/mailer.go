package mailer

import (
	"fmt"
	"src/infrastructure/mailer/wneessen_go_mail"
	"src/port/mailer"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

func Provide() {
	di.SingletonAs[mailer.Sender](func() mailer.Sender {
		provider := env.Get("MAILER_PROVIDER", "wneessen_go_mail")
		switch provider {
		case "wneessen_go_mail":
			config := wneessen_go_mail.Config{
				Host:        env.Get("MAILER_HOST", "localhost"),
				Port:        env.Get("MAILER_PORT", 587),
				Username:    env.Get("MAILER_USERNAME", ""),
				Password:    env.Get("MAILER_PASSWORD", ""),
				FromAddress: env.Get("MAILER_FROM", "noreply@bflow.dev"),
			}
			if _, err := wneessen_go_mail.ConfigSchema.Validate(&config); err != nil {
				panic(fmt.Errorf("mailer config validation failed: %w", err))
			}
			instance, err := wneessen_go_mail.NewSender(config)
			if err != nil {
				panic(fmt.Errorf("failed to create mailer: %w", err))
			}
			return instance
		default:
			panic(fmt.Errorf("invalid 'MAILER_PROVIDER': %s", provider))
		}
	})
}
