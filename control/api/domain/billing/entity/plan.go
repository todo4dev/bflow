// domain/billing/entity/plan.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Plan_IntervalEnum string

const (
	Plan_IntervalHourly  Plan_IntervalEnum = "hourly"
	Plan_IntervalDaily   Plan_IntervalEnum = "daily"
	Plan_IntervalWeekly  Plan_IntervalEnum = "weekly"
	Plan_IntervalMonthly Plan_IntervalEnum = "monthly"
	Plan_IntervalYearly  Plan_IntervalEnum = "yearly"
)

type PlanEntity struct {
	ID         uuid.UUID         `json:"id"`
	TS         time.Time         `json:"ts"`
	CreatedAt  time.Time         `json:"created_at"`
	DeletedAt  *time.Time        `json:"deleted_at"`
	Code       string            `json:"code"`
	Name       string            `json:"name"`
	Interval   Plan_IntervalEnum `json:"interval"`
	Currency   string            `json:"currency"`
	PriceCents int               `json:"price_cents"`
	Config     json.RawMessage   `json:"config"`
}

var _ json.Marshaler = (*PlanEntity)(nil)
var _ json.Unmarshaler = (*PlanEntity)(nil)

func (e *PlanEntity) MarshalJSON() ([]byte, error) {
	type Alias PlanEntity
	return json.Marshal((*Alias)(e))
}

func (e *PlanEntity) UnmarshalJSON(data []byte) error {
	type Alias PlanEntity
	return json.Unmarshal(data, (*Alias)(e))
}
