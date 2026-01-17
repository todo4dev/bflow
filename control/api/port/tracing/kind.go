// port/tracing/kind.go
package tracing

// Kind span type
type Kind string

// Kind values
const (
	Kind_PRODUCER Kind = "PRODUCER"
	Kind_CONSUMER Kind = "CONSUMER"
	Kind_INTERNAL Kind = "INTERNAL"
)
