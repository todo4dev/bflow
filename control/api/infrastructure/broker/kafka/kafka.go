// infrastructure/broker/kafka/kafka.go
package kafka

import (
	"fmt"
	"src/port/broker"
	"strings"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

func Provide() {
	config := Config{
		Brokers:         strings.Split(env.Get("API_BROKER_KAFKA_URL", "localhost:9092"), ","),
		TopicPrefix:     env.Get("API_BROKER_KAFKA_TOPIC_PREFIX", ""),
		ConsumerGroupID: env.Get("API_BROKER_KAFKA_CONSUMER_GROUP_ID", ""),
	}
	if err := config.Validate(); err != nil {
		panic(fmt.Errorf("broker config validation failed: %w", err))
	}

	instance, err := New(&config)
	if err != nil {
		panic(fmt.Errorf("failed to create broker client: %w", err))
	}

	di.SingletonAs[broker.Client](func() broker.Client { return instance })
}
