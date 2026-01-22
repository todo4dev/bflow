// domain/repository/document.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type Document interface {
	Create(document *entity.Document) error
	Save(document *entity.Document) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*entity.Document, error)
	FindByOrganizationId(organizationId uuid.UUID) ([]*entity.Document, error)
}
