// port/broker/broker.go
package broker

import (
	"context"
	"src/domain"
)

// Publisher represents a message publisher
type Publisher[TKind ~string] interface {
	// Publish publishes a message
	Publish(ctx context.Context, event domain.Event[TKind]) error

	// Close closes the connection
	Close() error
}

// Consumer represents a message consumer
type Consumer[TKind ~string] interface {
	// Subscribe subscribes to topics
	Subscribe(topics ...TKind) error

	// Consume consumes messages
	Consume(ctx context.Context, handler func(ctx context.Context, event domain.Event[TKind]) error) error

	// Close closes the connection
	Close() error
}

// Client represents a broker client
type Client[TKind ~string] interface {
	Publisher[TKind]
	Consumer[TKind]

	// Connect establishes connection
	Connect(ctx context.Context) error

	// Ping checks connection health
	Ping(ctx context.Context) error

	// Close closes the connection
	Close() error
}
