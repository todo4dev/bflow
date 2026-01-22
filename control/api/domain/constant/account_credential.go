// domain/constant/account_credential.go
package constant

import "src/domain/enum"

const (
	ACCOUNT_CREDENTIAL_PASSWORD_REGEX = `^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[^a-zA-Z0-9]).{8,50}$`
)

const (
	ACCOUNT_PREFERENCE_THEME_DEFAULT = enum.PreferenceTheme_LIGHT
	ACCOUNT_PROFILE_LANGUAGE_DEFAULT = "en"
	ACCOUNT_PROFILE_TIMEZONE_DEFAULT = "UTC"
)

