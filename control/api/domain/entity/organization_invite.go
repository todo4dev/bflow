// domain/entity/organization_invite.go
package entity

import (
	"encoding/json"
	"src/domain/enum"
	"time"

	"github.com/google/uuid"
)

type OrganizationInvite struct {
	ID               uuid.UUID                       `json:"id"`
	TS               time.Time                       `json:"ts"`
	CreatedAt        time.Time                       `json:"created_at"`
	Code             string                          `json:"code"`
	Email            string                          `json:"email"`
	Role             enum.OrganizationMembershipRole `json:"role"`
	Status           enum.OrganizationInviteStatus   `json:"status"`
	ExpiresAt        time.Time                       `json:"expires_at"`
	AcceptedAt       *time.Time                      `json:"accepted_at"`
	OrganizationID   uuid.UUID                       `json:"organization_id"`
	CreatorAccountID uuid.UUID                       `json:"creator_account_id"`
}

var _ json.Marshaler = (*OrganizationInvite)(nil)
var _ json.Unmarshaler = (*OrganizationInvite)(nil)

func (e *OrganizationInvite) MarshalJSON() ([]byte, error) {
	type Alias OrganizationInvite
	return json.Marshal((*Alias)(e))
}

func (e *OrganizationInvite) UnmarshalJSON(data []byte) error {
	type Alias OrganizationInvite
	return json.Unmarshal(data, (*Alias)(e))
}

