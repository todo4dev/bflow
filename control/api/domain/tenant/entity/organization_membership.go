// domain/tenant/entity/organization_membership.go
package entity

import (
	"encoding/json"
	"src/domain/tenant/enum"
	"time"

	"github.com/google/uuid"
)

type OrganizationMembership struct {
	ID             uuid.UUID                       `json:"id"`
	TS             time.Time                       `json:"ts"`
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
