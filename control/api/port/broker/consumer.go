// port/broker/consumer.go
package broker

import "context"

// Consumer represents a message consumer
type Consumer interface {
	// Subscribe subscribes to topics
	Subscribe(topics ...string) error

	// Consume consumes messages
	Consume(ctx context.Context, handler func(ctx context.Context, msg Message) error) error

	// Close closes the connection
	Close() error
}
