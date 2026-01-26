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

type RouteType struct {
	handlers  []Handler
	operation func(*oas.Operation)
}

func Route(handlers ...Handler) *RouteType {
	return &RouteType{handlers: handlers}
}

func (d *RouteType) Operation(spec func(*oas.Operation)) *RouteType {
	d.operation = spec
	return d
}
