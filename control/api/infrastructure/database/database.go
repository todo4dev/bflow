package database

import (
	"fmt"
	"src/infrastructure/database/pgx"
	"src/port/database"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

func Provide() {
	di.SingletonAs[database.Client](func() database.Client {
		provider := env.Get("DATABASE_PROVIDER", "pgx")
		switch provider {
		case "pgx":
			config := pgx.Config{
				DSN: env.Get("DATABASE_DSN", "postgres://user:pass@localhost:5432/bflow?sslmode=disable"),
			}
			if _, err := pgx.ConfigSchema.Validate(&config); err != nil {
				panic(fmt.Errorf("database config validation failed: %w", err))
			}
			instance, err := pgx.NewClient(config)
			if err != nil {
				panic(fmt.Errorf("failed to create database client: %w", err))
			}
			return instance
		default:
			panic(fmt.Errorf("invalid 'DATABASE_PROVIDER': %s", provider))
		}
	})
}
