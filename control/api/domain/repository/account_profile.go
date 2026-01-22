// domain/repository/account_profile.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type AccountProfile interface {
	Create(profile *entity.AccountProfile) error
	Save(profile *entity.AccountProfile) error
	FindByAccountId(accountId uuid.UUID) (*entity.AccountProfile, error)
}

