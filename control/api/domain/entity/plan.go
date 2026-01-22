// domain/entity/plan.go
package entity

import (
	"encoding/json"
	"src/domain/enum"
	"time"

	"github.com/google/uuid"
)

type Plan struct {
	ID         uuid.UUID         `json:"id"`
	TS         time.Time         `json:"ts"`
	CreatedAt  time.Time         `json:"created_at"`
	DeletedAt  *time.Time        `json:"deleted_at"`
	Code       string            `json:"code"`
	Name       string            `json:"name"`
	Interval   enum.PlanInterval `json:"interval"`
	Currency   string            `json:"currency"`
	PriceCents int               `json:"price_cents"`
	Config     json.RawMessage   `json:"config"`
}

var _ json.Marshaler = (*Plan)(nil)
var _ json.Unmarshaler = (*Plan)(nil)

func (e *Plan) MarshalJSON() ([]byte, error) {
	type Alias Plan
	return json.Marshal((*Alias)(e))
}

func (e *Plan) UnmarshalJSON(data []byte) error {
	type Alias Plan
	return json.Unmarshal(data, (*Alias)(e))
}

