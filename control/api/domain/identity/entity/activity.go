// domain/identity/entity/activity.go
package entity

import (
	"encoding/json"
	"src/domain/identity/enum"
	"time"

	"github.com/google/uuid"
)

type Activity struct {
	ID        uuid.UUID         `json:"id"`
	TS        time.Time         `json:"ts"`
	CreatedAt time.Time         `json:"created_at"`
	Kind      enum.ActivityKind `json:"kind"`
	Message   *string           `json:"message"`
	Metadata  *json.RawMessage  `json:"metadata"`
	AccountID *uuid.UUID        `json:"account_id"`
	ClusterID uuid.UUID         `json:"cluster_id"`
}

var _ json.Marshaler = (*Activity)(nil)
var _ json.Unmarshaler = (*Activity)(nil)

func (e *Activity) MarshalJSON() ([]byte, error) {
	type Alias Activity
	return json.Marshal((*Alias)(e))
}

func (e *Activity) UnmarshalJSON(data []byte) error {
	type Alias Activity
	return json.Unmarshal(data, (*Alias)(e))
}
