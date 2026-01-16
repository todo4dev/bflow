// domain/identity/entity/profile.go
package entity

import (
	"encoding/json"
	"src/domain/identity/enum"
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID         uuid.UUID         `json:"id"`
	TS         time.Time         `json:"ts"`
	GivenName  *string           `json:"given_name"`
	FamilyName *string           `json:"family_name"`
	Theme      enum.ProfileTheme `json:"theme"`
	Language   string            `json:"language"`
	Timezone   string            `json:"timezone"`
	AccountID  uuid.UUID         `json:"account_id"`
}

var _ json.Marshaler = (*Profile)(nil)
var _ json.Unmarshaler = (*Profile)(nil)

func (e *Profile) MarshalJSON() ([]byte, error) {
	type Alias Profile
	return json.Marshal((*Alias)(e))
}

func (e *Profile) UnmarshalJSON(data []byte) error {
	type Alias Profile
	return json.Unmarshal(data, (*Alias)(e))
}
