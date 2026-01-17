// port/sqldb/client.go
package sqldb

import "context"

// Client represents a generic database client
type Client interface {
	// Query executes a SELECT query and returns multiple rows
	Query(ctx context.Context, query string, args ...any) (Rows, error)

	// QueryRow executes a SELECT query and returns a single row
	QueryRow(ctx context.Context, query string, args ...any) Row

	// Exec executes INSERT, UPDATE, DELETE
	Exec(ctx context.Context, query string, args ...any) (Result, error)

	// Transaction executes a function within a transaction
	Transaction(ctx context.Context, fn func(tx Client) error) error

	// Close closes the connection
	Close() error

	// Ping verifies connectivity
	Ping(ctx context.Context) error
}
