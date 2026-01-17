// port/tracing/tracer.go
package tracing

import "context"

// Tracer represents a distributed tracing tracer
type Tracer interface {
	// StartSpan starts a new span
	StartSpan(ctx context.Context, name string, defs ...func(*Definition)) (context.Context, Span)

	// Extract extracts context from headers
	Extract(ctx context.Context, carrier map[string]string) context.Context

	// Inject injects context into headers
	Inject(ctx context.Context, carrier map[string]string)
}
