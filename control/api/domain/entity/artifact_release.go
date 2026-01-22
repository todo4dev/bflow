// domain/entity/artifact_release.go
package entity

import (
	"encoding/json"
	"src/domain/enum"
	"time"

	"github.com/google/uuid"
)

type ArtifactRelease struct {
	ID          uuid.UUID                   `json:"id"`
	TS          time.Time                   `json:"ts"`
	CreatedAt   time.Time                   `json:"created_at"`
	DeletedAt   *time.Time                  `json:"deleted_at"`
	Version     string                      `json:"version"`
	Channel     enum.ArtifactReleaseChannel `json:"channel"`
	Recommended bool                        `json:"recommended"`
	Notes       *string                     `json:"notes"`
	PublishedAt *time.Time                  `json:"published_at"`
	Metadata    json.RawMessage             `json:"metadata"`
	ArtifactID  uuid.UUID                   `json:"artifact_id"`
}

var _ json.Marshaler = (*ArtifactRelease)(nil)
var _ json.Unmarshaler = (*ArtifactRelease)(nil)

func (e *ArtifactRelease) MarshalJSON() ([]byte, error) {
	type Alias ArtifactRelease
	return json.Marshal((*Alias)(e))
}

func (e *ArtifactRelease) UnmarshalJSON(data []byte) error {
	type Alias ArtifactRelease
	return json.Unmarshal(data, (*Alias)(e))
}

