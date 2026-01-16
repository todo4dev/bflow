// domain/deployment/entity/runtime.go
package entity

import (
	"encoding/json"
	"src/domain/deployment/enum"
	"time"

	"github.com/google/uuid"
)

type Runtime struct {
	ID               uuid.UUID         `json:"id"`
	TS               time.Time         `json:"ts"`
	CreatedAt        time.Time         `json:"created_at"`
	DeletedAt        *time.Time        `json:"deleted_at"`
	State            enum.RuntimeState `json:"state"`
	ReadOnly         bool              `json:"readonly"`
	LastDeployedAt   *time.Time        `json:"last_deployed_at"`
	Config           json.RawMessage   `json:"config"`
	CurrentReleaseID *uuid.UUID        `json:"current_release_id"`
	ClusterID        uuid.UUID         `json:"cluster_id"`
}

var _ json.Marshaler = (*Runtime)(nil)
var _ json.Unmarshaler = (*Runtime)(nil)

func (e *Runtime) MarshalJSON() ([]byte, error) {
	type Alias Runtime
	return json.Marshal((*Alias)(e))
}

func (e *Runtime) UnmarshalJSON(data []byte) error {
	type Alias Runtime
	return json.Unmarshal(data, (*Alias)(e))
}
