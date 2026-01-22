// infrastructure/broker/broker.go
package broker

import (
	"fmt"
	"src/infrastructure/broker/segmentio_kafka"
	"src/port/broker"
	"strings"

	"github.com/leandroluk/gox/di"
	"github.com/leandroluk/gox/env"
)

func Provide() {
	provider := env.Get("BROKER_PROVIDER", "segmentio_kafka")
	switch provider {
	case "segmentio_kafka":
		config := segmentio_kafka.Config{
			Brokers:         strings.Split(env.Get("BROKER_URL", "localhost:9092"), ","),
			TopicPrefix:     env.Get("BROKER_TOPIC_PREFIX", ""),
			ConsumerGroupID: env.Get("BROKER_CONSUMER_GROUP_ID", ""),
		}
		if err := config.Validate(); err != nil {
			panic(fmt.Errorf("broker config validation failed: %w", err))
		}

		instance, err := segmentio_kafka.NewClient[string](&config)
		if err != nil {
			panic(fmt.Errorf("failed to create broker client: %w", err))
		}

		di.SingletonGeneric[broker.Client[string]](instance)

	// case "mocking_impl":
	// case "another_broker_impl":
	default:
		panic(fmt.Errorf("invalid 'BROKER_PROVIDER': %s", provider))
	}
}
