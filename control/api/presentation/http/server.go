package http

import (
	"src/presentation/http/resource/billing"
	"src/presentation/http/resource/deployment"
	"src/presentation/http/resource/identity"
	"src/presentation/http/resource/signing"
	"src/presentation/http/resource/system"
	"src/presentation/http/resource/tenant"
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

	r := router.Wrapper(a,
		billing.Group,
		deployment.Group,
		identity.Group,
		signing.Group,
		system.Group,
		tenant.Group,
	)

	return &Server{app: a, router: r}
}

func (s *Server) Listen(addr string) error {
	return s.router.Listen(addr)
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
