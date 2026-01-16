// domain/tenant/entity/membership.go
package entity

import (
	"encoding/json"
	"src/domain/tenant/enum"
	"time"

	"github.com/google/uuid"
)

type Membership struct {
	ID             uuid.UUID           `json:"id"`
	TS             time.Time           `json:"ts"`
	Role           enum.MembershipRole `json:"role"`
	AccountID      uuid.UUID           `json:"account_id"`
	OrganizationID uuid.UUID           `json:"organization_id"`
}

var _ json.Marshaler = (*Membership)(nil)
var _ json.Unmarshaler = (*Membership)(nil)

func (e *Membership) MarshalJSON() ([]byte, error) {
	type Alias Membership
	return json.Marshal((*Alias)(e))
}

func (e *Membership) UnmarshalJSON(data []byte) error {
	type Alias Membership
	return json.Unmarshal(data, (*Alias)(e))
}
