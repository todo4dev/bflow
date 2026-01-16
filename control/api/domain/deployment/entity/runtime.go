// domain/deployment/entity/runtime.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Runtime_StateEnum string

const (
	Runtime_StateEnumActive   Runtime_StateEnum = "active"
	Runtime_StateEnumInactive Runtime_StateEnum = "inactive"
)

type RuntimeEntity struct {
	ID               uuid.UUID         `json:"id"`
	TS               time.Time         `json:"ts"`
	CreatedAt        time.Time         `json:"created_at"`
	DeletedAt        *time.Time        `json:"deleted_at"`
	State            Runtime_StateEnum `json:"state"`
	ReadOnly         bool              `json:"readonly"`
	LastDeployedAt   *time.Time        `json:"last_deployed_at"`
	Config           json.RawMessage   `json:"config"`
	CurrentReleaseID *uuid.UUID        `json:"current_release_id"`
	ClusterID        uuid.UUID         `json:"cluster_id"`
}

var _ json.Marshaler = (*RuntimeEntity)(nil)
var _ json.Unmarshaler = (*RuntimeEntity)(nil)

func (e *RuntimeEntity) MarshalJSON() ([]byte, error) {
	type Alias RuntimeEntity
	return json.Marshal((*Alias)(e))
}

func (e *RuntimeEntity) UnmarshalJSON(data []byte) error {
	type Alias RuntimeEntity
	return json.Unmarshal(data, (*Alias)(e))
}
