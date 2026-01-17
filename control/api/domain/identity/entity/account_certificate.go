// domain/identity/entity/account_certificate.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type AccountCertificate struct {
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

var _ json.Marshaler = (*AccountCertificate)(nil)
var _ json.Unmarshaler = (*AccountCertificate)(nil)

func (e *AccountCertificate) MarshalJSON() ([]byte, error) {
	type Alias AccountCertificate
	return json.Marshal((*Alias)(e))
}

func (e *AccountCertificate) UnmarshalJSON(data []byte) error {
	type Alias AccountCertificate
	return json.Unmarshal(data, (*Alias)(e))
}
