// port/sqldb/client.go
package sqldb

import "context"

// Row represents a single row of results
type Row interface {
	Scan(dest ...any) error
}

// Rows represents multiple rows of results
type Rows interface {
	Next() bool
	Scan(dest ...any) error
	Close() error
	Err() error
}

// Result represents the result of a write operation
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

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
