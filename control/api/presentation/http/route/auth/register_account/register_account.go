// presentation/http/route/auth/register_account/register_account.go
package register_account

import (
	usecase "src/application/auth/register_account"
	"src/domain/issue"
	"src/presentation/http/server"

	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/oas"
)

var Route = server.
	Route(func(c *server.Context) error {
		var data usecase.Data
		if err := c.BodyParser(&data); err != nil {
			return c.SendStatus(fiber.StatusUnprocessableEntity)
		}

		if err := di.Resolve[*usecase.Handler]().Handle(c.Context(), &data); err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusCreated)
	}).
	Operation(func(o *oas.Operation) {
		o.Tags("Auth").Summary("Register Account").
			Description("Registers a new user account using email and password.")
		server.BodyJson(o, server.SchemaAs[usecase.Data]())
		server.ResponseStatus(o, fiber.StatusCreated, "Account created successfully")
		server.ResponseIssueAs[*issue.AccountEmailInUse](o, fiber.StatusConflict)
	})
