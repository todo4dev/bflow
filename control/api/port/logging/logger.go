// port/logger/logger.go
package logger

import "context"

// Field represents a log field
type Field struct {
	Key   string
	Value any
}

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

// String creates a string field
func String(key, value string) Field {
	return Field{Key: key, Value: value}
}

// Int creates an integer field
func Int(key string, value int) Field {
	return Field{Key: key, Value: value}
}

// Error creates an error field
func Error(err error) Field {
	return Field{Key: "error", Value: err.Error()}
}

// Any creates a field of any type
func Any(key string, value any) Field {
	return Field{Key: key, Value: value}
}
