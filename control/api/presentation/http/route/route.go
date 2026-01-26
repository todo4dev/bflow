// presentation/http/route/rest.go
package route

import (
	"src/presentation/http/route/auth"
	"src/presentation/http/route/health"
	"src/presentation/http/route/openapi_json"
	"src/presentation/http/route/openapi_ui"
	"src/presentation/http/server"
)

var Group = server.
	NewGroup("/", func(g *server.Grouper) {
		g.Group(auth.Group)
		g.Get(g.Config.GetOpenAPIPath(), openapi_ui.Route)
		g.Get(g.Config.GetOpenAPIJsonPath(), openapi_json.Route)
		g.Get("/health", health.Route)
	})
