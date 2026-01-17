// port/docdb/store.go
package docdb

import "context"

// DocumentStore represents a document database (MongoDB)
type DocumentStore interface {
	// Insert inserts one or more documents
	Insert(ctx context.Context, collection string, documents ...any) error

	// Find finds documents with filters
	Find(ctx context.Context, collection string, filter map[string]any) (Documents, error)

	// FindOne finds a single document
	FindOne(ctx context.Context, collection string, filter map[string]any, result any) error

	// Update updates documents
	Update(ctx context.Context, collection string, filter map[string]any, update map[string]any) error

	// Delete deletes documents
	Delete(ctx context.Context, collection string, filter map[string]any) error

	// Close closes the connection
	Close() error

	// Ping verifies connectivity
	Ping(ctx context.Context) error
}
