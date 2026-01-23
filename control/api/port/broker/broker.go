// port/broker/broker.go
package broker

import (
	"context"
	"src/domain"
)

// Client represents a broker client
type Client interface {
	// Publish publishes a message
	Publish(ctx context.Context, event domain.Event) error

	// Subscribe subscribes to topics
	Subscribe(topics ...string) error

	// Consume consumes messages
	Consume(ctx context.Context, handler func(ctx context.Context, event domain.Event) error) error

	// Ping checks connection health
	Ping(ctx context.Context) error

	// Close closes the connection
	Close() error
}
