// domain/identity/repository/account_credential.go
package repository

import (
	"src/domain/identity/entity"

	"github.com/google/uuid"
)

type AccountCredential interface {
	Create(credential *entity.Credential) error
	Save(credential *entity.Credential) error
	FindByAccountId(accountId uuid.UUID) (*entity.Credential, error)
}
