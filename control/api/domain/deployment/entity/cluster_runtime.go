// domain/deployment/entity/cluster_runtime.go
package entity

import (
	"encoding/json"
	"src/domain/deployment/enum"
	"time"

	"github.com/google/uuid"
)

type ClusterRuntime struct {
	ID                       uuid.UUID                `json:"id"`
	TS                       time.Time                `json:"ts"`
	CreatedAt                time.Time                `json:"created_at"`
	DeletedAt                *time.Time               `json:"deleted_at"`
	State                    enum.ClusterRuntimeState `json:"state"`
	ReadOnly                 bool                     `json:"readonly"`
	LastDeployedAt           *time.Time               `json:"last_deployed_at"`
	Config                   json.RawMessage          `json:"config"`
	CurrentArtifactReleaseID *uuid.UUID               `json:"current_artifact_release_id"`
	ClusterID                uuid.UUID                `json:"cluster_id"`
}

var _ json.Marshaler = (*ClusterRuntime)(nil)
var _ json.Unmarshaler = (*ClusterRuntime)(nil)

func (e *ClusterRuntime) MarshalJSON() ([]byte, error) {
	type Alias ClusterRuntime
	return json.Marshal((*Alias)(e))
}

func (e *ClusterRuntime) UnmarshalJSON(data []byte) error {
	type Alias ClusterRuntime
	return json.Unmarshal(data, (*Alias)(e))
}
