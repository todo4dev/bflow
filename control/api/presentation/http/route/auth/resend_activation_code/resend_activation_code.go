// presentation/http/route/auth/resend_activation_code/resend_activation_code.go
package resend_activation_code

import (
	usecase "src/application/auth/resend_activation_code"
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

		return c.SendStatus(fiber.StatusAccepted)
	}).
	Operation(func(o *oas.Operation) {
		o.Tags("Auth").Summary("Resend Activation Code").
			Description("Resends the activation code to the user's email address.")
		server.BodyJson(o, server.SchemaAs[usecase.Data]())
		server.ResponseStatus(o, fiber.StatusAccepted, "Activation code sent")
		server.ResponseIssueAs[*issue.AccountNotFound](o, fiber.StatusNotFound)
		server.ResponseIssueAs[*issue.AccountAlreadyActivated](o, fiber.StatusNotAcceptable)
	})
