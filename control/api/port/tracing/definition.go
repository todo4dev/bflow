package tracing

// Definition span configuration
type Definition struct {
	Attributes map[string]any
	Kind       Kind
}
