// port/broker/publisher.go
package broker

import "context"

// Publisher represents a message publisher
type Publisher interface {
	// Publish publishes a message to a topic
	Publish(ctx context.Context, topic string, key string, message []byte) error

	// PublishBatch publishes multiple messages
	PublishBatch(ctx context.Context, topic string, messages []Message) error

	// Close closes the connection
	Close() error
}
