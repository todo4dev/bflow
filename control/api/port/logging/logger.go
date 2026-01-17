// port/logging/logger.go
package logging

import "context"

// Logger represents a structured logger (JSON)
type Logger interface {
	// Debug registers a debug message
	Debug(ctx context.Context, msg string, fields ...Field)

	// Info registers an informational message
	Info(ctx context.Context, msg string, fields ...Field)

	// Warn registers a warning message
	Warn(ctx context.Context, msg string, fields ...Field)

	// Error registers an error message
	Error(ctx context.Context, msg string, err error, fields ...Field)

	// Fatal registers a fatal error and terminates the application
	Fatal(ctx context.Context, msg string, err error, fields ...Field)

	// With creates a child logger with additional fields
	With(fields ...Field) Logger
}
