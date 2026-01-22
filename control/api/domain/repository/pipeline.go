// domain/repository/pipeline.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type Pipeline interface {
	Create(pipeline *entity.Pipeline) error
	Save(pipeline *entity.Pipeline) error
	FindById(id uuid.UUID) (*entity.Pipeline, error)
	FindByOrganizationId(organizationId uuid.UUID) ([]*entity.Pipeline, error)
	FindByClusterRuntimeId(runtimeId uuid.UUID) ([]*entity.Pipeline, error)
}
