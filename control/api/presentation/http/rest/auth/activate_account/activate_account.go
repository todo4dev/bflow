package activate_account

import (
	usecase "src/application/auth/activate_account"
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
	if err := di.Resolve[*usecase.Handler]().Handle(c.UserContext(), &data); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}

func operation(o *oas.Operation) {
	o.Tags("Auth").Summary("Activate account").
		Description("Activate account using OTP sent by email")
	router.BodyJson(o, router.SchemaAs[usecase.Data]())
	router.ResponseStatus(o, fiber.StatusOK, "Account activated")
	router.ResponseIssueAs[*validate.ValidationError](o, fiber.StatusBadRequest)
	router.ResponseIssueAs[*issue.AccountNotFound](o, fiber.StatusNotFound)
	router.ResponseIssueAs[*issue.AccountAlreadyActivated](o, fiber.StatusNotAcceptable)
	router.ResponseIssueAs[*issue.AccountInvalidOTP](o, fiber.StatusConflict)
	router.ResponseStatus(o, fiber.StatusUnprocessableEntity, "Unprocessable entity")
}

var Route = router.
	Route(handler).
	Operation(operation)
