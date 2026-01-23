// presentation/http/server.go
package http

import (
	"context"
	"fmt"
	"src/port/logging"
	middleware "src/presentation/http/middleware"
	rest "src/presentation/http/rest"
	"src/presentation/http/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

type Server struct {
	app    *fiber.App
	router *router.Router
	logger logging.Logger
}

func NewServer() *Server {
	server := &Server{
		app: fiber.New(fiber.Config{
			DisableStartupMessage: true,
			ErrorHandler:          router.ErrorHandler(),
		}),
		logger: di.Resolve[logging.Logger](),
	}

	server.app.
		Use(recover.New(recover.Config{EnableStackTrace: true})).
		Use(middleware.LoggerHandler(middleware.LoggerConfig{
			Logger:    server.logger,
			SkipPaths: []string{"/system/health"},
		}))

	server.router = router.Wrapper(server.app, rest.Routes)

	return server
}

func (s *Server) Listen() error {
	port := env.Get("APP_PORT", "3000")
	name := env.Get("APP_NAME", "bflow-control")

	s.logger.Info(context.Background(), fmt.Sprintf("🚀 %s running on port :%s", name, port))
	return s.router.Listen(":" + port)
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
