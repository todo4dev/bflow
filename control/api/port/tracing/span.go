package tracing

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
