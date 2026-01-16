// domain/signing/entity/document.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Document_KindEnum string

const (
	Document_KindAgreement Document_KindEnum = "agreement"
	Document_KindInvoice   Document_KindEnum = "invoice"
	Document_KindReceipt   Document_KindEnum = "receipt"
	Document_KindReport    Document_KindEnum = "report"
	Document_KindOther     Document_KindEnum = "other"
)

type Document_StatusEnum string

const (
	Document_StatusActive   Document_StatusEnum = "active"
	Document_StatusReplaced Document_StatusEnum = "replaced"
	Document_StatusRevoked  Document_StatusEnum = "revoked"
)

type DocumentEntity struct {
	ID                 uuid.UUID           `json:"id"`
	TS                 time.Time           `json:"ts"`
	CreatedAt          time.Time           `json:"created_at"`
	DeletedAt          *time.Time          `json:"deleted_at"`
	Kind               Document_KindEnum   `json:"kind"`
	Status             Document_StatusEnum `json:"status"`
	Title              string              `json:"title"`
	StorageKey         string              `json:"storage_key"`
	MimeType           string              `json:"mime_type"`
	FileSize           int64               `json:"file_size"`
	ContentSHA256      string              `json:"content_sha256"`
	Metadata           json.RawMessage     `json:"metadata"`
	ReplacedDocumentID *uuid.UUID          `json:"replaced_document_id"`
	OrganizationID     *uuid.UUID          `json:"organization_id"`
}

var _ json.Marshaler = (*DocumentEntity)(nil)
var _ json.Unmarshaler = (*DocumentEntity)(nil)

func (e *DocumentEntity) MarshalJSON() ([]byte, error) {
	type Alias DocumentEntity
	return json.Marshal((*Alias)(e))
}

func (e *DocumentEntity) UnmarshalJSON(data []byte) error {
	type Alias DocumentEntity
	return json.Unmarshal(data, (*Alias)(e))
}
