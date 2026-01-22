// presentation/http/server.go
package http

import (
	"src/presentation/http/resource"
	"src/presentation/http/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	app    *fiber.App
	router *router.Router
}

func NewServer() *Server {
	a := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler:          errorHandler,
	})
	a.Use(recover.New())

	r := router.Wrapper(a, resource.Routes)

	return &Server{app: a, router: r}
}

func (s *Server) Listen(addr string) error {
	return s.router.Listen(addr)
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
