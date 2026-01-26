package server

import (
	"context"
	"src/application/health"
	"src/port/logger"
	"src/presentation/http/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
	"github.com/leandroluk/gox/util"
)

type Server struct {
	app    *fiber.App
	logger logger.Client
	config *Config
	router *Router
}

func New(logger logger.Client, config *Config, group GroupType) *Server {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler:          middleware.ErrorHandler(),
	})
	app.
		Use(middleware.Recover()).
		Use(middleware.Logger(logger)).
		Use(middleware.Cors(config.AppOrigin)).
		Use(middleware.Compress()).
		Use(middleware.Security(
			config.GetOpenAPIPath(),
			config.GetOpenAPIJsonPath()))

	router := &Router{fiber: app, path: "", Config: config}
	router.Group(group)

	return &Server{
		app:    app,
		logger: logger,
		config: config,
		router: router,
	}
}

func (s *Server) Listen() error {
	util.Must(di.Resolve[*health.Handler]().Handle())
	s.logger.Info(context.Background(), "🚀 running on port "+s.config.AppPort)
	return s.app.Listen(":" + s.config.AppPort)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.app.ShutdownWithContext(ctx)
}

func Provide(group GroupType) {
	config, err := configSchema.Validate(Config{
		AppName:             env.Get("API_NAME", "App Name"),
		AppPort:             env.Get("API_PORT", "30000"),
		AppPath:             env.Get("API_PATH", "/"),
		AppOrigin:           env.Get("API_ORIGIN", "*"),
		OpenAPIEnable:       env.Get("API_OPENAPI_ENABLE", "true") == "true",
		OpenAPIUIPath:       env.Get("API_OPENAPI_UI_PATH", "/"),
		OpenAPIJSONPath:     env.Get("API_OPENAPI_JSON_PATH", "/openapi.json"),
		OpenAPITitle:        env.Get("API_OPENAPI_TITLE", "Title"),
		OpenAPIDescription:  env.Get("API_OPENAPI_DESCRIPTION", ""),
		OpenAPIContactName:  env.Get("API_OPENAPI_CONTACT_NAME", "Contact Name"),
		OpenAPIContactURL:   env.Get("API_OPENAPI_CONTACT_URL", "https://example.com"),
		OpenAPIContactEmail: env.Get("API_OPENAPI_CONTACT_EMAIL", "contact@email.com"),
		OpenAPILicenseName:  env.Get("API_OPENAPI_LICENSE_NAME", "License Name"),
		OpenAPILicenseURL:   env.Get("API_OPENAPI_LICENSE_URL", "https://example.com"),
		OpenAPIVersion:      env.Get("API_OPENAPI_VERSION", "1.0.0"),
	})
	if err != nil {
		panic(err)
	}
	di.SingletonAs[*Server](func(logger logger.Client) *Server {
		return New(logger, &config, group)
	})
}
