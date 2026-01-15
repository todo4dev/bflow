// control/api/domain/signing/entity/document.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Signature_StateEnum string

const (
	Signature_StatePending  Signature_StateEnum = "pending"
	Signature_StateSigned   Signature_StateEnum = "signed"
	Signature_StateFailed   Signature_StateEnum = "failed"
	Signature_StateCanceled Signature_StateEnum = "canceled"
)

type SignatureEntity struct {
	ID            uuid.UUID           `json:"id"`
	TS            time.Time           `json:"ts"`
	CreatedAt     time.Time           `json:"created_at"`
	State         Signature_StateEnum `json:"state"`
	SignedAt      *time.Time          `json:"signed_at"`
	FailureReason *string             `json:"failure_reason"`
	Value         *string             `json:"value"`
	Hash          *string             `json:"hash"`
	Metadata      json.RawMessage     `json:"metadata"`
	DocumentID    uuid.UUID           `json:"document_id"`
	AccountID     uuid.UUID           `json:"account_id"`
	CertificateID *uuid.UUID          `json:"certificate_id"`
}

var _ json.Marshaler = (*SignatureEntity)(nil)
var _ json.Unmarshaler = (*SignatureEntity)(nil)

func (e *SignatureEntity) MarshalJSON() ([]byte, error) {
	type Alias SignatureEntity
	return json.Marshal((*Alias)(e))
}

func (e *SignatureEntity) UnmarshalJSON(data []byte) error {
	type Alias SignatureEntity
	return json.Unmarshal(data, (*Alias)(e))
}
