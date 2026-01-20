// presentation/http/route/system/system.go
package system

import (
	"path"
	"src/presentation/http/route/system/healthcheck"
	"src/presentation/http/route/system/swagger_json"
	"src/presentation/http/route/system/swagger_ui"
	"src/presentation/http/router"

	"github.com/leandroluk/gox/di"
)

var Group = router.Group("/", func(g *router.GroupRouter) {
	config := di.Resolve[*router.Config]()
	swaggerUIPath := path.Clean(config.BasePath + config.SwaggerPath)
	swaggerJSONPath := path.Clean(swaggerUIPath + "/openapi.json")

	g.Get("/health", healthcheck.Route)
	g.Get(swaggerUIPath, swagger_ui.Route)
	g.Get(swaggerJSONPath, swagger_json.Route)
})
