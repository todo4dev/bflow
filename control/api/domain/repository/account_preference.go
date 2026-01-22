// domain/repository/account_preference.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type AccountPreference interface {
	Create(preference *entity.AccountPreference) error
	Save(preference *entity.AccountPreference) error
	FindByAccountId(accountId uuid.UUID) (*entity.AccountPreference, error)
}

