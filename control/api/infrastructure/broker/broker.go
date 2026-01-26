// infrastructure/broker/broker.go
package broker

import (
	"fmt"
	"src/infrastructure/broker/kafka"

	"github.com/leandroluk/gox/env"
)

func Provide() {
	provider := env.Get("API_BROKER_PROVIDER", "kafka")
	switch provider {
	case "kafka":
		kafka.Provide()
	default:
		panic(fmt.Errorf("invalid 'API_BROKER_PROVIDER': %s", provider))
	}
}
