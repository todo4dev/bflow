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
		return err
	}
	return c.JSON(result)
}

func operation(o *oas.Operation) {
	o.Tags("System").Summary("Healthcheck").
		Description("Checks connectivity status of services")
	router.ResponseStatus(o, fiber.StatusOK, "Healthcheck status", router.SchemaAs[healthcheck.Result]())
}

var Route = router.
	Route(handler).
	Operation(operation)
