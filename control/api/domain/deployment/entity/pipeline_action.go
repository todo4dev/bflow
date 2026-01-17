// domain/deployment/entity/pipeline_action.go
package entity

import (
	"encoding/json"
	"src/domain/deployment/enum"
	"time"

	"github.com/google/uuid"
)

type PipelineAction struct {
	ID                      uuid.UUID                 `json:"id"`
	TS                      time.Time                 `json:"ts"`
	CreatedAt               time.Time                 `json:"created_at"`
	DeletedAt               *time.Time                `json:"deleted_at"`
	Kind                    enum.PipelineActionKind   `json:"kind"`
	Status                  enum.PipelineActionStatus `json:"status"`
	StartedAt               *time.Time                `json:"started_at"`
	FinishedAt              *time.Time                `json:"finished_at"`
	ErrorMessage            *string                   `json:"error_message"`
	PipelineID              uuid.UUID                 `json:"pipeline_id"`
	ExecutionClusterAgentID uuid.UUID                 `json:"execution_cluster_agent_id"`
}

var _ json.Marshaler = (*PipelineAction)(nil)
var _ json.Unmarshaler = (*PipelineAction)(nil)

func (e *PipelineAction) MarshalJSON() ([]byte, error) {
	type Alias PipelineAction
	return json.Marshal((*Alias)(e))
}

func (e *PipelineAction) UnmarshalJSON(data []byte) error {
	type Alias PipelineAction
	return json.Unmarshal(data, (*Alias)(e))
}
