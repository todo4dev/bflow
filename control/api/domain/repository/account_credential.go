// domain/repository/account_credential.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type AccountCredential interface {
	Create(credential *entity.AccountCredential) error
	Save(credential *entity.AccountCredential) error
	FindByAccountId(accountId uuid.UUID) (*entity.AccountCredential, error)
}

