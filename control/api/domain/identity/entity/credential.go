// domain/identity/entity/credential.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Credential struct {
	ID           uuid.UUID `json:"id"`
	TS           time.Time `json:"ts"`
	CreatedAt    time.Time `json:"created_at"`
	PasswordHash string    `json:"password_hash"`
	AccountID    uuid.UUID `json:"account_id"`
}

var _ json.Marshaler = (*Credential)(nil)
var _ json.Unmarshaler = (*Credential)(nil)

func (e *Credential) MarshalJSON() ([]byte, error) {
	type Alias Credential
	return json.Marshal((*Alias)(e))
}

func (e *Credential) UnmarshalJSON(data []byte) error {
	type Alias Credential
	return json.Unmarshal(data, (*Alias)(e))
}
