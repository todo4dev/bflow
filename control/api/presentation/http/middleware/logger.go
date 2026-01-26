// presentation/http/middleware/logger.go
package middleware

import (
	"fmt"
	"src/port/logger"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logger(loggerClient logger.Client) fiber.Handler {
	skipPaths := map[string]bool{
		"/":             true,
		"/openapi.json": true,
		"/health":       true,
	}

	return func(c *fiber.Ctx) error {
		if skipPaths[c.Path()] {
			return c.Next()
		}

		start, err := time.Now(), c.Next()
		duration := time.Since(start)
		status := c.Response().StatusCode()

		if status >= 200 && status < 300 {
			return err
		}

		ctx := c.UserContext()

		fields := []logger.Field{
			logger.String("method", c.Method()),
			logger.String("path", c.Path()),
			logger.Int("status", status),
			logger.Any("duration_ms", duration.Milliseconds()),
			logger.String("ip", c.IP()),
			logger.String("user_agent", c.Get("User-Agent")),
		}

		if len(c.Request().URI().QueryString()) > 0 {
			fields = append(fields, logger.String("query", string(c.Request().URI().QueryString())))
		}

		if err != nil {
			fields = append(fields, logger.Error(err))
			loggerClient.Error(ctx, "HTTP request failed", err, fields...)
			return err
		}

		message := fmt.Sprintf("%s %s", c.Method(), c.Path())

		if status >= 500 {
			loggerClient.Error(ctx, message, nil, fields...)
		} else if status >= 400 {
			loggerClient.Warn(ctx, message, fields...)
		} else {
			loggerClient.Info(ctx, message, fields...)
		}

		return nil
	}
}
