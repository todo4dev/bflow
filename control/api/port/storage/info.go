// port/storage/info.go
package storage

import "time"

// Info represents information about a stored object
type Info struct {
	Path         string
	Size         int64
	LastModified time.Time
	ContentType  string
	ETag         string
	Metadata     map[string]string
}
