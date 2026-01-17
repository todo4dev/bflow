// domain/signing/entity/document.go
package entity

import (
	"encoding/json"
	"src/domain"
	"src/domain/signing/enum"
	"src/domain/signing/event"
	"time"

	"github.com/google/uuid"
)

type Document struct {
	domain.Aggregate[event.Document]

	ID                 uuid.UUID           `json:"id"`
	TS                 time.Time           `json:"ts"`
	CreatedAt          time.Time           `json:"created_at"`
	DeletedAt          *time.Time          `json:"deleted_at"`
	Kind               enum.DocumentKind   `json:"kind"`
	Status             enum.DocumentStatus `json:"status"`
	Title              string              `json:"title"`
	StorageKey         string              `json:"storage_key"`
	Mimetype           string              `json:"mimetype"`
	FileSize           int64               `json:"file_size"`
	ContentSHA256      string              `json:"content_sha256"`
	Metadata           json.RawMessage     `json:"metadata"`
	ReplacedDocumentID *uuid.UUID          `json:"replaced_document_id"`
	OrganizationID     *uuid.UUID          `json:"organization_id"`
}

var _ json.Marshaler = (*Document)(nil)
var _ json.Unmarshaler = (*Document)(nil)

func (e *Document) MarshalJSON() ([]byte, error) {
	type Alias Document
	return json.Marshal((*Alias)(e))
}

func (e *Document) UnmarshalJSON(data []byte) error {
	type Alias Document
	return json.Unmarshal(data, (*Alias)(e))
}
