// port/jwt/provider.go
package jwt

import "time"

// Kind represents a JWT kind
type Kind string

const (
	Kind_ACCESS  Kind = "access"
	Kind_REFRESH Kind = "refresh"
)

// Claims represents a JWT claim
type Claims struct {
	Subject    string
	Email      string
	GivenName  *string
	FamilyName *string
	Picture    *string
	Language   *string
	Theme      *string
	Timezone   *string
}

// Token represents a JWT token
type Token struct {
	TokenType    string
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
}

// Decoded represents a decoded JWT token
type Decoded struct {
	Kind      Kind
	SessionID string
	Claims    Claims
	IssuedAt  time.Time
	ExpiresAt time.Time
}

// Provider generates and validates JWT tokens
type Provider interface {
	// Create creates a new JWT token
	Create(sessionID string, claims Claims, optionalIncludeRefresh ...bool) (*Token, error)

	// Decode decodes a JWT token
	Decode(token string) (*Decoded, error)
}
