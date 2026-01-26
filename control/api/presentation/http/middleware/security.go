// presentation/http/middleware/security.go
package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func Security(ignoredRoute ...string) fiber.Handler {
	config := helmet.Config{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "SAMEORIGIN",
		HSTSMaxAge:            31536000,
		ContentSecurityPolicy: "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; img-src 'self' data: https://validator.swagger.io",
		ReferrerPolicy:        "no-referrer",
		PermissionPolicy:      "geolocation=(self)",
	}

	helmetHandler := helmet.New(config)

	return func(ctx *fiber.Ctx) error {
		for _, route := range ignoredRoute {
			if strings.HasPrefix(ctx.Path(), route) {
				return ctx.Next()
			}
		}
		return helmetHandler(ctx)
	}
}
