// infrastructure/broker/segmentio_kafka/config.go
package segmentio_kafka

import v "github.com/leandroluk/gox/validate"

type Config struct {
	Brokers         []string
	TopicPrefix     string
	ConsumerGroupID string
}

var ConfigSchema = v.Object(func(t *Config, s *v.ObjectSchema[Config]) {
	s.Field(&t.Brokers).Array(v.Text()).Required().Min(1)
	s.Field(&t.TopicPrefix).Text().Default("")
	s.Field(&t.ConsumerGroupID).Text().Required()
})
