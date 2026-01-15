// control/api/application/port/security/oidc.go
package security

import "io"

// Successful Token Response
// see https://openid.net/specs/openid-connect-core-1_0.html#TokenResponse
type OIDCToken struct {
	AccessToken  string
	RefreshToken string
}

// Openid user info
// see https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims
type OIDCInfo struct {
	Sub        string
	Email      string
	GivenName  *string
	FamilyName *string
	PictureURL *string
	Theme      *string
	Language   *string
	Timezone   *string
}

type OIDCProviderEnum string

const (
	OIDCProvider_Google    OIDCProviderEnum = "google"
	OIDCProvider_Microsoft OIDCProviderEnum = "microsoft"
)

type OIDCConnector interface {
	CreateRedirectUrl(state string) string
	ExchangeCode(code string) (OIDCToken, error)
	Refresh(token string) (OIDCToken, error)
	GetInfo(token string) (OIDCInfo, error)
	GetPicture(token string) (io.ReadCloser, error)
}

type OIDCProvider interface {
	GetConnector(kind OIDCProviderEnum) (OIDCConnector, error)
}
