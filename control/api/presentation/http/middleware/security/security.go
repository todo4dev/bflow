// presentation/http/middleware/security/security.go
package security

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func New(optionalConfig ...helmet.Config) fiber.Handler {
	config := helmet.Config{
		XSSProtection:      "1; mode=block",
		ContentTypeNosniff: "nosniff",
		XFrameOptions:      "SAMEORIGIN",
		HSTSMaxAge:         31536000,
		// CSP ajustado para permitir o Swagger UI
		ContentSecurityPolicy: "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; img-src 'self' data: https://validator.swagger.io",
		ReferrerPolicy:        "no-referrer",
		PermissionPolicy:      "geolocation=(self)",
	}

	if len(optionalConfig) > 0 {
		config = optionalConfig[0]
	}

	helmetHandler := helmet.New(config)

	return func(ctx *fiber.Ctx) error {
		if strings.HasPrefix(ctx.Path(), "/_system") {
			return ctx.Next()
		}
		return helmetHandler(ctx)
	}
}
