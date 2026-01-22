// domain/entity/artifact.go
package entity

import (
	"encoding/json"
	"src/domain/enum"
	"time"

	"github.com/google/uuid"
)

type Artifact struct {
	ID        uuid.UUID         `json:"id"`
	TS        time.Time         `json:"ts"`
	CreatedAt time.Time         `json:"created_at"`
	DeletedAt *time.Time        `json:"deleted_at"`
	Kind      enum.ArtifactKind `json:"kind"`
	Name      string            `json:"name"`
	Metadata  json.RawMessage   `json:"metadata"`
}

var _ json.Marshaler = (*Artifact)(nil)
var _ json.Unmarshaler = (*Artifact)(nil)

func (a *Artifact) MarshalJSON() ([]byte, error) {
	type Alias Artifact
	return json.Marshal((*Alias)(a))
}

func (a *Artifact) UnmarshalJSON(data []byte) error {
	type Alias Artifact
	return json.Unmarshal(data, (*Alias)(a))
}

