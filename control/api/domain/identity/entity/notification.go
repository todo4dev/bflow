// control/api/domain/identity/entity/notification.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Notification_KindEnum string

const (
	Notification_KindSystem   Notification_KindEnum = "system"
	Notification_KindBilling  Notification_KindEnum = "billing"
	Notification_KindSecurity Notification_KindEnum = "security"
	Notification_KindRelease  Notification_KindEnum = "release"
	Notification_KindRuntime  Notification_KindEnum = "runtime"
	Notification_KindPipeline Notification_KindEnum = "pipeline"
	Notification_KindAgent    Notification_KindEnum = "agent"
	Notification_KindDocument Notification_KindEnum = "document"
)

type Notification_LevelEnum string

const (
	Notification_LevelInfo    Notification_LevelEnum = "info"
	Notification_LevelWarning Notification_LevelEnum = "warning"
	Notification_LevelError   Notification_LevelEnum = "error"
)

type Notification_StatusEnum string

const (
	Notification_StatusActive    Notification_StatusEnum = "active"
	Notification_StatusDismissed Notification_StatusEnum = "dismissed"
)

type NotificationEntity struct {
	ID         uuid.UUID               `json:"id"`
	TS         time.Time               `json:"ts"`
	CreatedAt  time.Time               `json:"created_at"`
	DeletedAt  *time.Time              `json:"deleted_at"`
	Kind       Notification_KindEnum   `json:"kind"`
	Level      Notification_LevelEnum  `json:"level"`
	Status     Notification_StatusEnum `json:"status"`
	Message    string                  `json:"message"`
	Metadata   *json.RawMessage        `json:"metadata"`
	ArtifactID *uuid.UUID              `json:"artifact_id"`
	ClusterID  uuid.UUID               `json:"cluster_id"`
	AccountID  uuid.UUID               `json:"account_id"`
}

var _ json.Marshaler = (*NotificationEntity)(nil)
var _ json.Unmarshaler = (*NotificationEntity)(nil)

func (e *NotificationEntity) MarshalJSON() ([]byte, error) {
	type Alias NotificationEntity
	return json.Marshal((*Alias)(e))
}

func (e *NotificationEntity) UnmarshalJSON(data []byte) error {
	type Alias NotificationEntity
	return json.Unmarshal(data, (*Alias)(e))
}
