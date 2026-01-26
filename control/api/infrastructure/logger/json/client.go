// infrastructure/logger/json/logger.go
package json

import (
	"context"
	"io"
	"os"
	"src/port/logger"

	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
)

type Client struct {
	logger zerolog.Logger
}

var _ logger.Client = (*Client)(nil)

func New(config *Config) (*Client, error) {
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

	return &Client{
		logger: zLogger,
	}, nil
}

func (l *Client) toZerologFields(fields []logger.Field) map[string]any {
	m := make(map[string]any, len(fields))
	for _, f := range fields {
		m[f.Key] = f.Value
	}
	return m
}

func (l *Client) Debug(ctx context.Context, msg string, fields ...logger.Field) {
	l.logger.Debug().Fields(l.toZerologFields(fields)).Msg(msg)
}

func (l *Client) Info(ctx context.Context, msg string, fields ...logger.Field) {
	l.logger.Info().Fields(l.toZerologFields(fields)).Msg(msg)
}

func (l *Client) Warn(ctx context.Context, msg string, fields ...logger.Field) {
	sentry.CaptureMessage(msg)
	l.logger.Warn().Fields(l.toZerologFields(fields)).Msg(msg)
}

func (l *Client) Error(ctx context.Context, msg string, err error, fields ...logger.Field) {
	sentry.CaptureException(err)
	l.logger.Error().Err(err).Fields(l.toZerologFields(fields)).Msg(msg)
}

func (l *Client) Fatal(ctx context.Context, msg string, err error, fields ...logger.Field) {
	sentry.CaptureException(err)
	l.logger.Fatal().Err(err).Fields(l.toZerologFields(fields)).Msg(msg)
}

func (l *Client) With(fields ...logger.Field) logger.Client {
	newLogger := l.logger.With().Fields(l.toZerologFields(fields)).Logger()
	return &Client{logger: newLogger}
}
