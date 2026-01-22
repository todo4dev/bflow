// domain/entity/cluster_agent_enrollment.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type ClusterAgentEnrollment struct {
	ID               uuid.UUID `json:"id"`
	TS               time.Time `json:"ts"`
	CreatedAt        time.Time `json:"created_at"`
	Token            string    `json:"token"`
	ExpiresAt        time.Time `json:"expires_at"`
	Used             bool      `json:"used"`
	ClusterID        uuid.UUID `json:"cluster_id"`
	CreatorAccountID uuid.UUID `json:"creator_account_id"`
}

var _ json.Marshaler = (*ClusterAgentEnrollment)(nil)
var _ json.Unmarshaler = (*ClusterAgentEnrollment)(nil)

func (e *ClusterAgentEnrollment) MarshalJSON() ([]byte, error) {
	type Alias ClusterAgentEnrollment
	return json.Marshal((*Alias)(e))
}

func (e *ClusterAgentEnrollment) UnmarshalJSON(data []byte) error {
	type Alias ClusterAgentEnrollment
	return json.Unmarshal(data, (*Alias)(e))
}

