// domain/repository/organization_invite.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type OrganizationInvite interface {
	Create(invite *entity.OrganizationInvite) error
	Save(invite *entity.OrganizationInvite) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*entity.OrganizationInvite, error)
	FindByCode(code string) (*entity.OrganizationInvite, error)
	FindByOrganizationId(organizationId uuid.UUID) ([]*entity.OrganizationInvite, error)
}
