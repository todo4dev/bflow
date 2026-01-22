// infrastructure/broker/segmentio_kafka/client.go
package segmentio_kafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"src/domain"
	"src/port/broker"
	"time"

	"github.com/segmentio/kafka-go"
)

type Client[TKind ~string] struct {
	config *Config
	writer *kafka.Writer
	dialer *kafka.Dialer

	// Consumer state
	topics []string
	reader *kafka.Reader
}

var _ broker.Client[string] = (*Client[string])(nil)

func NewClient[TKind ~string](config *Config) (*Client[TKind], error) {
	dialer := &kafka.Dialer{Timeout: 10 * time.Second, DualStack: true}
	client := &Client[TKind]{config: config, dialer: dialer}
	return client, nil
}

func (c *Client[TKind]) Connect(ctx context.Context) error {
	// Initialize writer shared instance
	c.writer = &kafka.Writer{
		Addr:     kafka.TCP(c.config.Brokers...),
		Balancer: &kafka.LeastBytes{},
	}
	// Verify connection
	return c.Ping(ctx)
}

func (c *Client[TKind]) Ping(ctx context.Context) error {
	if len(c.config.Brokers) == 0 {
		return fmt.Errorf("no brokers configured")
	}
	// Try to dial the first broker to check connectivity
	conn, err := c.dialer.DialContext(ctx, "tcp", c.config.Brokers[0])
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

func (c *Client[TKind]) Close() error {
	if c.writer != nil {
		return c.writer.Close()
	}
	return nil
}

func (c *Client[TKind]) Subscribe(topics ...TKind) error {
	if c.reader != nil {
		return errors.New("consumer already started, cannot subscribe")
	}

	// Apply prefix
	prefixed := make([]string, len(topics))
	for i, t := range topics {
		prefixed[i] = c.config.TopicPrefix + string(t)
	}

	c.topics = append(c.topics, prefixed...)
	return nil
}

func (c *Client[TKind]) Consume(ctx context.Context, handler func(ctx context.Context, event domain.Event[TKind]) error) error {
	if len(c.topics) == 0 {
		return errors.New("no topics subscribed")
	}

	c.reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:     c.config.Brokers,
		GroupTopics: c.topics,
		GroupID:     c.config.ConsumerGroupID,
		Dialer:      c.dialer,
		MinBytes:    10e3, // 10KB
		MaxBytes:    10e6, // 10MB
		MaxWait:     1 * time.Second,
	})

	defer func() {
		_ = c.reader.Close()
		c.reader = nil
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			m, err := c.reader.FetchMessage(ctx)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return nil
				}
				return err
			}

			// Reconstruct domain.Event
			var occurredAt time.Time
			var idempotencyKey string

			for _, h := range m.Headers {
				switch h.Key {
				case "OccurredAt":
					occurredAt, _ = time.Parse(time.RFC3339, string(h.Value))
				case "IdempotencyKey":
					idempotencyKey = string(h.Value)
				}
			}

			var payload any
			if err := json.Unmarshal(m.Value, &payload); err != nil {
				// Log error?
				continue
			}

			// Extract Kind from topic (remove prefix)
			// Assuming prefix length is fixed or we can strip it.
			// Currently prefix is c.config.TopicPrefix.
			topic := m.Topic
			kindStr := topic
			if len(c.config.TopicPrefix) > 0 && len(topic) > len(c.config.TopicPrefix) {
				kindStr = topic[len(c.config.TopicPrefix):]
			}

			event := domain.Event[TKind]{
				Kind:           TKind(kindStr),
				IdempotencyKey: idempotencyKey,
				OccurredAt:     occurredAt,
				Payload:        payload,
			}

			if err := handler(ctx, event); err != nil {
				continue
			}

			if err := c.reader.CommitMessages(ctx, m); err != nil {
				return err
			}
		}
	}
}

func (c *Client[TKind]) Publish(ctx context.Context, event domain.Event[TKind]) error {
	b, err := json.Marshal(event.Payload)
	if err != nil {
		return err
	}

	headers := []kafka.Header{
		{Key: "OccurredAt", Value: []byte(event.OccurredAt.Format(time.RFC3339))},
		{Key: "IdempotencyKey", Value: []byte(event.IdempotencyKey)},
	}

	msg := kafka.Message{
		Topic:   c.config.TopicPrefix + string(event.Kind),
		Value:   b,
		Headers: headers,
	}
	return c.writer.WriteMessages(ctx, msg)
}
