// port/cache/store.go
package cache

import (
	"context"
	"time"
)

// Store represents a cache storage (Redis)
type Store interface {
	// Set stores a value with TTL
	Set(ctx context.Context, key string, value any, ttl time.Duration) error

	// Get retrieves a value
	Get(ctx context.Context, key string) (string, error)

	// GetBytes retrieves bytes
	GetBytes(ctx context.Context, key string) ([]byte, error)

	// Delete removes a key
	Delete(ctx context.Context, keys ...string) error

	// Exists checks if a key exists
	Exists(ctx context.Context, keys ...string) (int64, error)

	// Expire defines TTL for an existing key
	Expire(ctx context.Context, key string, ttl time.Duration) error

	// Increment increments an integer value
	Increment(ctx context.Context, key string) (int64, error)

	// Decrement decrements an integer value
	Decrement(ctx context.Context, key string) (int64, error)

	// Close closes the connection
	Close() error

	// Ping verifies connectivity
	Ping(ctx context.Context) error
}
