package activate_account

import (
	usecase "src/application/auth/activate_account"
	"src/domain/issue"
	"src/presentation/http/router"

	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/oas"
)

var Route = router.
	Route(func(c *fiber.Ctx) error {
		var data usecase.Data
		if err := c.BodyParser(&data); err != nil {
			return c.SendStatus(fiber.StatusUnprocessableEntity)
		}
		_, err := di.Resolve[*usecase.Handler]().Handle(c.UserContext(), &data)
		if err != nil {
			return err
		}
		return c.SendStatus(fiber.StatusOK)
	}).
	Operation(func(o *oas.Operation) {
		o.Tags("Auth").Summary("Activate account").
			Description("Activate account using OTP sent by email")
		router.BodyJson(o, func(s *oas.Schema) {
			s.Object().
				Property("email", func(p *oas.Schema) { p.String().Format("email") }).
				Property("otp", func(p *oas.Schema) { p.String() })
		})
		router.ResponseStatus(o, "200", "Account activated")
		router.ResponseValidationError(o)
		router.ResponseIssue[*issue.AccountNotFound](o, "404")
		router.ResponseIssue[*issue.AccountAlreadyActivated](o, "406")
		router.ResponseIssue[*issue.AccountInvalidOTP](o, "409")
	})
