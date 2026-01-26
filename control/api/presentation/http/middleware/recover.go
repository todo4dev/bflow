package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Recover() fiber.Handler {
	config := recover.Config{
		EnableStackTrace: true,
	}

	return recover.New(config)
}
