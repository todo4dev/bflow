// control/api/domain/deployment/entity/agent.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Agent_StatusEnum string

const (
	Agent_StatusPending  Agent_StatusEnum = "pending"
	Agent_StatusOnline   Agent_StatusEnum = "online"
	Agent_StatusOffline  Agent_StatusEnum = "offline"
	Agent_StatusDisabled Agent_StatusEnum = "disabled"
	Agent_StatusDraining Agent_StatusEnum = "draining"
)

type Agent_AuthTypeEnum string

const (
	Agent_AuthTypeSecret Agent_AuthTypeEnum = "secret"
	Agent_AuthTypeKey    Agent_AuthTypeEnum = "key"
)

type AgentEntity struct {
	ID         uuid.UUID        `json:"id"`
	TS         time.Time        `json:"ts"`
	CreatedAt  time.Time        `json:"created_at"`
	DeletedAt  *time.Time       `json:"deleted_at"`
	Status     Agent_StatusEnum `json:"status"`
	Version    string           `json:"version"`
	LastSeenAt *time.Time       `json:"last_seen_at"`
	PublicKey  *string          `json:"public_key"`
	Metadata   *json.RawMessage `json:"metadata"`
	ClusterID  uuid.UUID        `json:"cluster_id"`
}

var _ json.Marshaler = (*AgentEntity)(nil)
var _ json.Unmarshaler = (*AgentEntity)(nil)

func (a *AgentEntity) MarshalJSON() ([]byte, error) {
	type Alias AgentEntity
	return json.Marshal((*Alias)(a))
}

func (a *AgentEntity) UnmarshalJSON(data []byte) error {
	type Alias AgentEntity
	return json.Unmarshal(data, (*Alias)(a))
}
