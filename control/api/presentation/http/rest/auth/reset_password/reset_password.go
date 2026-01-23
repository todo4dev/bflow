// presentation/http/resource/auth/reset_password/reset_password.go
package reset_password

import (
	usecase "src/application/auth/reset_password"
	"src/domain/issue"
	"src/presentation/http/router"

	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/oas"
	"github.com/leandroluk/gox/validate"
)

func handler(c *fiber.Ctx) error {
	var data usecase.Data
	if err := c.BodyParser(&data); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	if err := di.Resolve[*usecase.Handler]().Handle(c.Context(), &data); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func operation(o *oas.Operation) {
	o.Tags("Auth").Summary("Reset Password").
		Description("Resets the user's password using the OTP code sent by email.")
	router.BodyJson(o, router.SchemaAs[usecase.Data]())
	router.ResponseStatus(o, fiber.StatusOK, "Password reset successfully")
	router.ResponseIssueAs[*issue.AccountNotFound](o, fiber.StatusNotFound)
	router.ResponseIssueAs[*issue.AccountInvalidOTP](o, fiber.StatusConflict)
	router.ResponseIssueAs[*validate.ValidationError](o, fiber.StatusUnprocessableEntity)
}

var Route = router.
	Route(handler).
	Operation(operation)
