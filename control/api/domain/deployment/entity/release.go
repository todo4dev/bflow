// domain/deployment/entity/release.go
package entity

import (
	"encoding/json"
	"src/domain/deployment/enum"
	"time"

	"github.com/google/uuid"
)

type Release struct {
	ID          uuid.UUID           `json:"id"`
	TS          time.Time           `json:"ts"`
	CreatedAt   time.Time           `json:"created_at"`
	DeletedAt   *time.Time          `json:"deleted_at"`
	Version     string              `json:"version"`
	Channel     enum.ReleaseChannel `json:"channel"`
	Recommended bool                `json:"recommended"`
	Notes       *string             `json:"notes"`
	PublishedAt *time.Time          `json:"published_at"`
	Metadata    json.RawMessage     `json:"metadata"`
	ArtifactID  uuid.UUID           `json:"artifact_id"`
}

var _ json.Marshaler = (*Release)(nil)
var _ json.Unmarshaler = (*Release)(nil)

func (e *Release) MarshalJSON() ([]byte, error) {
	type Alias Release
	return json.Marshal((*Alias)(e))
}

func (e *Release) UnmarshalJSON(data []byte) error {
	type Alias Release
	return json.Unmarshal(data, (*Alias)(e))
}
