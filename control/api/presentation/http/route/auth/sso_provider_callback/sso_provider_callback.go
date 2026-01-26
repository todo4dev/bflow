// presentation/http/route/auth/sso_provider_callback/sso_provider_callback.go
package sso_provider_callback

import (
	usecase "src/application/auth/sso_provider_callback"
	"src/presentation/http/server"

	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/oas"
)

var Route = server.
	Route(func(c *server.Context) error {
		provider := c.Params("provider")
		code := c.Query("code")
		state := c.Query("state")
		errCode := c.Query("error")

		var err *string
		if errCode != "" {
			err = &errCode
		}

		data := usecase.Data{
			Provider: provider,
			Code:     code,
			State:    state,
			Error:    err,
		}

		result, appErr := di.Resolve[*usecase.Handler]().Handle(c.Context(), &data)
		if appErr != nil {
			return appErr
		}

		return c.Redirect(result.RedirectURL, fiber.StatusTemporaryRedirect)
	}).
	Operation(func(o *oas.Operation) {
		o.Tags("Auth").Summary("SSO Provider Callback").
			Description("Callback endpoint for SSO providers to return authorization code.")
		server.InPath(o, "provider", func(s *oas.Schema) {
			s.Description("The SSO provider (google or microsoft)").
				Enum("google", "microsoft").Example("google")
		})
		server.InQuery(o, "code", func(s *oas.Schema) {
			s.Description("Authorization code from the provider").
				String().Example("auth_code_xyz")
		})
		server.InQuery(o, "state", func(s *oas.Schema) {
			s.Description("State parameter containing encrypted callback URL").
				String().Example("ey...")
		})
		server.InQuery(o, "error", func(s *oas.Schema) {
			s.Description("Error code from provider if any").
				String().Example("access_denied")
		})
		server.ResponseStatus(o, fiber.StatusTemporaryRedirect, "Redirect to Client with Code or Error")
	})
