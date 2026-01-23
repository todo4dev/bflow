// presentation/http/rest/auth/check_email_available/route.go
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

func handler(c *fiber.Ctx) error {
	email, _ := url.QueryUnescape(c.Params("email"))
	data := &usecase.Data{Email: email}
	if err := di.Resolve[*usecase.Handler]().Handle(c.Context(), data); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}

func operation(o *oas.Operation) {
	o.Tags("Auth").Summary("Check Email Availability").
		Description("Verifies whether an email address is available for registration.")
	router.InPath(o, "email", func(s *oas.Schema) { s.String().Example("john.doe@email.com") })
	router.ResponseStatus(o, fiber.StatusOK, "Email is available")
	router.ResponseIssueAs[*issue.AccountEmailInUse](o, fiber.StatusConflict)
}

var Route = router.
	Route(handler).
	Operation(operation)
