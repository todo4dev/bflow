// presentation/http/server.go
package http

import (
	"src/presentation/http/route"
	"src/presentation/http/server"
)

func Provide() {
	server.Provide(route.Group)
}
