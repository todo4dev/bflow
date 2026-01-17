// domain/identity/entity/account_activity.go
package entity

import (
	"encoding/json"
	"src/domain/identity/enum"
	"time"

	"github.com/google/uuid"
)

type AccountActivity struct {
	ID        uuid.UUID                `json:"id"`
	TS        time.Time                `json:"ts"`
	CreatedAt time.Time                `json:"created_at"`
	Kind      enum.AccountActivityKind `json:"kind"`
	Message   *string                  `json:"message"`
	Metadata  *json.RawMessage         `json:"metadata"`
	AccountID *uuid.UUID               `json:"account_id"`
	ClusterID uuid.UUID                `json:"cluster_id"`
}

var _ json.Marshaler = (*AccountActivity)(nil)
var _ json.Unmarshaler = (*AccountActivity)(nil)

func (e *AccountActivity) MarshalJSON() ([]byte, error) {
	type Alias AccountActivity
	return json.Marshal((*Alias)(e))
}

func (e *AccountActivity) UnmarshalJSON(data []byte) error {
	type Alias AccountActivity
	return json.Unmarshal(data, (*Alias)(e))
}
