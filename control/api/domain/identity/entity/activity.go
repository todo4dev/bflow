// domain/identity/entity/activity.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Activity_TypeEnum string

const (
	Activity_TypeUnknown Activity_TypeEnum = "unknown"
)

type ActivityEntity struct {
	ID        uuid.UUID         `json:"id"`
	TS        time.Time         `json:"ts"`
	CreatedAt time.Time         `json:"created_at"`
	Type      Activity_TypeEnum `json:"type"`
	Message   *string           `json:"message"`
	Metadata  *json.RawMessage  `json:"metadata"`
	AccountID *uuid.UUID        `json:"account_id"`
	ClusterID uuid.UUID         `json:"cluster_id"`
}

var _ json.Marshaler = (*ActivityEntity)(nil)
var _ json.Unmarshaler = (*ActivityEntity)(nil)

func (e *ActivityEntity) MarshalJSON() ([]byte, error) {
	type Alias ActivityEntity
	return json.Marshal((*Alias)(e))
}

func (e *ActivityEntity) UnmarshalJSON(data []byte) error {
	type Alias ActivityEntity
	return json.Unmarshal(data, (*Alias)(e))
}
