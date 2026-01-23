// presentation/http/middleware/logger.go
package logger

import (
	"fmt"
	"time"

	"src/port/logging"

	"github.com/gofiber/fiber/v2"
)

type LoggerConfig struct {
	Logger logging.Logger

	SkipPaths []string

	SkipSuccessStatus bool
}

func LoggerHandler(config LoggerConfig) fiber.Handler {
	skipPaths := make(map[string]bool)
	for _, path := range config.SkipPaths {
		skipPaths[path] = true
	}

	return func(c *fiber.Ctx) error {

		if skipPaths[c.Path()] {
			return c.Next()
		}

		start := time.Now()

		err := c.Next()

		duration := time.Since(start)
		status := c.Response().StatusCode()

		if config.SkipSuccessStatus && status >= 200 && status < 300 {
			return err
		}

		ctx := c.UserContext()

		fields := []logging.Field{
			logging.String("method", c.Method()),
			logging.String("path", c.Path()),
			logging.Int("status", status),
			logging.Any("duration_ms", duration.Milliseconds()),
			logging.String("ip", c.IP()),
			logging.String("user_agent", c.Get("User-Agent")),
		}

		if len(c.Request().URI().QueryString()) > 0 {
			fields = append(fields, logging.String("query", string(c.Request().URI().QueryString())))
		}

		if err != nil {
			fields = append(fields, logging.Error(err))
			config.Logger.Error(ctx, "HTTP request failed", err, fields...)
			return err
		}

		message := fmt.Sprintf("%s %s", c.Method(), c.Path())

		if status >= 500 {
			config.Logger.Error(ctx, message, nil, fields...)
		} else if status >= 400 {
			config.Logger.Warn(ctx, message, fields...)
		} else {
			config.Logger.Info(ctx, message, fields...)
		}

		return nil
	}
}
