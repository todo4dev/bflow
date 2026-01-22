// port/cache/client.go
package cache

import (
	"context"
	"time"
)

// Client represents a cache storage (Redis)
type Client interface {
	// Set stores a value with TTL
	Set(ctx context.Context, key string, value any, ttl time.Duration) error

	// Get retrieves a value
	Get(ctx context.Context, match string) (key string, value string, err error)

	// GetBytes retrieves bytes
	GetBytes(ctx context.Context, match string) (key string, value []byte, err error)

	// Delete removes a key
	Delete(ctx context.Context, matches ...string) error

	// Exists checks if a key exists
	Exists(ctx context.Context, matches ...string) (int64, error)

	// Expire defines TTL for an existing key
	Expire(ctx context.Context, match string, ttl time.Duration) error

	// Increment increments an integer value
	Increment(ctx context.Context, match string) (int64, error)

	// Decrement decrements an integer value
	Decrement(ctx context.Context, match string) (int64, error)

	// Close closes the connection
	Close() error

	// Ping verifies connectivity
	Ping(ctx context.Context) error
}
