// port/auth/oidc_provider.go
package auth

import "context"

// OIDCClaims OIDC ID token claims
type OIDCClaims struct {
	Subject       string
	Email         string
	EmailVerified bool
	Name          string
	Picture       string
	Custom        map[string]any
}

// OIDCTokens tokens returned by OIDC
type OIDCTokens struct {
	AccessToken  string
	IDToken      string
	RefreshToken string
	ExpiresIn    int64
}

// OIDCProvider authenticates via OIDC (OpenID Connect)
type OIDCProvider interface {
	// GetAuthURL returns authentication URL
	GetAuthURL(state string, scopes []string) string

	// Exchange exchanges code for tokens
	Exchange(ctx context.Context, code string) (*OIDCTokens, error)

	// VerifyIDToken verifies ID token
	VerifyIDToken(ctx context.Context, idToken string) (*OIDCClaims, error)
}
