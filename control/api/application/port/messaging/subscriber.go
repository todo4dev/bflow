// control/api/application/port/messaging/subscriber.go
package messaging

type Subscriber interface {
	Subscribe(topic string, handler func(payload []byte) error) error
	Unsubscribe(topic string) error
}
