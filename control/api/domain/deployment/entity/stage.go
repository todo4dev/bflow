// domain/deployment/entity/stage.go
package entity

import (
	"encoding/json"
	"src/domain/deployment/enum"
	"time"

	"github.com/google/uuid"
)

type Stage struct {
	ID         uuid.UUID        `json:"id"`
	TS         time.Time        `json:"ts"`
	CreatedAt  time.Time        `json:"created_at"`
	DeletedAt  *time.Time       `json:"deleted_at"`
	Name       string           `json:"name"`
	Position   int              `json:"position"`
	Status     enum.StageStatus `json:"status"`
	StartedAt  *time.Time       `json:"started_at"`
	FinishedAt *time.Time       `json:"finished_at"`
	Summary    *string          `json:"summary"`
	OutputMeta json.RawMessage  `json:"output_meta"`
	ActionID   uuid.UUID        `json:"action_id"`
}

var _ json.Marshaler = (*Stage)(nil)
var _ json.Unmarshaler = (*Stage)(nil)

func (e *Stage) MarshalJSON() ([]byte, error) {
	type Alias Stage
	return json.Marshal((*Alias)(e))
}

func (e *Stage) UnmarshalJSON(data []byte) error {
	type Alias Stage
	return json.Unmarshal(data, (*Alias)(e))
}
