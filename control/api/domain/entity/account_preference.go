// domain/entity/account_preference.go
package entity

import (
	"encoding/json"
	"src/domain/enum"
	"time"

	"github.com/google/uuid"
)

type AccountPreference struct {
	ID                      uuid.UUID            `json:"id"`
	TS                      time.Time            `json:"ts"`
	Theme                   enum.PreferenceTheme `json:"theme"`
	NotifyOnPipelineSuccess bool                 `json:"notify_on_pipeline_success"`
	NotifyOnInfraAlerts     bool                 `json:"notify_on_infra_alerts"`
	AccountID               uuid.UUID            `json:"account_id"`
}

var _ json.Marshaler = (*AccountPreference)(nil)
var _ json.Unmarshaler = (*AccountPreference)(nil)

func (e *AccountPreference) MarshalJSON() ([]byte, error) {
	type Alias AccountPreference
	return json.Marshal((*Alias)(e))
}

func (e *AccountPreference) UnmarshalJSON(data []byte) error {
	type Alias AccountPreference
	return json.Unmarshal(data, (*Alias)(e))
}

