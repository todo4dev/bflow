// domain/identity/entity/notification.go
package entity

import (
	"encoding/json"
	"src/domain/identity/enum"
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID         uuid.UUID               `json:"id"`
	TS         time.Time               `json:"ts"`
	CreatedAt  time.Time               `json:"created_at"`
	DeletedAt  *time.Time              `json:"deleted_at"`
	Kind       enum.NotificationKind   `json:"kind"`
	Level      enum.NotificationLevel  `json:"level"`
	Status     enum.NotificationStatus `json:"status"`
	Message    string                  `json:"message"`
	Metadata   *json.RawMessage        `json:"metadata"`
	ArtifactID *uuid.UUID              `json:"artifact_id"`
	ClusterID  uuid.UUID               `json:"cluster_id"`
	AccountID  uuid.UUID               `json:"account_id"`
}

var _ json.Marshaler = (*Notification)(nil)
var _ json.Unmarshaler = (*Notification)(nil)

func (e *Notification) MarshalJSON() ([]byte, error) {
	type Alias Notification
	return json.Marshal((*Alias)(e))
}

func (e *Notification) UnmarshalJSON(data []byte) error {
	type Alias Notification
	return json.Unmarshal(data, (*Alias)(e))
}
