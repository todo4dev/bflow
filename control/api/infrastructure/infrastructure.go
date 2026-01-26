// infrastructure/infrastructure.go
package infrastructure

import (
	"src/infrastructure/broker"
	"src/infrastructure/cache"
	"src/infrastructure/crypto"
	"src/infrastructure/database"
	"src/infrastructure/image"
	"src/infrastructure/jwt"
	"src/infrastructure/logger"
	"src/infrastructure/mailing"
	"src/infrastructure/oidc"
	"src/infrastructure/storage"
)

func Provide() {
	broker.Provide()
	cache.Provide()
	crypto.Provide()
	database.Provide()
	image.Provide()
	jwt.Provide()
	logger.Provide()
	mailing.Provide()
	oidc.Provide()
	storage.Provide()
}
