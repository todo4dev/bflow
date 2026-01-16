// domain/identity/entity/account.go
package entity

import (
	"encoding/json"
	"src/domain/identity/event"
	"src/domain/shared"
	"time"

	"github.com/google/uuid"
)

type Account_RoleEnum string

const (
	Account_RoleAdmin  Account_RoleEnum = "admin"
	Account_RoleMember Account_RoleEnum = "member"
)

type AccountEntity struct {
	shared.Aggregate[event.AccountEventEnum]

	ID              uuid.UUID        `db:"id"`
	TS              time.Time        `db:"ts"`
	CreatedAt       time.Time        `db:"created_at"`
	DeletedAt       *time.Time       `db:"deleted_at"`
	Email           string           `db:"email"`
	EmailVerifiedAt *time.Time       `db:"email_verified_at"`
	Role            Account_RoleEnum `db:"role"`
}

var _ json.Marshaler = (*AccountEntity)(nil)
var _ json.Unmarshaler = (*AccountEntity)(nil)

func (e *AccountEntity) MarshalJSON() ([]byte, error) {
	type Alias AccountEntity
	return json.Marshal((*Alias)(e))
}

func (e *AccountEntity) UnmarshalJSON(data []byte) error {
	type Alias AccountEntity
	return json.Unmarshal(data, (*Alias)(e))
}
