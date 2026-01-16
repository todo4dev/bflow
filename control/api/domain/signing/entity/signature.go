// domain/signing/entity/signature.go
package entity

import (
	"encoding/json"
	"src/domain/signing/enum"
	"time"

	"github.com/google/uuid"
)

type Signature struct {
	ID            uuid.UUID           `json:"id"`
	TS            time.Time           `json:"ts"`
	CreatedAt     time.Time           `json:"created_at"`
	State         enum.SignatureState `json:"state"`
	SignedAt      *time.Time          `json:"signed_at"`
	FailureReason *string             `json:"failure_reason"`
	Value         *string             `json:"value"`
	Hash          *string             `json:"hash"`
	Metadata      json.RawMessage     `json:"metadata"`
	DocumentID    uuid.UUID           `json:"document_id"`
	AccountID     uuid.UUID           `json:"account_id"`
	CertificateID *uuid.UUID          `json:"certificate_id"`
}

var _ json.Marshaler = (*Signature)(nil)
var _ json.Unmarshaler = (*Signature)(nil)

func (e *Signature) MarshalJSON() ([]byte, error) {
	type Alias Signature
	return json.Marshal((*Alias)(e))
}

func (e *Signature) UnmarshalJSON(data []byte) error {
	type Alias Signature
	return json.Unmarshal(data, (*Alias)(e))
}
