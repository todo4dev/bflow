// domain/entity/pipeline_action_stage.go
package entity

import (
	"encoding/json"
	"src/domain/enum"
	"time"

	"github.com/google/uuid"
)

type PipelineActionStage struct {
	ID               uuid.UUID                      `json:"id"`
	TS               time.Time                      `json:"ts"`
	CreatedAt        time.Time                      `json:"created_at"`
	DeletedAt        *time.Time                     `json:"deleted_at"`
	Name             string                         `json:"name"`
	Position         int                            `json:"position"`
	Status           enum.PipelineActionStageStatus `json:"status"`
	StartedAt        *time.Time                     `json:"started_at"`
	FinishedAt       *time.Time                     `json:"finished_at"`
	Summary          *string                        `json:"summary"`
	OutputMeta       json.RawMessage                `json:"output_meta"`
	PipelineActionID uuid.UUID                      `json:"pipeline_action_id"`
}

var _ json.Marshaler = (*PipelineActionStage)(nil)
var _ json.Unmarshaler = (*PipelineActionStage)(nil)

func (e *PipelineActionStage) MarshalJSON() ([]byte, error) {
	type Alias PipelineActionStage
	return json.Marshal((*Alias)(e))
}

func (e *PipelineActionStage) UnmarshalJSON(data []byte) error {
	type Alias PipelineActionStage
	return json.Unmarshal(data, (*Alias)(e))
}

