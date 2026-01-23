// presentation/http/rest/system/healthcheck/route.go
package healthcheck

import (
	"src/application/system/healthcheck"
	"src/presentation/http/router"

	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/oas"
)

func handler(c *fiber.Ctx) error {
	usecase := di.Resolve[*healthcheck.Handler]()
	result, err := usecase.Handle(c.Context())
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(result)
	}
	return c.JSON(result)
}

func operation(o *oas.Operation) {
	o.Tags("System").Summary("Healthcheck").
		Description("Checks connectivity status of services")
	resultSchema := router.SchemaAs[healthcheck.Result]()
	router.ResponseStatus(o, fiber.StatusOK, "Healthy", resultSchema)
	router.ResponseStatus(o, fiber.StatusServiceUnavailable, "Unhealthy", resultSchema)
}

var Route = router.
	Route(handler).
	Operation(operation)
