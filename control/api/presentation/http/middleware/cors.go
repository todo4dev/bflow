package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors(origin string) fiber.Handler {
	config := cors.Config{
		AllowOrigins:     origin,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: false,
		ExposeHeaders:    "Content-Length, Content-Type",
		MaxAge:           86400,
	}
	return cors.New(config)
}
