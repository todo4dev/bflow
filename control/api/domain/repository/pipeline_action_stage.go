// domain/repository/pipeline_action_stage.go
package repository

import (
	"src/domain/entity"

	"github.com/google/uuid"
)

type PipelineActionStage interface {
	Create(stage *entity.PipelineActionStage) error
	Save(stage *entity.PipelineActionStage) error
	FindById(id uuid.UUID) (*entity.PipelineActionStage, error)
	FindByPipelineActionId(actionId uuid.UUID) ([]*entity.PipelineActionStage, error)
}
