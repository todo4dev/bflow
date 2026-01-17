// port/docdb/client.go
package docdb

import "context"

// Cursor represents multiple documents
type Cursor interface {
	Next(ctx context.Context) bool
	Decode(result any) error
	Close(ctx context.Context) error
	Err() error
}

// Client represents a document database (MongoDB)
type Client interface {
	// Insert inserts one or more documents
	Insert(ctx context.Context, collection string, documents ...any) error

	// Find finds documents with filters
	Find(ctx context.Context, collection string, filter map[string]any) (Cursor, error)

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
