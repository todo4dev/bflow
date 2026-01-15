// control/api/application/port/storage/file.go
package storage

import (
	"context"
	"io"
	"time"
)

type File struct {
	Key         string
	ContentType string
	ETag        string
	UpdatedAt   time.Time
}

type Bucket interface {
	Put(ctx context.Context, key string, contentType string, body io.Reader) (File, error)
	Get(ctx context.Context, key string) (File, io.ReadCloser, error)
	Delete(ctx context.Context, key string) error
}
