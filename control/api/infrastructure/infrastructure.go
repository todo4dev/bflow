// infrastructure/infrastructure.go
package infrastructure

import (
	"src/infrastructure/broker"
	"src/infrastructure/cache"
	"src/infrastructure/database"
	"src/infrastructure/jwt"
	"src/infrastructure/logging"
	"src/infrastructure/mailing"
	"src/infrastructure/oidc"
	"src/infrastructure/storage"
)

func Provide() {
	broker.Provide()
	cache.Provide()
	database.Provide()
	jwt.Provide()
	logging.Provide()
	mailing.Provide()
	oidc.Provide()
	storage.Provide()
}
