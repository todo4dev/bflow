package swagger_json

import (
	"encoding/json"

	"src/presentation/http/router"

	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/di"
)

func handler(c *fiber.Ctx) error {
	r := di.Resolve[*router.Router]()
	if r == nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Router not initialized")
	}

	data, err := json.Marshal(r.GenerateOAS())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to serialize OpenAPI")
	}

	c.Set("Content-Type", "application/json")
	return c.Send(data)
}

var Route = router.Route(handler)
