// presentation/http/route/auth/check_email_available/check_email_available.go
package check_email_available

import (
	"net/url"
	usecase "src/application/auth/check_email_available"
	"src/domain/issue"
	"src/presentation/http/server"

	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/oas"
	"github.com/leandroluk/gox/validate"
)

var Route = server.
	Route(func(c *server.Context) error {
		email, err := url.QueryUnescape(c.Params("email"))
		if err != nil {
			return fiber.NewError(fiber.StatusUnprocessableEntity, "invalid query parameter")
		}

		data := &usecase.Data{Email: email}
		if err := di.Resolve[*usecase.Handler]().Handle(c.Context(), data); err != nil {
			return err
		}
		return c.SendStatus(fiber.StatusOK)
	}).
	Operation(func(o *oas.Operation) {
		o.Tags("Auth").Summary("Check Email Availability").
			Description("Verifies whether an email address is available for registration.")
		server.InPath(o, "email", func(s *oas.Schema) { s.String().Example("john.doe@email.com") })
		server.ResponseStatus(o, fiber.StatusOK, "Email is available")
		server.ResponseIssueAs[*issue.AccountEmailInUse](o, fiber.StatusConflict)
		server.ResponseIssueAs[*validate.ValidationError](o, fiber.StatusUnprocessableEntity)
	})
