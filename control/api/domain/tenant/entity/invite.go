// domain/tenant/entity/invite.go
package entity

import (
	"encoding/json"
	"src/domain/tenant/enum"
	"time"

	"github.com/google/uuid"
)

type Invite struct {
	ID               uuid.UUID           `json:"id"`
	TS               time.Time           `json:"ts"`
	CreatedAt        time.Time           `json:"created_at"`
	Code             string              `json:"code"`
	Email            string              `json:"email"`
	Role             enum.MembershipRole `json:"role"`
	Status           enum.InviteStatus   `json:"status"`
	ExpiresAt        time.Time           `json:"expires_at"`
	AcceptedAt       *time.Time          `json:"accepted_at"`
	OrganizationID   uuid.UUID           `json:"organization_id"`
	CreatorAccountID uuid.UUID           `json:"creator_account_id"`
}

var _ json.Marshaler = (*Invite)(nil)
var _ json.Unmarshaler = (*Invite)(nil)

func (e *Invite) MarshalJSON() ([]byte, error) {
	type Alias Invite
	return json.Marshal((*Alias)(e))
}

func (e *Invite) UnmarshalJSON(data []byte) error {
	type Alias Invite
	return json.Unmarshal(data, (*Alias)(e))
}
