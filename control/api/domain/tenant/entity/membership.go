// domain/tenant/entity/membership.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Membership_RoleEnum string

const (
	Membership_RoleAdmin   Membership_RoleEnum = "admin"
	Membership_RoleManager Membership_RoleEnum = "manager"
	Membership_RoleViewer  Membership_RoleEnum = "viewer"
	Membership_RoleMember  Membership_RoleEnum = "member"
)

type MembershipEntity struct {
	ID             uuid.UUID `json:"id"`
	TS             time.Time `json:"ts"`
	Role           string    `json:"role"`
	AccountID      uuid.UUID `json:"account_id"`
	OrganizationID uuid.UUID `json:"organization_id"`
}

var _ json.Marshaler = (*MembershipEntity)(nil)
var _ json.Unmarshaler = (*MembershipEntity)(nil)

func (e *MembershipEntity) MarshalJSON() ([]byte, error) {
	type Alias MembershipEntity
	return json.Marshal((*Alias)(e))
}

func (e *MembershipEntity) UnmarshalJSON(data []byte) error {
	type Alias MembershipEntity
	return json.Unmarshal(data, (*Alias)(e))
}
