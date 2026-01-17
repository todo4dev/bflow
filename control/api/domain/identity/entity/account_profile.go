// domain/identity/entity/account_profile.go
package entity

import (
	"encoding/json"
	"src/domain/identity/enum"
	"time"

	"github.com/google/uuid"
)

type AccountProfile struct {
	ID         uuid.UUID         `json:"id"`
	TS         time.Time         `json:"ts"`
	GivenName  *string           `json:"given_name"`
	FamilyName *string           `json:"family_name"`
	Theme      enum.ProfileTheme `json:"theme"`
	Language   string            `json:"language"`
	Timezone   string            `json:"timezone"`
	AccountID  uuid.UUID         `json:"account_id"`
}

var _ json.Marshaler = (*AccountProfile)(nil)
var _ json.Unmarshaler = (*AccountProfile)(nil)

func (e *AccountProfile) MarshalJSON() ([]byte, error) {
	type Alias AccountProfile
	return json.Marshal((*Alias)(e))
}

func (e *AccountProfile) UnmarshalJSON(data []byte) error {
	type Alias AccountProfile
	return json.Unmarshal(data, (*Alias)(e))
}
