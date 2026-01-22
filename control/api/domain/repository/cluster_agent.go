// domain/repository/cluster_agent.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type ClusterAgent interface {
	Create(agent *entity.ClusterAgent) error
	Save(agent *entity.ClusterAgent) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*entity.ClusterAgent, error)
	FindByClusterId(clusterId uuid.UUID) ([]*entity.ClusterAgent, error)
}
