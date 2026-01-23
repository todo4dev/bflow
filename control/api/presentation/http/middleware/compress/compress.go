package compress

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func New(optionalConfig ...compress.Config) fiber.Handler {
	config := compress.Config{
		Level: compress.LevelBestSpeed,
	}

	if len(optionalConfig) > 0 {
		config = optionalConfig[0]
	}

	return compress.New(config)
}
