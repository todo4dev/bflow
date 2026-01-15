// control/api/domain/identity/entity/certificate.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type CertificateEntity struct {
	ID        uuid.UUID  `json:"id"`
	TS        time.Time  `json:"ts"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Kind      string     `json:"kind"`
	Content   string     `json:"content"`
	AccountID uuid.UUID  `json:"account_id"`
}

var _ json.Marshaler = (*CertificateEntity)(nil)
var _ json.Unmarshaler = (*CertificateEntity)(nil)

func (e *CertificateEntity) MarshalJSON() ([]byte, error) {
	type Alias CertificateEntity
	return json.Marshal((*Alias)(e))
}

func (e *CertificateEntity) UnmarshalJSON(data []byte) error {
	type Alias CertificateEntity
	return json.Unmarshal(data, (*Alias)(e))
}
