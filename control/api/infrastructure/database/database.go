// infrastructure/database/database.go
package database

import (
	"fmt"
	"src/infrastructure/database/postgres"

	"github.com/leandroluk/gox/env"
)

func Provide() {
	provider := env.Get("API_DATABASE_PROVIDER", "postgres")
	switch provider {
	case "postgres":
		postgres.Provide()
	default:
		panic(fmt.Errorf("invalid 'API_DATABASE_PROVIDER': %s", provider))
	}
}
