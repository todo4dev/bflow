// presentation/http/route/health/route.go
package health

import (
	"src/application/health"
	"src/presentation/http/server"

	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/oas"
)

var Route = server.
	NewRoute(func(c *server.Context) error {
		result, err := di.Resolve[*health.Handler]().Handle()
		if err != nil {
			return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
		}
		return c.JSON(result)
	}).
	Operation(func(o *oas.Operation) {
		o.Tags("System").Summary("Healthcheck").
			Description("Checks connectivity status of services")
		server.ResponseStatus(o, fiber.StatusOK, "Healthy", server.SchemaAs[health.Result]())
		server.ResponseStatus(o, fiber.StatusServiceUnavailable, "Unhealthy")
	})
