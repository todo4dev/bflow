// domain/identity/entity/credential.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type CredentialEntity struct {
	ID           uuid.UUID `json:"id"`
	TS           time.Time `json:"ts"`
	CreatedAt    time.Time `json:"created_at"`
	PasswordHash string    `json:"password_hash"`
	AccountID    uuid.UUID `json:"account_id"`
}

var _ json.Marshaler = (*CredentialEntity)(nil)
var _ json.Unmarshaler = (*CredentialEntity)(nil)

func (e *CredentialEntity) MarshalJSON() ([]byte, error) {
	type Alias CredentialEntity
	return json.Marshal((*Alias)(e))
}

func (e *CredentialEntity) UnmarshalJSON(data []byte) error {
	type Alias CredentialEntity
	return json.Unmarshal(data, (*Alias)(e))
}
