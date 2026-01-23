// presentation/http/rest/system/system.go
package system

import (
	"path"
	"src/presentation/http/rest/system/healthcheck"
	"src/presentation/http/rest/system/swagger_json"
	"src/presentation/http/rest/system/swagger_ui"
	"src/presentation/http/router"

	"github.com/leandroluk/gox/di"
)

var Group = router.Group("/_system", func(g *router.GroupRouter) {
	config := di.Resolve[*router.Config]()
	swaggerUIPath := path.Clean(config.BasePath + config.SwaggerPath)
	swaggerJSONPath := path.Clean(swaggerUIPath + "/openapi.json")

	g.Get(swaggerUIPath, swagger_ui.Route)
	g.Get(swaggerJSONPath, swagger_json.Route)
	g.Get("/health", healthcheck.Route)
})
