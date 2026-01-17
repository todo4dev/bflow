// domain/identity/entity/account_notification.go
package entity

import (
	"encoding/json"
	"src/domain/identity/enum"
	"time"

	"github.com/google/uuid"
)

type AccountNotification struct {
	ID         uuid.UUID                      `json:"id"`
	TS         time.Time                      `json:"ts"`
	CreatedAt  time.Time                      `json:"created_at"`
	DeletedAt  *time.Time                     `json:"deleted_at"`
	Kind       enum.AccountNotificationKind   `json:"kind"`
	Level      enum.AccountNotificationLevel  `json:"level"`
	Status     enum.AccountNotificationStatus `json:"status"`
	Message    string                         `json:"message"`
	Metadata   *json.RawMessage               `json:"metadata"`
	ArtifactID *uuid.UUID                     `json:"artifact_id"`
	ClusterID  uuid.UUID                      `json:"cluster_id"`
	AccountID  uuid.UUID                      `json:"account_id"`
}

var _ json.Marshaler = (*AccountNotification)(nil)
var _ json.Unmarshaler = (*AccountNotification)(nil)

func (e *AccountNotification) MarshalJSON() ([]byte, error) {
	type Alias AccountNotification
	return json.Marshal((*Alias)(e))
}

func (e *AccountNotification) UnmarshalJSON(data []byte) error {
	type Alias AccountNotification
	return json.Unmarshal(data, (*Alias)(e))
}
