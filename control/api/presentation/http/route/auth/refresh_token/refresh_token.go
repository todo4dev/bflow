// presentation/http/route/auth/refresh_token/refresh_token.go
package refresh_token

import (
	usecase "src/application/auth/refresh_token"
	"src/domain/issue"
	"src/presentation/http/server"

	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/oas"
	"github.com/leandroluk/gox/validate"
)

var Route = server.
	Route(func(c *server.Context) error {
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
		o.Tags("Auth").Summary("Refresh Authorization Token").
			Description("Refresh the authorization token extending the session expiration.")
		server.BodyJson(o, server.SchemaAs[usecase.Data]())
		server.ResponseStatus(o, fiber.StatusOK, "Refresh successful", server.SchemaAs[usecase.Result]())
		server.ResponseIssueAs[*issue.AccountInvalidToken](o, fiber.StatusUnauthorized)
		server.ResponseIssueAs[*issue.AccountSessionExpired](o, fiber.StatusUnauthorized)
		server.ResponseIssueAs[*issue.AccountDeactivated](o, fiber.StatusForbidden)
		server.ResponseIssueAs[*validate.ValidationError](o, fiber.StatusUnprocessableEntity)
	})
