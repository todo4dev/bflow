// presentation/http/route/openapi_ui/route.go
package openapi_ui

import (
	"fmt"
	"src/presentation/http/server"
)

var Route = server.
	Route(func(c *server.Context) error {
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
			c.Config.OpenAPITitle,
			c.Config.GetOpenAPIJsonPath(),
		)

		c.Set("Content-Type", "text/html; charset=utf-8")
		return c.SendString(html)
	})
