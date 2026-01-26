package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"src/application"
	"src/infrastructure"
	"src/presentation"
	"src/presentation/http/server"

	"github.com/getsentry/sentry-go"
	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
	"github.com/leandroluk/gox/util"
)

func main() {
	env.Load(".env", "../.env")

	sentry.Init(sentry.ClientOptions{Dsn: env.Get("API_TRACE_URL", "")})
	defer sentry.Flush(2 * time.Second)

	infrastructure.Provide()
	application.Provide()
	presentation.Provide()

	app := di.Resolve[*server.Server]()

	errChan := make(chan error, 1)
	go func() {
		errChan <- app.Listen()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errChan:
		util.Check(err)
	case <-stop:
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		util.Check(app.Shutdown(ctx))
	}
}
