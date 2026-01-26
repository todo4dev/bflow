// presentation/http/server/route_definition.go
package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/oas"
)

type Context struct {
	*fiber.Ctx
	Config *Config
	Router *Router
}

type Handler func(*Context) error

type Route struct {
	handlers  []Handler
	operation func(*oas.Operation)
}

func NewRoute(handlers ...Handler) *Route {
	return &Route{handlers: handlers}
}

func (d *Route) Operation(spec func(*oas.Operation)) *Route {
	d.operation = spec
	return d
}
