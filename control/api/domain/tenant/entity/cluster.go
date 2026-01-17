// domain/tenant/entity/cluster.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"

	"src/domain"
	"src/domain/tenant/enum"
	"src/domain/tenant/event"
)

type Cluster struct {
	domain.Aggregate[event.Cluster]

	ID                uuid.UUID         `json:"id"`
	TS                time.Time         `json:"ts"`
	CreatedAt         time.Time         `json:"created_at"`
	DeletedAt         *time.Time        `json:"deleted_at"`
	State             enum.ClusterState `json:"state"`
	PromotedAt        *time.Time        `json:"promoted_at"`
	LegacyAt          *time.Time        `json:"legacy_at"`
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	Provider          *string           `json:"provider"`
	ExternalID        *string           `json:"external_id"`
	KubernetesVersion *string           `json:"kubernetes_version"`
	Metadata          json.RawMessage   `json:"metadata"`
	OrganizationID    uuid.UUID         `json:"organization_id"`
}

var _ json.Marshaler = (*Cluster)(nil)
var _ json.Unmarshaler = (*Cluster)(nil)

func (e *Cluster) MarshalJSON() ([]byte, error) {
	type Alias Cluster
	return json.Marshal((*Alias)(e))
}

func (e *Cluster) UnmarshalJSON(data []byte) error {
	type Alias Cluster
	return json.Unmarshal(data, (*Alias)(e))
}
