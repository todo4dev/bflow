package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Invite_StatusEnum string

const (
	Invite_StatusPending  Invite_StatusEnum = "pending"
	Invite_StatusAccepted Invite_StatusEnum = "accepted"
	Invite_StatusExpired  Invite_StatusEnum = "expired"
	Invite_StatusRevoked  Invite_StatusEnum = "revoked"
)

type InviteEntity struct {
	ID               uuid.UUID           `json:"id"`
	TS               time.Time           `json:"ts"`
	CreatedAt        time.Time           `json:"created_at"`
	Code             string              `json:"code"`
	Email            string              `json:"email"`
	Role             Membership_RoleEnum `json:"role"`
	Status           Invite_StatusEnum   `json:"status"`
	ExpiresAt        time.Time           `json:"expires_at"`
	AcceptedAt       *time.Time          `json:"accepted_at"`
	OrganizationID   uuid.UUID           `json:"organization_id"`
	CreatorAccountID uuid.UUID           `json:"creator_account_id"`
}

var _ json.Marshaler = (*InviteEntity)(nil)
var _ json.Unmarshaler = (*InviteEntity)(nil)

func (e *InviteEntity) MarshalJSON() ([]byte, error) {
	type Alias InviteEntity
	return json.Marshal((*Alias)(e))
}

func (e *InviteEntity) UnmarshalJSON(data []byte) error {
	type Alias InviteEntity
	return json.Unmarshal(data, (*Alias)(e))
}
