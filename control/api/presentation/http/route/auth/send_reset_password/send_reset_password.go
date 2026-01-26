// presentation/http/route/auth/send_reset_password/send_reset_password.go
package send_reset_password

import (
	usecase "src/application/auth/send_reset_password"
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

		if err := di.Resolve[*usecase.Handler]().Handle(c.Context(), &data); err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusAccepted)
	}).
	Operation(func(o *oas.Operation) {
		o.Tags("Auth").Summary("Send Reset Password").
			Description("Initiates the password recovery process by sending an email with a reset code.")
		server.BodyJson(o, server.SchemaAs[usecase.Data]())
		server.ResponseStatus(o, fiber.StatusAccepted, "Recovery email sent")
		server.ResponseIssueAs[*issue.AccountNotFound](o, fiber.StatusNotFound)
		server.ResponseIssueAs[*validate.ValidationError](o, fiber.StatusUnprocessableEntity)
	})
