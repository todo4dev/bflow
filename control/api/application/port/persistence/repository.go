// control/api/application/port/persistence/repository.go
package persistence

import "context"

type Insert[TEntity any] interface {
	Insert(ctx context.Context, entity TEntity) error
}

type Update[TEntity any] interface {
	Update(ctx context.Context, entity TEntity) error
}

type Delete[TID comparable] interface {
	Delete(ctx context.Context, id TID) error
}
