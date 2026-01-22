// domain/entity/account_credential.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type AccountCredential struct {
	ID           uuid.UUID `json:"id"`
	TS           time.Time `json:"ts"`
	CreatedAt    time.Time `json:"created_at"`
	PasswordHash string    `json:"password_hash"`
	AccountID    uuid.UUID `json:"account_id"`
}

var _ json.Marshaler = (*AccountCredential)(nil)
var _ json.Unmarshaler = (*AccountCredential)(nil)

func (e *AccountCredential) MarshalJSON() ([]byte, error) {
	type Alias AccountCredential
	return json.Marshal((*Alias)(e))
}

func (e *AccountCredential) UnmarshalJSON(data []byte) error {
	type Alias AccountCredential
	return json.Unmarshal(data, (*Alias)(e))
}

