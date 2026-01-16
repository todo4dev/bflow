// domain/identity/entity/profile.go
package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Profile_ThemeEnum string

const (
	Profile_ThemeEnumLight  Profile_ThemeEnum = "light"
	Profile_ThemeEnumDark   Profile_ThemeEnum = "dark"
	Profile_ThemeEnumSystem Profile_ThemeEnum = "system"
)

type ProfileEntity struct {
	ID         uuid.UUID         `json:"id"`
	TS         time.Time         `json:"ts"`
	GivenName  *string           `json:"given_name"`
	FamilyName *string           `json:"family_name"`
	Theme      Profile_ThemeEnum `json:"theme"`
	Language   string            `json:"language"`
	Timezone   string            `json:"timezone"`
	AccountID  uuid.UUID         `json:"account_id"`
}

var _ json.Marshaler = (*ProfileEntity)(nil)
var _ json.Unmarshaler = (*ProfileEntity)(nil)

func (e *ProfileEntity) MarshalJSON() ([]byte, error) {
	type Alias ProfileEntity
	return json.Marshal((*Alias)(e))
}

func (e *ProfileEntity) UnmarshalJSON(data []byte) error {
	type Alias ProfileEntity
	return json.Unmarshal(data, (*Alias)(e))
}
