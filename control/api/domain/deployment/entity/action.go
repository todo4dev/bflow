// domain/deployment/entity/action.go
package entity

import (
	"encoding/json"
	"src/domain/deployment/enum"
	"time"

	"github.com/google/uuid"
)

type Action struct {
	ID               uuid.UUID         `json:"id"`
	TS               time.Time         `json:"ts"`
	CreatedAt        time.Time         `json:"created_at"`
	DeletedAt        *time.Time        `json:"deleted_at"`
	Kind             enum.ActionKind   `json:"kind"`
	Status           enum.ActionStatus `json:"status"`
	StartedAt        *time.Time        `json:"started_at"`
	FinishedAt       *time.Time        `json:"finished_at"`
	ErrorMessage     *string           `json:"error_message"`
	PipelineID       uuid.UUID         `json:"pipeline_id"`
	ExecutionAgentID uuid.UUID         `json:"execution_agent_id"`
}

var _ json.Marshaler = (*Action)(nil)
var _ json.Unmarshaler = (*Action)(nil)

func (e *Action) MarshalJSON() ([]byte, error) {
	type Alias Action
	return json.Marshal((*Alias)(e))
}

func (e *Action) UnmarshalJSON(data []byte) error {
	type Alias Action
	return json.Unmarshal(data, (*Alias)(e))
}
