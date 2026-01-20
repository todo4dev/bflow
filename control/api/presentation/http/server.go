package http

import (
	"src/presentation/http/route/billing"
	"src/presentation/http/route/deployment"
	"src/presentation/http/route/identity"
	"src/presentation/http/route/signing"
	"src/presentation/http/route/system"
	"src/presentation/http/route/tenant"
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

	r := router.Wrapper(a).
		Group(billing.Group).
		Group(deployment.Group).
		Group(identity.Group).
		Group(signing.Group).
		Group(system.Group).
		Group(tenant.Group)

	return &Server{app: a, router: r}
}

func (s *Server) Listen(addr string) error {
	return s.router.Listen(addr)
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
