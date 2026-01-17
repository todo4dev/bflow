// port/broker/broker.go
package broker

import "context"

// Message represents a message
type Message struct {
	Key     string
	Value   []byte
	Headers map[string]string
}

// Publisher represents a message publisher
type Publisher interface {
	// Publish publishes a message to a topic
	Publish(ctx context.Context, topic string, key string, message []byte) error

	// PublishBatch publishes multiple messages
	PublishBatch(ctx context.Context, topic string, messages []Message) error

	// Close closes the connection
	Close() error
}

// Consumer represents a message consumer
type Consumer interface {
	// Subscribe subscribes to topics
	Subscribe(topics ...string) error

	// Consume consumes messages
	Consume(ctx context.Context, handler func(ctx context.Context, msg Message) error) error

	// Close closes the connection
	Close() error
}

// Client represents a broker client
type Client interface {
	Publisher
	Consumer

	// Connect establishes connection
	Connect(ctx context.Context) error

	// Ping checks connection health
	Ping(ctx context.Context) error

	// Close closes the connection
	Close() error
}
