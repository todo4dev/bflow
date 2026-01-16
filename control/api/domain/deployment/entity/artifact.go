// domain/deployment/entity/artifact.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Artifact_KindEnum string

const (
	Artifact_KindDeployment Artifact_KindEnum = "deployment"
	Artifact_KindService    Artifact_KindEnum = "service"
)

type ArtifactEntity struct {
	ID        uuid.UUID         `json:"id"`
	TS        time.Time         `json:"ts"`
	CreatedAt time.Time         `json:"created_at"`
	DeletedAt *time.Time        `json:"deleted_at"`
	Kind      Artifact_KindEnum `json:"kind"`
	Name      string            `json:"name"`
	Metadata  json.RawMessage   `json:"metadata"`
}

var _ json.Marshaler = (*ArtifactEntity)(nil)
var _ json.Unmarshaler = (*ArtifactEntity)(nil)

func (a *ArtifactEntity) MarshalJSON() ([]byte, error) {
	type Alias ArtifactEntity
	return json.Marshal((*Alias)(a))
}

func (a *ArtifactEntity) UnmarshalJSON(data []byte) error {
	type Alias ArtifactEntity
	return json.Unmarshal(data, (*Alias)(a))
}
