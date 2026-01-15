// control/api/domain/deployment/entity/action.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Action_TypeEnum string

const (
	Action_TypeDeploy   Action_TypeEnum = "deploy"
	Action_TypeConfig   Action_TypeEnum = "config"
	Action_TypeRollback Action_TypeEnum = "rollback"
	Action_TypeDestroy  Action_TypeEnum = "destroy"
)

type Action_StatusEnum string

const (
	Action_StatusPending  Action_StatusEnum = "pending"
	Action_StatusRunning  Action_StatusEnum = "running"
	Action_StatusSuccess  Action_StatusEnum = "success"
	Action_StatusFailure  Action_StatusEnum = "failure"
	Action_StatusCanceled Action_StatusEnum = "canceled"
)

type ActionEntity struct {
	ID               uuid.UUID         `json:"id"`
	TS               time.Time         `json:"ts"`
	CreatedAt        time.Time         `json:"created_at"`
	DeletedAt        *time.Time        `json:"deleted_at"`
	Type             Action_TypeEnum   `json:"type"`
	Status           Action_StatusEnum `json:"status"`
	StartedAt        *time.Time        `json:"started_at"`
	FinishedAt       *time.Time        `json:"finished_at"`
	ErrorMessage     *string           `json:"error_message"`
	PipelineID       uuid.UUID         `json:"pipeline_id"`
	ExecutionAgentID uuid.UUID         `json:"execution_agent_id"`
}

var _ json.Marshaler = (*ActionEntity)(nil)
var _ json.Unmarshaler = (*ActionEntity)(nil)

func (e *ActionEntity) MarshalJSON() ([]byte, error) {
	type Alias ActionEntity
	return json.Marshal((*Alias)(e))
}

func (e *ActionEntity) UnmarshalJSON(data []byte) error {
	type Alias ActionEntity
	return json.Unmarshal(data, (*Alias)(e))
}
