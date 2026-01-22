// domain/repository/cluster_runtime.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type ClusterRuntime interface {
	Create(runtime *entity.ClusterRuntime) error
	Save(runtime *entity.ClusterRuntime) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*entity.ClusterRuntime, error)
	FindByClusterId(clusterId uuid.UUID) ([]*entity.ClusterRuntime, error)
}
