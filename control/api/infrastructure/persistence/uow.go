// application/port/persistence/uow.go
package persistence

import "context"

type UnitOfWork interface {
	Do(ctx context.Context, fn func(context.Context) error) error
}
