// domain/identity/entity/certificate.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Certificate struct {
	ID             uuid.UUID  `json:"id"`
	TS             time.Time  `json:"ts"`
	CreatedAt      time.Time  `json:"created_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
	ExpiresAt      *time.Time `json:"expires_at"`
	DisplayName    string     `json:"display_name"`
	DocumentNumber string     `json:"document_number"`
	OwnerName      string     `json:"owner_name"`
	Thumbprint     string     `json:"thumbprint"`
	IsActive       bool       `json:"is_active"`
	AccountID      uuid.UUID  `json:"account_id"`
}

var _ json.Marshaler = (*Certificate)(nil)
var _ json.Unmarshaler = (*Certificate)(nil)

func (e *Certificate) MarshalJSON() ([]byte, error) {
	type Alias Certificate
	return json.Marshal((*Alias)(e))
}

func (e *Certificate) UnmarshalJSON(data []byte) error {
	type Alias Certificate
	return json.Unmarshal(data, (*Alias)(e))
}
