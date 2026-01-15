// control/api/domain/tenant/entity/cluster.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Cluster_StateEnum string

const (
	Cluster_StateActive    Cluster_StateEnum = "active"
	Cluster_StateMigrating Cluster_StateEnum = "migrating"
	Cluster_StateLegacy    Cluster_StateEnum = "legacy"
	Cluster_StateDisabled  Cluster_StateEnum = "disabled"
)

type ClusterEntity struct {
	ID                uuid.UUID         `json:"id"`
	TS                time.Time         `json:"ts"`
	CreatedAt         time.Time         `json:"created_at"`
	DeletedAt         *time.Time        `json:"deleted_at"`
	State             Cluster_StateEnum `json:"state"`
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

var _ json.Marshaler = (*ClusterEntity)(nil)
var _ json.Unmarshaler = (*ClusterEntity)(nil)

func (e *ClusterEntity) MarshalJSON() ([]byte, error) {
	type Alias ClusterEntity
	return json.Marshal((*Alias)(e))
}

func (e *ClusterEntity) UnmarshalJSON(data []byte) error {
	type Alias ClusterEntity
	return json.Unmarshal(data, (*Alias)(e))
}
