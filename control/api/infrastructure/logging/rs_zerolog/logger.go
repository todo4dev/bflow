// infrastructure/logging/rs_zerolog/logger.go
package rs_zerolog

import (
	"context"
	"io"
	"os"
	"src/port/logging"

	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
}

var _ logging.Logger = (*Logger)(nil)

func NewLogger(config *Config) (*Logger, error) {
	level, err := zerolog.ParseLevel(config.Level)
	if err != nil {
		return nil, err
	}

	// Default to stdout
	var output io.Writer = os.Stdout

	zLogger := zerolog.New(output).
		Level(level).
		With().
		Timestamp().
		Str("service", config.ServiceName).
		Logger()

	return &Logger{
		logger: zLogger,
	}, nil
}

func (l *Logger) toZerologFields(fields []logging.Field) map[string]any {
	m := make(map[string]any, len(fields))
	for _, f := range fields {
		m[f.Key] = f.Value
	}
	return m
}

func (l *Logger) Debug(ctx context.Context, msg string, fields ...logging.Field) {
	l.logger.Debug().Fields(l.toZerologFields(fields)).Msg(msg)
}

func (l *Logger) Info(ctx context.Context, msg string, fields ...logging.Field) {
	l.logger.Info().Fields(l.toZerologFields(fields)).Msg(msg)
}

func (l *Logger) Warn(ctx context.Context, msg string, fields ...logging.Field) {
	sentry.CaptureMessage(msg)
	l.logger.Warn().Fields(l.toZerologFields(fields)).Msg(msg)
}

func (l *Logger) Error(ctx context.Context, msg string, err error, fields ...logging.Field) {
	sentry.CaptureException(err)
	l.logger.Error().Err(err).Fields(l.toZerologFields(fields)).Msg(msg)
}

func (l *Logger) Fatal(ctx context.Context, msg string, err error, fields ...logging.Field) {
	sentry.CaptureException(err)
	l.logger.Fatal().Err(err).Fields(l.toZerologFields(fields)).Msg(msg)
}

func (l *Logger) With(fields ...logging.Field) logging.Logger {
	newLogger := l.logger.With().Fields(l.toZerologFields(fields)).Logger()
	return &Logger{logger: newLogger}
}
