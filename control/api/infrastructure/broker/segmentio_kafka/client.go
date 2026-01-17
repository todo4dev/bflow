// infrastructure/broker/segmentio_kafka/client.go
package segmentio_kafka

import (
	"context"
	"errors"
	"fmt"
	"src/port/broker"
	"time"

	"github.com/segmentio/kafka-go"
)

type SegmentioKafkaClient struct {
	config SegmentioKafkaConfig
	writer *kafka.Writer
	dialer *kafka.Dialer

	// Consumer state
	topics []string
	reader *kafka.Reader
}

var _ broker.Client = (*SegmentioKafkaClient)(nil)

func NewSegmentioKafkaClient(rawConfig SegmentioKafkaConfig) (*SegmentioKafkaClient, error) {
	config, err := SegmentioKafkaConfigSchema.Validate(rawConfig)
	if err != nil {
		return nil, err
	}

	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}

	return &SegmentioKafkaClient{
		config: config,
		dialer: dialer,
	}, nil
}

func (c *SegmentioKafkaClient) Connect(ctx context.Context) error {
	// Initialize writer shared instance
	c.writer = &kafka.Writer{
		Addr:     kafka.TCP(c.config.Brokers...),
		Balancer: &kafka.LeastBytes{},
	}
	// Verify connection
	return c.Ping(ctx)
}

func (c *SegmentioKafkaClient) Ping(ctx context.Context) error {
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

func (c *SegmentioKafkaClient) Close() error {
	if c.writer != nil {
		return c.writer.Close()
	}
	return nil
}

func (c *SegmentioKafkaClient) Subscribe(topics ...string) error {
	if c.reader != nil {
		return errors.New("consumer already started, cannot subscribe")
	}

	// Apply prefix
	prefixed := make([]string, len(topics))
	for i, t := range topics {
		prefixed[i] = c.config.TopicPrefix + t
	}

	c.topics = append(c.topics, prefixed...)
	return nil
}

func (c *SegmentioKafkaClient) Consume(ctx context.Context, handler func(ctx context.Context, msg broker.Message) error) error {
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
				// If context canceled, stop
				if errors.Is(err, context.Canceled) {
					return nil
				}
				// Log error? Or continue? For now continue/return
				// If connection lost, kafka-go retries automatically usually.
				// If fatal, return err.
				return err
			}

			// Process
			headers := make(map[string]string)
			for _, h := range m.Headers {
				headers[h.Key] = string(h.Value)
			}

			bMsg := broker.Message{
				Key:     string(m.Key),
				Value:   m.Value,
				Headers: headers,
			}

			if err := handler(ctx, bMsg); err != nil {
				// Handler failed?
				// Kafka-go AutoCommit assumes we process it.
				// If we want manual commit, we should use CommitMessages.
				// For now, FetchMessage + CommitMessages is safer than ReadMessage (auto-commit).
				// BUT ReadMessage is simpler for "at least once" if we don't crash.
				// Let's assume manual commit pattern for safety:
				// Fetch -> Process -> Commit
				// But interface doesn't expose commit control.
				// Returning error from handler implies failure to process.
				// We should NOT commit if handler failed.
				// Retry?
				// Simple approach: Log error and continue (lossy) or Retry loop?
				// Let's implement basics: Fetch -> Handle -> Commit if success.
				continue
			}

			if err := c.reader.CommitMessages(ctx, m); err != nil {
				return err
			}
		}
	}
}

func (c *SegmentioKafkaClient) Publish(ctx context.Context, topic string, key string, message []byte) error {
	msg := kafka.Message{
		Topic: c.config.TopicPrefix + topic,
		Key:   []byte(key),
		Value: message,
	}
	return c.writer.WriteMessages(ctx, msg)
}

func (c *SegmentioKafkaClient) PublishBatch(ctx context.Context, topic string, messages []broker.Message) error {
	msgs := make([]kafka.Message, len(messages))
	fullTopic := c.config.TopicPrefix + topic

	for i, m := range messages {
		headers := make([]kafka.Header, 0, len(m.Headers))
		for k, v := range m.Headers {
			headers = append(headers, kafka.Header{Key: k, Value: []byte(v)})
		}

		msgs[i] = kafka.Message{
			Topic:   fullTopic,
			Key:     []byte(m.Key),
			Value:   m.Value,
			Headers: headers,
		}
	}

	return c.writer.WriteMessages(ctx, msgs...)
}
