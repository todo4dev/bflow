// domain/repository/organization_membership.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type OrganizationMembership interface {
	Create(membership *entity.OrganizationMembership) error
	Save(membership *entity.OrganizationMembership) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*entity.OrganizationMembership, error)
	FindByAccountId(accountId uuid.UUID) ([]*entity.OrganizationMembership, error)
	FindByOrganizationId(organizationId uuid.UUID) ([]*entity.OrganizationMembership, error)
}
