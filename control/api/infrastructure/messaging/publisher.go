// application/port/messaging/publisher.go
package messaging

type Publisher interface {
	Publish(topic string, payload []byte) error
}
