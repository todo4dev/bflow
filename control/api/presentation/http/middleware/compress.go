// presentation/http/middleware/compress.go
package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func Compress() fiber.Handler {
	config := compress.Config{
		Level: compress.LevelBestSpeed,
	}

	return compress.New(config)
}
