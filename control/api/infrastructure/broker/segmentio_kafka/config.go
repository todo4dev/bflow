// infrastructure/broker/segmentio_kafka/config.go
package segmentio_kafka

import "github.com/leandroluk/gox/validate"

var configSchema = validate.Object(func(t *Config, s *validate.ObjectSchema[Config]) {
	s.Field(&t.Brokers).Array(validate.Text()).Required().Min(1)
	s.Field(&t.TopicPrefix).Text().Default("")
	s.Field(&t.ConsumerGroupID).Text().Required()
})

type Config struct {
	Brokers         []string
	TopicPrefix     string
	ConsumerGroupID string
}

func (c *Config) Validate() (err error) {
	_, err = configSchema.Validate(c)
	return err
}
