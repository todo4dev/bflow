package recover

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func New(optionalConfig ...recover.Config) fiber.Handler {
	config := recover.Config{
		EnableStackTrace: true,
	}

	if len(optionalConfig) > 0 {
		config = optionalConfig[0]
	}

	return recover.New(config)
}
