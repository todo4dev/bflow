// domain/entity/organization_membership.go
package entity

import (
	"encoding/json"
	"src/domain/enum"
	"time"

	"github.com/google/uuid"
)

type OrganizationMembership struct {
	ID             uuid.UUID                       `json:"id"`
	TS             time.Time                       `json:"ts"`
	CreatedAt      time.Time                       `json:"created_at"`
	Role           enum.OrganizationMembershipRole `json:"role"`
	AccountID      uuid.UUID                       `json:"account_id"`
	OrganizationID uuid.UUID                       `json:"organization_id"`
}

var _ json.Marshaler = (*OrganizationMembership)(nil)
var _ json.Unmarshaler = (*OrganizationMembership)(nil)

func (e *OrganizationMembership) MarshalJSON() ([]byte, error) {
	type Alias OrganizationMembership
	return json.Marshal((*Alias)(e))
}

func (e *OrganizationMembership) UnmarshalJSON(data []byte) error {
	type Alias OrganizationMembership
	return json.Unmarshal(data, (*Alias)(e))
}

