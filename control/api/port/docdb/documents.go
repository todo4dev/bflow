// port/docdb/documents.go
package docdb

import "context"

// Documents represents multiple documents
type Documents interface {
	Next(ctx context.Context) bool
	Decode(result any) error
	Close(ctx context.Context) error
	Err() error
}
