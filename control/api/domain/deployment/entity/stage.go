// control/api/domain/deployment/entity/stage.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type StageEntity struct {
	ID         uuid.UUID       `json:"id"`
	TS         time.Time       `json:"ts"`
	CreatedAt  time.Time       `json:"created_at"`
	DeletedAt  *time.Time      `json:"deleted_at"`
	Name       string          `json:"name"`
	Position   int             `json:"position"`
	Status     string          `json:"status"`
	StartedAt  *time.Time      `json:"started_at"`
	FinishedAt *time.Time      `json:"finished_at"`
	Summary    *string         `json:"summary"`
	OutputMeta json.RawMessage `json:"output_meta"`
	ActionID   uuid.UUID       `json:"action_id"`
}

var _ json.Marshaler = (*StageEntity)(nil)
var _ json.Unmarshaler = (*StageEntity)(nil)

func (e *StageEntity) MarshalJSON() ([]byte, error) {
	type Alias StageEntity
	return json.Marshal((*Alias)(e))
}

func (e *StageEntity) UnmarshalJSON(data []byte) error {
	type Alias StageEntity
	return json.Unmarshal(data, (*Alias)(e))
}
