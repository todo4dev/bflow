// infrastructure/broker/segmentio_kafka/config.go
package segmentio_kafka

import "github.com/leandroluk/go/v"

type SegmentioKafkaConfig struct {
	Brokers         []string
	TopicPrefix     string
	ConsumerGroupID string
}

var SegmentioKafkaConfigSchema = v.Object(func(t *SegmentioKafkaConfig, s *v.ObjectSchema[SegmentioKafkaConfig]) {
	s.Field(&t.Brokers).Array(v.Text().URL()).Required().Min(1)
	s.Field(&t.TopicPrefix).Text().Default("")
	s.Field(&t.ConsumerGroupID).Text().Required()
})
