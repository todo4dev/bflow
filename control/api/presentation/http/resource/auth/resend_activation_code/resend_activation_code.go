// presentation/http/resource/auth/resend_activation_code/resend_activation_code.go
package resend_activation_code

import (
	usecase "src/application/auth/resend_activation_code"
	"src/domain/issue"
	"src/presentation/http/router"

	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/oas"
)

func handler(c *fiber.Ctx) error {
	var data usecase.Data
	if err := c.BodyParser(&data); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	_, err := di.Resolve[*usecase.Handler]().Handle(c.Context(), &data)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusAccepted)
}

func operation(o *oas.Operation) {
	o.Tags("Auth").Summary("Resend Activation Code").
		Description("Resends the activation code to the user's email address.")
	router.BodyJson(o, router.SchemaAs[usecase.Data]())
	router.ResponseStatus(o, fiber.StatusAccepted, "Activation code sent")
	router.ResponseIssueAs[*issue.AccountNotFound](o, fiber.StatusNotFound)
	router.ResponseIssueAs[*issue.AccountAlreadyActivated](o, fiber.StatusNotAcceptable)
}

var Route = router.
	Route(handler).
	Operation(operation)
