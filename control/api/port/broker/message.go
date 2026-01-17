// port/broker/message.go
package broker

// Message represents a message
type Message struct {
	Key     string
	Value   []byte
	Headers map[string]string
}
