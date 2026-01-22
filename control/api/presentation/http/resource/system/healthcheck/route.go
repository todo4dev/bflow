// presentation/http/resource/system/healthcheck/route.go
package healthcheck

import (
	"errors"
	"src/application/system/healthcheck"
	"src/presentation/http/router"

	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/oas"
)

var Route = router.
	Route(func(c *fiber.Ctx) error {
		usecase := di.Resolve[*healthcheck.Handler]()
		result, err := usecase.Handle(c.Context())
		if err != nil {
			return err
		}
		return c.JSON(result)
	}).
	Operation(func(o *oas.Operation) {
		o.Tags("System").Summary("Healthcheck").
			Description("Checks connectivity status of services").
			Response("200", func(r *oas.Response) {
				r.Description("Ok").
					Json(func(m *oas.MediaType) {
						m.Schema(func(s *oas.Schema) {
							s.Object().
								Required("uptime", func(p *oas.Schema) { p.String().Example("1h15m30s") }).
								Required("status", func(p *oas.Schema) { p.Boolean().Example(false) }).
								Required("services", func(p *oas.Schema) {
									p.Object().
										Required("database", func(item *oas.Schema) { item.String().Nullable().Example("failed to connect") }).
										Required("cache", func(item *oas.Schema) { item.String().Nullable().Example("failed to connect") }).
										Required("broker", func(item *oas.Schema) { item.String().Nullable().Example("failed to connect") }).
										Required("storage", func(item *oas.Schema) { item.String().Nullable().Example("failed to connect") })
								})
						}).Example(healthcheck.Result{
							Uptime: "1h15m30s",
							Status: false,
							Services: map[string]error{
								"database": nil,
								"cache":    errors.New("failed to connect"),
								"broker":   nil,
								"storage":  nil,
							},
						})
					})
			})
	})
