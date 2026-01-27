package interpolate

// Renderer renders a template with the given data
type Renderer interface {
	Render(templateName string, data map[string]any) (string, error)
}
