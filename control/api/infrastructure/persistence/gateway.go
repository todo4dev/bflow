// infrastructure/persistence/gateway.go
// application/port/persistence/gateway.go
package persistence

import "context"

type FindOne[TType any, TQuery any] interface {
	FindOne(ctx context.Context, query TQuery) (TType, error)
}

type FindMany[TType any, TQuery any] interface {
	FindMany(ctx context.Context, query TQuery) ([]TType, error)
}

type Count[TQuery any] interface {
	Count(ctx context.Context, query TQuery) (int, error)
}
