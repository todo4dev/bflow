// presentation/http/server.go
package http

import (
	"context"
	"fmt"
	"src/port/logging"
	"src/presentation/http/middleware/compress"
	"src/presentation/http/middleware/cors"
	"src/presentation/http/middleware/logger"
	"src/presentation/http/middleware/recover"
	"src/presentation/http/middleware/security"
	"src/presentation/http/rest"
	"src/presentation/http/router"

	"github.com/gofiber/fiber/v2"
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
		Use(recover.New()).
		Use(logger.New(server.logger)).
		Use(cors.New(env.Get("API_ORIGIN", "*"))).
		Use(compress.New()).
		Use(security.New())

	server.router = router.Wrapper(server.app, rest.Routes)

	return server
}

func (s *Server) Listen() error {
	port := env.Get("API_PORT", "3000")
	name := env.Get("API_NAME", "bflow-control")

	s.logger.Info(context.Background(), fmt.Sprintf("🚀 %s running on port :%s", name, port))
	return s.router.Listen(":" + port)
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
