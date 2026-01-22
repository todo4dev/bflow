// presentation/http/resource/auth/register_account/route.go
package register_account

import (
	usecase "src/application/auth/register_account"
	"src/domain/issue"
	"src/presentation/http/router"

	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/oas"
)

func handler(c *fiber.Ctx) error {
	var data usecase.Data
	if err := c.BodyParser(&data); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	_, err := di.Resolve[*usecase.Handler]().Handle(c.Context(), &data)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func operation(o *oas.Operation) {
	o.Tags("Auth").Summary("Register Account").
		Description("Registers a new user account using email and password.")
	router.BodyJson(o, router.SchemaAs[usecase.Data]())
	router.ResponseStatus(o, fiber.StatusCreated, "Account created successfully")
	router.ResponseIssueAs[*issue.AccountEmailInUse](o, fiber.StatusConflict)
}

var Route = router.
	Route(handler).
	Operation(operation)
