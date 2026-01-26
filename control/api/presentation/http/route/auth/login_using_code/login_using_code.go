// presentation/http/route/auth/login_using_code/login_using_code.go
package login_using_code

import (
	usecase "src/application/auth/login_using_code"
	"src/domain/issue"
	"src/presentation/http/server"

	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/oas"
	"github.com/leandroluk/gox/validate"
)

var Route = server.
	NewRoute(func(c *server.Context) error {
		var data usecase.Data
		if err := c.BodyParser(&data); err != nil {
			return fiber.NewError(fiber.StatusUnprocessableEntity, "invalid body")
		}

		result, err := di.Resolve[*usecase.Handler]().Handle(c.Context(), &data)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}).
	Operation(func(o *oas.Operation) {
		o.Tags("Auth").Summary("Login Using Code").
			Description("Authenticate using an encrypted code received from SSO callback.")
		server.BodyJson(o, server.SchemaAs[usecase.Data]())
		server.ResponseStatus(o, fiber.StatusOK, "Login Result", server.SchemaAs[usecase.Result]())
		server.ResponseIssueAs[*issue.AccountInvalidCredentials](o, fiber.StatusUnauthorized)
		server.ResponseIssueAs[*issue.AccountDeactivated](o, fiber.StatusNotAcceptable)
		server.ResponseIssueAs[*validate.ValidationError](o, fiber.StatusUnprocessableEntity)
	})
