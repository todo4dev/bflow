package swagger_ui

import (
	"fmt"
	"path"
	"src/presentation/http/router"

	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/di"
)

func handler(c *fiber.Ctx) error {
	config := di.Resolve[*router.Config]()
	swaggerUIPath := path.Clean(config.BasePath + config.SwaggerPath)
	swaggerJSONPath := path.Clean(swaggerUIPath + "/openapi.json")

	html := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
			<head>
				<meta charset="utf-8">
				<title>%s</title>
				<link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist/swagger-ui.css" />
			</head>
			<body>
				<div id="swagger-ui"></div>
				<script src="https://unpkg.com/swagger-ui-dist/swagger-ui-bundle.js"></script>
				<script src="https://unpkg.com/swagger-ui-dist/swagger-ui-standalone-preset.js"></script>
				<script>
					window.onload = function() {
						window.ui = SwaggerUIBundle({
							url: %q,
							dom_id: '#swagger-ui',
							presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
							layout: "BaseLayout",
						});
					};
				</script>
			</body>
		</html>`,
		config.SwaggerTitle,
		swaggerJSONPath,
	)

	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(html)
}

var Route = router.Route(handler)
