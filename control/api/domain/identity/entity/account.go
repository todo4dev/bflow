// domain/identity/entity/account.go
package entity

import (
	"encoding/json"
	"src/domain"
	"src/domain/identity/enum"
	"src/domain/identity/event"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	domain.Aggregate[event.Account]

	ID              uuid.UUID        `json:"id"`
	TS              time.Time        `json:"ts"`
	CreatedAt       time.Time        `json:"created_at"`
	DeletedAt       *time.Time       `json:"deleted_at"`
	Email           string           `json:"email"`
	EmailVerifiedAt *time.Time       `json:"email_verified_at"`
	Role            enum.AccountRole `json:"role"`
}

var _ json.Marshaler = (*Account)(nil)
var _ json.Unmarshaler = (*Account)(nil)

func (e *Account) MarshalJSON() ([]byte, error) {
	type Alias Account
	return json.Marshal((*Alias)(e))
}

func (e *Account) UnmarshalJSON(data []byte) error {
	type Alias Account
	return json.Unmarshal(data, (*Alias)(e))
}
