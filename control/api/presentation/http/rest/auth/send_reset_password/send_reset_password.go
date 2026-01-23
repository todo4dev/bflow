// presentation/http/rest/auth/send_reset_password/send_reset_password.go
package send_reset_password

import (
	usecase "src/application/auth/send_reset_password"
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

	return c.SendStatus(fiber.StatusAccepted)
}

func operation(o *oas.Operation) {
	o.Tags("Auth").Summary("Send Reset Password").
		Description("Initiates the password recovery process by sending an email with a reset code.")
	router.BodyJson(o, router.SchemaAs[usecase.Data]())
	router.ResponseStatus(o, fiber.StatusAccepted, "Recovery email sent")
	router.ResponseIssueAs[*issue.AccountNotFound](o, fiber.StatusNotFound)
	router.ResponseIssueAs[*validate.ValidationError](o, fiber.StatusUnprocessableEntity)
}

var Route = router.
	Route(handler).
	Operation(operation)
