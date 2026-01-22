// domain/repository/pipeline_action.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type PipelineAction interface {
	Create(action *entity.PipelineAction) error
	Save(action *entity.PipelineAction) error
	FindById(id uuid.UUID) (*entity.PipelineAction, error)
	FindByPipelineId(pipelineId uuid.UUID) ([]*entity.PipelineAction, error)
}
