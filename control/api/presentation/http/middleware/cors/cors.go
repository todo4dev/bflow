package cors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func New(optionalOrigin ...string) fiber.Handler {
	config := cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
		ExposeHeaders:    "Content-Length, Content-Type",
		MaxAge:           86400,
	}

	if len(optionalOrigin) > 0 {
		config.AllowOrigins = optionalOrigin[0]
	}

	config.AllowCredentials = config.AllowOrigins != "*"

	return cors.New(config)
}
