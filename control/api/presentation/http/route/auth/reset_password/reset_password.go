// presentation/http/resource/auth/reset_password/reset_password.go
package reset_password

import (
	usecase "src/application/auth/reset_password"
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

		return c.SendStatus(fiber.StatusOK)
	}).
	Operation(func(o *oas.Operation) {
		o.Tags("Auth").Summary("Reset Password").
			Description("Resets the user's password using the OTP code sent by email.")
		server.BodyJson(o, server.SchemaAs[usecase.Data]())
		server.ResponseStatus(o, fiber.StatusOK, "Password reset successfully")
		server.ResponseIssueAs[*issue.AccountNotFound](o, fiber.StatusNotFound)
		server.ResponseIssueAs[*issue.AccountInvalidOTP](o, fiber.StatusConflict)
		server.ResponseIssueAs[*validate.ValidationError](o, fiber.StatusUnprocessableEntity)
	})
