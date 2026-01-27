// infrastructure/interpolate/std/render.go
package std

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sync"

	"src/port/interpolate"
)

type Renderer struct {
	config    *Config
	templates map[string]*template.Template
	mu        sync.RWMutex
}

var _ interpolate.Renderer = (*Renderer)(nil)

func New(config *Config) (*Renderer, error) {
	renderer := &Renderer{
		config:    config,
		templates: make(map[string]*template.Template),
	}

	if err := renderer.loadTemplates(); err != nil {
		return nil, err
	}

	return renderer, nil
}

func (r *Renderer) loadTemplates() error {
	entries, err := os.ReadDir(r.config.Path)
	if err != nil {
		return fmt.Errorf("failed to read template directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		fullPath := filepath.Join(r.config.Path, name)

		content, err := os.ReadFile(fullPath)
		if err != nil {
			return fmt.Errorf("failed to read template file %s: %w", fullPath, err)
		}

		t, err := template.New(name).Parse(string(content))
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", name, err)
		}

		r.templates[name] = t
	}

	return nil
}

func (r *Renderer) getTemplate(name string) (*template.Template, error) {
	r.mu.RLock()
	tmpl, ok := r.templates[name]
	r.mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("template not found: %s", name)
	}
	return tmpl, nil
}

func (r *Renderer) Render(templateName string, data map[string]any) (string, error) {
	tmpl, err := r.getTemplate(templateName)
	if err != nil {
		return "", err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return "", err
	}

	return body.String(), nil
}
