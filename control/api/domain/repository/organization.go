// domain/repository/organization.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type Organization interface {
	Create(organization *entity.Organization) error
	Save(organization *entity.Organization) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*entity.Organization, error)
	FindBySlug(slug string) (*entity.Organization, error)
}
