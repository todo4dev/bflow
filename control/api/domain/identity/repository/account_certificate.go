// domain/identity/repository/account_certificate.go
package repository

import (
	"src/domain/identity/entity"

	"github.com/google/uuid"
)

type AccountCertificate interface {
	Create(cert *entity.AccountCertificate) error
	Save(cert *entity.AccountCertificate) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*entity.AccountCertificate, error)
	FindByAccountId(accountId uuid.UUID) ([]*entity.AccountCertificate, error)
	FindActiveByAccountId(accountId uuid.UUID) (*entity.AccountCertificate, error)
}
