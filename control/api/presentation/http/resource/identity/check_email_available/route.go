package check_email_available

import (
	"net/url"
	"src/application/identity/usecase/check_email_available"
	"src/presentation/http/router"

	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/oas"
)

var Route = router.
	Route(func(c *fiber.Ctx) error {
		useCase := di.Resolve[*check_email_available.Handler]()
		email, err := url.QueryUnescape(c.Params("email"))
		if err != nil {
			return err
		}
		if _, err := useCase.Handle(c.Context(), &check_email_available.Data{Email: email}); err != nil {
			return err
		}
		return c.SendStatus(fiber.StatusOK)
	}).
	Operation(func(o *oas.Operation) {
		o.Tags("Identity").Summary("Check Email Availability").
			Description("Verifies whether an email address is available for registration.").
			Parameter(func(p *oas.Parameter) {
				p.Name("email").In("path").Required(true).
					Description("The email address to check").
					Schema(func(s *oas.Schema) { s.String() })
			}).
			Response("200", func(r *oas.Response) {
				r.Description("Email is available")
			}).
			Response("409", func(r *oas.Response) {
				r.Description("Email is already in use").
					Json(func(m *oas.MediaType) {
						m.Schema(func(s *oas.Schema) {
							s.Object().
								Required("code", func(p *oas.Schema) { p.String().Example("EMAIL_ALREADY_IN_USE") }).
								Required("message", func(p *oas.Schema) { p.String().Example("The provided email is already associated with an account.") })
						})
					})
			})
	})
