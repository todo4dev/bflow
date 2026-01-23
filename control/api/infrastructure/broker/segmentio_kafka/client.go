// infrastructure/broker/segmentio_kafka/client.go
package segmentio_kafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"src/domain"
	"src/port/broker"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

type Client struct {
	config *Config
	writer *kafka.Writer
	dialer *kafka.Dialer

	// Consumer state
	topics []string
	reader *kafka.Reader

	// Topic creation cache
	createdTopics map[string]bool
	topicMutex    sync.RWMutex
}

var _ broker.Client = (*Client)(nil)

func NewClient(config *Config) (*Client, error) {
	dialer := &kafka.Dialer{Timeout: 10 * time.Second, DualStack: true}
	client := &Client{
		config:        config,
		dialer:        dialer,
		createdTopics: make(map[string]bool),
		writer: &kafka.Writer{
			Addr:                   kafka.TCP(config.Brokers...),
			Balancer:               &kafka.LeastBytes{},
			AllowAutoTopicCreation: true,
			Transport:              &kafka.Transport{Dial: dialer.DialFunc},
		},
	}
	return client, nil
}

func (c *Client) Ping(ctx context.Context) error {
	if len(c.config.Brokers) == 0 {
		return fmt.Errorf("no brokers configured")
	}

	conn, err := c.dialer.DialContext(ctx, "tcp", c.config.Brokers[0])
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

func (c *Client) Close() error {
	if c.writer != nil {
		return c.writer.Close()
	}
	return nil
}

func (c *Client) Subscribe(topics ...string) error {
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

func (c *Client) Consume(ctx context.Context, handler func(ctx context.Context, event domain.Event) error) error {
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
				case "occurred_at":
					occurredAt, _ = time.Parse(time.RFC3339, string(h.Value))
				case "idempotency_key":
					idempotencyKey = string(h.Value)
				}
			}

			var payload any
			if err := json.Unmarshal(m.Value, &payload); err != nil {
				// Log error?
				continue
			}

			// Extract Kind from topic (remove prefix)
			topic := m.Topic
			kindStr := topic
			if len(c.config.TopicPrefix) > 0 && len(topic) > len(c.config.TopicPrefix) {
				kindStr = topic[len(c.config.TopicPrefix):]
			}

			event := domain.Event{
				Kind:           kindStr,
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

// ensureTopicExists creates the topic if it doesn't exist
func (c *Client) ensureTopicExists(ctx context.Context, topic string) error {
	// Check cache first
	c.topicMutex.RLock()
	if c.createdTopics[topic] {
		c.topicMutex.RUnlock()
		return nil
	}
	c.topicMutex.RUnlock()

	// Try to create topic
	c.topicMutex.Lock()
	defer c.topicMutex.Unlock()

	// Double-check after acquiring write lock
	if c.createdTopics[topic] {
		return nil
	}

	conn, err := c.dialer.DialLeader(ctx, "tcp", c.config.Brokers[0], topic, 0)
	if err != nil {
		// If we can't dial the leader, try to create the topic
		conn, err = c.dialer.DialContext(ctx, "tcp", c.config.Brokers[0])
		if err != nil {
			return fmt.Errorf("failed to connect to broker: %w", err)
		}
		defer conn.Close()

		// Create topic with default configuration
		topicConfigs := []kafka.TopicConfig{
			{
				Topic:             topic,
				NumPartitions:     1,
				ReplicationFactor: 1,
			},
		}

		err = conn.CreateTopics(topicConfigs...)
		if err != nil {
			// Ignore "topic already exists" errors
			var kafkaErr kafka.Error
			if errors.As(err, &kafkaErr) && kafkaErr.Title() == "Topic Already Exists" {
				c.createdTopics[topic] = true
				return nil
			}
			return fmt.Errorf("failed to create topic %s: %w", topic, err)
		}
	} else {
		conn.Close()
	}

	// Mark as created
	c.createdTopics[topic] = true
	return nil
}

func (c *Client) Publish(ctx context.Context, event domain.Event) error {
	topic := c.config.TopicPrefix + string(event.Kind)

	// Ensure topic exists before publishing
	if err := c.ensureTopicExists(ctx, topic); err != nil {
		return fmt.Errorf("failed to ensure topic exists: %w", err)
	}

	value, err := json.Marshal(event.Payload)
	if err != nil {
		return err
	}

	headers := []kafka.Header{
		{Key: "occurred_at", Value: []byte(event.OccurredAt.Format(time.RFC3339))},
		{Key: "idempotency_key", Value: []byte(event.IdempotencyKey)},
	}

	msg := kafka.Message{Topic: topic, Value: value, Headers: headers}
	return c.writer.WriteMessages(ctx, msg)
}
