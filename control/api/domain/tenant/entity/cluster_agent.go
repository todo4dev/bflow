// domain/tenant/entity/cluster_agent.go
package entity

import (
	"encoding/json"
	"src/domain/tenant/enum"
	"time"

	"github.com/google/uuid"
)

type ClusterAgent struct {
	ID         uuid.UUID               `json:"id"`
	TS         time.Time               `json:"ts"`
	CreatedAt  time.Time               `json:"created_at"`
	DeletedAt  *time.Time              `json:"deleted_at"`
	Status     enum.ClusterAgentStatus `json:"status"`
	Version    string                  `json:"version"`
	LastSeenAt *time.Time              `json:"last_seen_at"`
	PublicKey  *string                 `json:"public_key"`
	Metadata   *json.RawMessage        `json:"metadata"`
	ClusterID  uuid.UUID               `json:"cluster_id"`
}

var _ json.Marshaler = (*ClusterAgent)(nil)
var _ json.Unmarshaler = (*ClusterAgent)(nil)

func (a *ClusterAgent) MarshalJSON() ([]byte, error) {
	type Alias ClusterAgent
	return json.Marshal((*Alias)(a))
}

func (a *ClusterAgent) UnmarshalJSON(data []byte) error {
	type Alias ClusterAgent
	return json.Unmarshal(data, (*Alias)(a))
}
