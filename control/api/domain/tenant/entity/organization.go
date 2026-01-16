// domain/tenant/entity/organization.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID               uuid.UUID        `json:"id"`
	TS               time.Time        `json:"ts"`
	CreatedAt        time.Time        `json:"created_at"`
	DeletedAt        *time.Time       `json:"deleted_at"`
	Name             string           `json:"name"`
	Slug             string           `json:"slug"`
	Config           *json.RawMessage `json:"config"`
	CreatorAccountID uuid.UUID        `json:"creator_account_id"`
}

var _ json.Marshaler = (*Organization)(nil)
var _ json.Unmarshaler = (*Organization)(nil)

func (e *Organization) MarshalJSON() ([]byte, error) {
	type Alias Organization
	return json.Marshal((*Alias)(e))
}

func (e *Organization) UnmarshalJSON(data []byte) error {
	type Alias Organization
	return json.Unmarshal(data, (*Alias)(e))
}
