// port/auth/jwt_provider.go
package auth

import (
	"context"
	"time"
)

// JWTClaims represents JWT claims
type JWTClaims struct {
	Subject   string
	Email     string
	Role      string
	IssuedAt  time.Time
	ExpiresAt time.Time
	Custom    map[string]any
}

// JWTProvider generates and validates JWT tokens
type JWTProvider interface {
	// Generate generates a new JWT token
	Generate(ctx context.Context, claims JWTClaims) (string, error)

	// Validate validates and decodes a JWT token
	Validate(ctx context.Context, token string) (*JWTClaims, error)

	// Refresh refreshes a token
	Refresh(ctx context.Context, refreshToken string) (string, error)
}
