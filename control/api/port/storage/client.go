// port/storage/client.go
package storage

import (
	"context"
	"io"
	"time"
)

// Info represents information about a stored object
type Info struct {
	Path         string
	Size         int64
	LastModified time.Time
	ContentType  string
	ETag         string
	Metadata     map[string]string
}

// Client represents a generic object storage
type Client interface {
	// Upload uploads data to the specified path
	Upload(ctx context.Context, path string, reader io.Reader, metadata map[string]string) error

	// Download downloads data from the specified path
	Download(ctx context.Context, path string) (io.ReadCloser, error)

	// Delete removes data at the specified path
	Delete(ctx context.Context, path string) error

	// Exists checks if data exists at the specified path
	Exists(ctx context.Context, path string) (bool, error)

	// GetInfo returns information about the stored object
	GetInfo(ctx context.Context, path string) (*Info, error)

	// GenerateTemporaryURL generates a temporary access URL
	GenerateTemporaryURL(ctx context.Context, path string, expiration time.Duration) (string, error)

	// List lists objects with the given prefix
	List(ctx context.Context, prefix string) ([]Info, error)
}
