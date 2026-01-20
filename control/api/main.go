// main.go
package main

import (
	"context"
	"time"

	"src/application"
	"src/application/system/usecase/healthcheck"
	"src/infrastructure"
	"src/presentation/http"

	"github.com/getsentry/sentry-go"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

func main() {
	env.Load(".env", "../.env")

	sentry.Init(sentry.ClientOptions{Dsn: env.Get("APP_TRACE_URL", "")})
	defer sentry.Flush(2 * time.Second)

	infrastructure.Provide()
	application.Provide()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	handler := di.Resolve[*healthcheck.Handler]()
	if _, err := handler.Handle(ctx); err != nil {
		panic(err)
	}

	server := http.NewServer()

	port := env.Get("APP_PORT", "3000")
	name := env.Get("APP_NAME", "bflow-control")

	println("🚀", name, "running on port :"+port)

	if err := server.Listen(":" + port); err != nil {
		panic(err)
	}
}
