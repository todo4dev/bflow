// domain/deployment/entity/pipeline.go
package entity

import (
	"encoding/json"
	"src/domain/deployment/enum"
	"time"

	"github.com/google/uuid"
)

type Pipeline struct {
	ID                 uuid.UUID           `json:"id"`
	TS                 time.Time           `json:"ts"`
	CreatedAt          time.Time           `json:"created_at"`
	DeletedAt          *time.Time          `json:"deleted_at"`
	Kind               enum.PipelineKind   `json:"kind"`
	Status             enum.PipelineStatus `json:"status"`
	Payload            *json.RawMessage    `json:"payload"`
	StartedAt          *time.Time          `json:"started_at"`
	FinishedAt         *time.Time          `json:"finished_at"`
	ErrorMessage       *string             `json:"error_message"`
	TargetReleaseID    *uuid.UUID          `json:"target_release_id"`
	RuntimeID          uuid.UUID           `json:"runtime_id"`
	RequesterAccountID uuid.UUID           `json:"requester_account_id"`
	PreviousPipelineID *uuid.UUID          `json:"previous_pipeline_id"`
}

var _ json.Marshaler = (*Pipeline)(nil)
var _ json.Unmarshaler = (*Pipeline)(nil)

func (e *Pipeline) MarshalJSON() ([]byte, error) {
	type Alias Pipeline
	return json.Marshal((*Alias)(e))
}

func (e *Pipeline) UnmarshalJSON(data []byte) error {
	type Alias Pipeline
	return json.Unmarshal(data, (*Alias)(e))
}
