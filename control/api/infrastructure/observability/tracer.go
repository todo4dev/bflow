// application/port/observability/tracer.go
package observability

import "context"

type Attribute struct {
	Key   string
	Value any
}

type Span interface {
	End()
	AddEvent(name string, attributes ...Attribute)
	RecordError(err error, attributes ...Attribute)
	SetAttributes(attributes ...Attribute)
}

type Tracer interface {
	Start(ctx context.Context, name string, attributes ...Attribute) (context.Context, Span)
}
