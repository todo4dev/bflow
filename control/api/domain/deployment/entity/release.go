// domain/deployment/entity/release.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Release_ChannelEnum string

const (
	Release_ChannelEnumStable Release_ChannelEnum = "stable"
	Release_ChannelEnumBeta   Release_ChannelEnum = "beta"
	Release_ChannelEnumAlpha  Release_ChannelEnum = "alpha"
)

type ReleaseEntity struct {
	ID          uuid.UUID           `json:"id"`
	TS          time.Time           `json:"ts"`
	CreatedAt   time.Time           `json:"created_at"`
	DeletedAt   *time.Time          `json:"deleted_at"`
	Version     string              `json:"version"`
	Channel     Release_ChannelEnum `json:"channel"`
	Recommended bool                `json:"recommended"`
	Notes       *string             `json:"notes"`
	PublishedAt *time.Time          `json:"published_at"`
	Metadata    json.RawMessage     `json:"metadata"`
	ArtifactID  uuid.UUID           `json:"artifact_id"`
}

var _ json.Marshaler = (*ReleaseEntity)(nil)
var _ json.Unmarshaler = (*ReleaseEntity)(nil)

func (e *ReleaseEntity) MarshalJSON() ([]byte, error) {
	type Alias ReleaseEntity
	return json.Marshal((*Alias)(e))
}

func (e *ReleaseEntity) UnmarshalJSON(data []byte) error {
	type Alias ReleaseEntity
	return json.Unmarshal(data, (*Alias)(e))
}
