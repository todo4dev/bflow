// domain/deployment/entity/agent.go
package entity

import (
	"encoding/json"
	"src/domain/deployment/enum"
	"time"

	"github.com/google/uuid"
)

type Agent struct {
	ID         uuid.UUID        `json:"id"`
	TS         time.Time        `json:"ts"`
	CreatedAt  time.Time        `json:"created_at"`
	DeletedAt  *time.Time       `json:"deleted_at"`
	Status     enum.AgentStatus `json:"status"`
	Version    string           `json:"version"`
	LastSeenAt *time.Time       `json:"last_seen_at"`
	PublicKey  *string          `json:"public_key"`
	Metadata   *json.RawMessage `json:"metadata"`
	ClusterID  uuid.UUID        `json:"cluster_id"`
}

var _ json.Marshaler = (*Agent)(nil)
var _ json.Unmarshaler = (*Agent)(nil)

func (a *Agent) MarshalJSON() ([]byte, error) {
	type Alias Agent
	return json.Marshal((*Alias)(a))
}

func (a *Agent) UnmarshalJSON(data []byte) error {
	type Alias Agent
	return json.Unmarshal(data, (*Alias)(a))
}
