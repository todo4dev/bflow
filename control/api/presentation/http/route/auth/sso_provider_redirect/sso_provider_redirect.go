// presentation/http/route/auth/sso_provider_redirect/sso_provider_redirect.go
package sso_provider_redirect

import (
	"net/url"
	usecase "src/application/auth/sso_provider_redirect"
	"src/presentation/http/server"

	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/oas"
)

var Route = server.
	NewRoute(func(c *server.Context) error {
		callbackURL, err := url.QueryUnescape(c.Query("callback_url"))
		if err != nil {
			return fiber.NewError(fiber.StatusUnprocessableEntity, "invalid query parameter")
		}

		data := usecase.Data{Provider: c.Params("provider"), CallbackURL: callbackURL}
		result, err := di.Resolve[*usecase.Handler]().Handle(c.Context(), &data)
		if err != nil {
			return err
		}

		return c.Redirect(result.RedirectURL)
	}).
	Operation(func(o *oas.Operation) {
		o.Tags("Auth").Summary("SSO Provider Redirect").
			Description("Redirects the user to the SSO provider authentication page.")
		server.InPath(o, "provider", func(s *oas.Schema) {
			s.Description("The SSO provider (google or microsoft)").
				Enum("google", "microsoft").Example("google")
		})
		server.InQuery(o, "callback_url", func(s *oas.Schema) {
			s.Description("The URL to return to after authentication").
				String().Example("http://localhost:3000")
		})
		server.ResponseStatus(o, fiber.StatusFound, "Redirect to SSO Provider")
	})
