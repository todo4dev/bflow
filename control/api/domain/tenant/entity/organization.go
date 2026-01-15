// control/api/domain/tenant/entity/membership.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type OrganizationEntity struct {
	ID               uuid.UUID        `json:"id"`
	TS               time.Time        `json:"ts"`
	CreatedAt        time.Time        `json:"created_at"`
	DeletedAt        *time.Time       `json:"deleted_at"`
	Name             string           `json:"name"`
	Slug             string           `json:"slug"`
	Config           *json.RawMessage `json:"config"`
	CreatorAccountID uuid.UUID        `json:"creator_account_id"`
}

var _ json.Marshaler = (*OrganizationEntity)(nil)
var _ json.Unmarshaler = (*OrganizationEntity)(nil)

func (e *OrganizationEntity) MarshalJSON() ([]byte, error) {
	type Alias OrganizationEntity
	return json.Marshal((*Alias)(e))
}

func (e *OrganizationEntity) UnmarshalJSON(data []byte) error {
	type Alias OrganizationEntity
	return json.Unmarshal(data, (*Alias)(e))
}
