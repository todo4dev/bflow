package billing

import (
	"src/presentation/http/router"

	"github.com/gofiber/fiber/v2"
)

var Group = router.Group("/billing", func(g *router.GroupRouter) {
	g.Get("/", router.Route(func(c *fiber.Ctx) error {
		return c.SendString("billing")
	}))
})
