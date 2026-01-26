// presentation/http/route/auth/login_using_credential/login_using_credential.go
package login_using_credential

import (
	usecase "src/application/auth/login_using_credential"
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
			return c.SendStatus(fiber.StatusUnprocessableEntity)
		}

		result, err := di.Resolve[*usecase.Handler]().Handle(c.Context(), &data)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}).
	Operation(func(o *oas.Operation) {
		o.Tags("Auth").Summary("Login using credential").
			Description("Authenticates a user by verifying their email and password.")
		server.BodyJson(o, server.SchemaAs[usecase.Data]())
		server.ResponseStatus(o, fiber.StatusOK, "Login successful", server.SchemaAs[usecase.Result]())
		server.ResponseIssueAs[*issue.AccountInvalidCredentials](o, fiber.StatusUnauthorized)
		server.ResponseIssueAs[*issue.AccountNotVerified](o, fiber.StatusForbidden)
		server.ResponseIssueAs[*issue.AccountDeactivated](o, fiber.StatusNotAcceptable)
		server.ResponseIssueAs[*validate.ValidationError](o, fiber.StatusUnprocessableEntity)
	})
