// presentation/http/resource/auth/check_email_available/route.go
package check_email_available

import (
	"net/url"
	usecase "src/application/auth/check_email_available"
	"src/domain/issue"
	"src/presentation/http/router"

	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/oas"
)

var Route = router.
	Route(func(c *fiber.Ctx) error {
		email, _ := url.QueryUnescape(c.Params("email"))
		data := &usecase.Data{Email: email}
		_, err := di.Resolve[*usecase.Handler]().Handle(c.Context(), data)
		if err != nil {
			return err
		}
		return c.SendStatus(fiber.StatusOK)
	}).
	Operation(func(o *oas.Operation) {
		o.Tags("Auth").Summary("Check Email Availability").
			Description("Verifies whether an email address is available for registration.")
		router.InPath(o, "email", func(s *oas.Schema) { s.String().Example("john.doe@email.com") })
		router.ResponseStatus(o, "200", "Email is available")
		router.ResponseIssue[*issue.AccountEmailInUse](o, "409")
	})
