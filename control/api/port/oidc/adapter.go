// port/oidc/adapter.go
package oidc

import (
	"context"
	"io"
)

// Claims OIDC ID token claims
type Claims struct {
	Subject       string
	Email         string
	EmailVerified bool
	Name          string
	Picture       string
	Custom        map[string]any
}

// Tokens tokens returned by OIDC
type Tokens struct {
	AccessToken  string
	IDToken      string
	RefreshToken string
	ExpiresIn    int64
}

// Adapter authenticates via OIDC (OpenID Connect)
type Adapter interface {
	// GetAuthURL returns authentication URL
	GetAuthURL(state string, scopes []string) string

	// Exchange exchanges code for tokens
	Exchange(ctx context.Context, code string) (*Tokens, error)

	// GetToken exchanges refresh token for tokens
	GetToken(ctx context.Context, refreshToken string) (*Tokens, error)

	// GetInfo returns user info
	GetInfo(ctx context.Context, accessToken string) (*Claims, error)

	// GetPicture returns user picture
	GetPicture(ctx context.Context, accessToken string) (io.ReadCloser, error)
}

// ProviderName represents an OIDC provider
type ProviderName string

const (
	ProviderName_MICROSOFT ProviderName = "microsoft"
	ProviderName_GOOGLE    ProviderName = "google"
)

// Provider creates an Adapter for the given provider
type Provider interface {
	// GetAdapter returns an Adapter for the given provider
	GetAdapter(provider ProviderName) (Adapter, error)
}
