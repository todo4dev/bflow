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
		if _, err := segmentio_kafka.ConfigSchema.Validate(&config); err != nil {
			panic(err)
		}

		instance, err := segmentio_kafka.NewClient(&config)
		if err != nil {
			panic(err)
		}
		di.SingletonAs[broker.Client](func() broker.Client { return instance })
	// case "mocking_impl":
	// case "another_broker_impl":
	default:
		panic(fmt.Errorf("invalid 'BROKER_PROVIDER': %s", provider))
	}
}
