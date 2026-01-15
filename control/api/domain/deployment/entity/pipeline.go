// control/api/domain/deployment/entity/pipeline.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Pipeline_TypeEnum string

const (
	Pipeline_TypeBuild  Pipeline_TypeEnum = "build"
	Pipeline_TypeDeploy Pipeline_TypeEnum = "deploy"
)

type Pipeline_StatusEnum string

const (
	Pipeline_StatusPending   Pipeline_StatusEnum = "pending"
	Pipeline_StatusSucceeded Pipeline_StatusEnum = "succeeded"
	Pipeline_StatusFailed    Pipeline_StatusEnum = "failed"
	Pipeline_StatusCanceled  Pipeline_StatusEnum = "canceled"
)

type PipelineEntity struct {
	ID                 uuid.UUID           `json:"id"`
	TS                 time.Time           `json:"ts"`
	CreatedAt          time.Time           `json:"created_at"`
	DeletedAt          *time.Time          `json:"deleted_at"`
	Type               Pipeline_TypeEnum   `json:"type"`
	Status             Pipeline_StatusEnum `json:"status"`
	Payload            *json.RawMessage    `json:"payload"`
	StartedAt          *time.Time          `json:"started_at"`
	FinishedAt         *time.Time          `json:"finished_at"`
	ErrorMessage       *string             `json:"error_message"`
	TargetReleaseID    *uuid.UUID          `json:"target_release_id"`
	RuntimeID          uuid.UUID           `json:"runtime_id"`
	RequesterAccountID uuid.UUID           `json:"requester_account_id"`
	PreviousPipelineID *uuid.UUID          `json:"previous_pipeline_id"`
}

var _ json.Marshaler = (*PipelineEntity)(nil)
var _ json.Unmarshaler = (*PipelineEntity)(nil)

func (e *PipelineEntity) MarshalJSON() ([]byte, error) {
	type Alias PipelineEntity
	return json.Marshal((*Alias)(e))
}

func (e *PipelineEntity) UnmarshalJSON(data []byte) error {
	type Alias PipelineEntity
	return json.Unmarshal(data, (*Alias)(e))
}
