// domain/repository/cluster.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type Cluster interface {
	Create(cluster *entity.Cluster) error
	Save(cluster *entity.Cluster) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*entity.Cluster, error)
	FindByOrganizationId(organizationId uuid.UUID) ([]*entity.Cluster, error)
}
