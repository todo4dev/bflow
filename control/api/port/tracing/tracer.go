// port/tracing/tracer.go
package tracing

import "context"

// Kind span type
type Kind string

// Kind values
const (
	Kind_PRODUCER Kind = "PRODUCER"
	Kind_CONSUMER Kind = "CONSUMER"
	Kind_INTERNAL Kind = "INTERNAL"
)

// Definition span configuration
type Definition struct {
	Attributes map[string]any
	Kind       Kind
}

// Span represents a span of trace
type Span interface {
	// End ends the span
	End()

	// SetAttribute defines attribute
	SetAttribute(key string, value any)

	// SetError marks span as error
	SetError(err error)

	// AddEvent adds event to span
	AddEvent(name string, attributes map[string]any)
}

// Tracer represents a distributed tracing tracer
type Tracer interface {
	// StartSpan starts a new span
	StartSpan(ctx context.Context, name string, defs ...func(*Definition)) (context.Context, Span)

	// Extract extracts context from headers
	Extract(ctx context.Context, carrier map[string]string) context.Context

	// Inject injects context into headers
	Inject(ctx context.Context, carrier map[string]string)
}
