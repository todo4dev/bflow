// presentation/http/route/openapi_json/route.go
package openapi_json

import (
	"encoding/json"
	"src/presentation/http/server"

	"github.com/gofiber/fiber/v2"
)

var Route = server.
	NewRoute(func(c *server.Context) error {
		if c.Router == nil {
			return fiber.NewError(fiber.StatusServiceUnavailable, "Router not initialized")
		}

		data, err := json.Marshal(c.Router.GenerateJSON())
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to serialize OpenAPI")
		}

		c.Set("Content-Type", "application/json")
		return c.Send(data)
	})
