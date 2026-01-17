// domain/identity/repository/account.go
package repository

import (
	"src/domain/identity/entity"

	"github.com/google/uuid"
)

type Account interface {
	Create(account *entity.Account) error
	Save(account *entity.Account) error
	Disable(id uuid.UUID) error
	FindById(id uuid.UUID) (*entity.Account, error)
	FindByEmail(email string) (*entity.Account, error)
	ExistsByEmail(email string) (bool, error)
}
