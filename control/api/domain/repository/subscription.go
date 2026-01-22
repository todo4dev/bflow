// domain/repository/subscription.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type Subscription interface {
	Create(subscription *entity.Subscription) error
	Save(subscription *entity.Subscription) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*entity.Subscription, error)
	FindByOrganizationId(organizationId uuid.UUID) ([]*entity.Subscription, error)
}
