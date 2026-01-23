// presentation/http/server.go
package http

import (
	rest "src/presentation/http/rest"
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
		ErrorHandler:          router.ErrorHandler(),
	})
	a.Use(recover.New(recover.Config{EnableStackTrace: true}))

	r := router.Wrapper(a, rest.Routes)

	return &Server{app: a, router: r}
}

func (s *Server) Listen(addr string) error {
	return s.router.Listen(addr)
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
