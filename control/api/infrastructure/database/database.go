// infrastructure/database/database.go
package database

import (
	"fmt"
	"src/domain"
	"src/infrastructure/database/jackc_pgx"
	"src/port/database"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

func Provide() {
	provider := env.Get("DATABASE_PROVIDER", "jackc_pgx")
	switch provider {
	case "jackc_pgx":
		config := jackc_pgx.Config{
			DSN: env.Get("DATABASE_DSN", "postgres://user:pass@localhost:5432/bflow?sslmode=disable"),
		}
		if err := config.Validate(); err != nil {
			panic(fmt.Errorf("database config validation failed: %w", err))
		}

		instance, err := jackc_pgx.NewClient(&config)
		if err != nil {
			panic(fmt.Errorf("failed to create database client: %w", err))
		}

		di.SingletonAs[database.Client](func() database.Client { return instance })
		di.SingletonAs[domain.Uow](jackc_pgx.NewUow)
	default:
		panic(fmt.Errorf("invalid 'DATABASE_PROVIDER': %s", provider))
	}
}
