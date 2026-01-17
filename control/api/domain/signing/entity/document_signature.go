// domain/signing/entity/document_signature.go
package entity

import (
	"encoding/json"
	"src/domain/signing/enum"
	"time"

	"github.com/google/uuid"
)

type DocumentSignature struct {
	ID                   uuid.UUID                   `json:"id"`
	TS                   time.Time                   `json:"ts"`
	CreatedAt            time.Time                   `json:"created_at"`
	State                enum.DocumentSignatureState `json:"state"`
	SignedAt             *time.Time                  `json:"signed_at"`
	FailureReason        *string                     `json:"failure_reason"`
	Value                *string                     `json:"value"`
	Hash                 *string                     `json:"hash"`
	Metadata             json.RawMessage             `json:"metadata"`
	DocumentID           uuid.UUID                   `json:"document_id"`
	AccountCertificateID uuid.UUID                   `json:"account_certificate_id"`
}

var _ json.Marshaler = (*DocumentSignature)(nil)
var _ json.Unmarshaler = (*DocumentSignature)(nil)

func (e *DocumentSignature) MarshalJSON() ([]byte, error) {
	type Alias DocumentSignature
	return json.Marshal((*Alias)(e))
}

func (e *DocumentSignature) UnmarshalJSON(data []byte) error {
	type Alias DocumentSignature
	return json.Unmarshal(data, (*Alias)(e))
}
