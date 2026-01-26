// presentation/http/route/auth/activate_account/activate_account.go
package activate_account

import (
	usecase "src/application/auth/activate_account"
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
		if err := di.Resolve[*usecase.Handler]().Handle(c.UserContext(), &data); err != nil {
			return err
		}
		return c.SendStatus(fiber.StatusOK)
	}).
	Operation(func(o *oas.Operation) {
		o.Tags("Auth").Summary("Activate account").
			Description("Activate account using OTP sent by email")
		server.BodyJson(o, server.SchemaAs[usecase.Data]())
		server.ResponseStatus(o, fiber.StatusOK, "Account activated")
		server.ResponseIssueAs[*validate.ValidationError](o, fiber.StatusBadRequest)
		server.ResponseIssueAs[*issue.AccountNotFound](o, fiber.StatusNotFound)
		server.ResponseIssueAs[*issue.AccountAlreadyActivated](o, fiber.StatusNotAcceptable)
		server.ResponseIssueAs[*issue.AccountInvalidOTP](o, fiber.StatusConflict)
		server.ResponseStatus(o, fiber.StatusUnprocessableEntity, "Unprocessable entity")
	})
