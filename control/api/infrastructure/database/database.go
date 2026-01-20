package database

import (
	"fmt"
	"src/domain/identity/repository"
	"src/infrastructure/database/pgx"
	"src/infrastructure/database/pgx/identity"
	"src/port/database"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

func Provide() {
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
		di.SingletonAs[database.Client](func() database.Client { return instance })
		di.SingletonAs[repository.Account](identity.NewAccountRepository)
	default:
		panic(fmt.Errorf("invalid 'DATABASE_PROVIDER': %s", provider))
	}
}
