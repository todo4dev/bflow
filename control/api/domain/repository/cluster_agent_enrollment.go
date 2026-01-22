// domain/repository/cluster_agent_enrollment.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type ClusterAgentEnrollment interface {
	Create(enrollment *entity.ClusterAgentEnrollment) error
	Save(enrollment *entity.ClusterAgentEnrollment) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*entity.ClusterAgentEnrollment, error)
	FindByToken(token string) (*entity.ClusterAgentEnrollment, error)
}
