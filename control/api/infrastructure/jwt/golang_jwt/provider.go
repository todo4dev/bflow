// infrastructure/jwt/golang_jwt/provider.go
package golang_jwt

import (
	"errors"
	"src/port/jwt"
	"time"

	lib "github.com/golang-jwt/jwt/v5"
)

type customClaims struct {
	lib.RegisteredClaims
	Email      string  `json:"email"`
	GivenName  *string `json:"given_name,omitempty"`
	FamilyName *string `json:"family_name,omitempty"`
	Picture    *string `json:"picture,omitempty"`
	Language   *string `json:"language,omitempty"`
	Theme      *string `json:"theme,omitempty"`
	Timezone   *string `json:"timezone,omitempty"`
}

type Provider struct {
	algorithm  lib.SigningMethod
	issuer     string
	audience   string
	accessTTL  time.Duration
	refreshTTL time.Duration
	signKey    any
	verifyKey  any
}

var _ jwt.Provider = (*Provider)(nil)

func NewProvider(config *Config) (*Provider, error) {
	algorithm, ok := (map[any]lib.SigningMethod{
		"HS256": lib.SigningMethodHS256,
		"RS256": lib.SigningMethodRS256,
	})[config.Algorithm]
	if !ok {
		return nil, errors.New("unsupported algorithm")
	}

	signKey, verifyKey, err := parseKeys(config)
	if err != nil {
		return nil, err
	}

	return &Provider{
		algorithm:  algorithm,
		issuer:     config.Issuer,
		audience:   config.Audience,
		accessTTL:  config.AccessTTL,
		refreshTTL: config.RefreshTTL,
		signKey:    signKey,
		verifyKey:  verifyKey,
	}, nil
}

func (p *Provider) createToken(
	sessionID string,
	claims jwt.Claims,
	kind jwt.Kind,
	now time.Time,
	ttl time.Duration,
) (string, error) {
	custom := customClaims{
		RegisteredClaims: lib.RegisteredClaims{
			ID:        sessionID,
			Subject:   claims.Subject,
			Issuer:    p.issuer,
			Audience:  lib.ClaimStrings{p.audience},
			IssuedAt:  lib.NewNumericDate(now),
			ExpiresAt: lib.NewNumericDate(now.Add(ttl)),
		},
		Email:      claims.Email,
		GivenName:  claims.GivenName,
		FamilyName: claims.FamilyName,
		Picture:    claims.Picture,
		Language:   claims.Language,
		Theme:      claims.Theme,
		Timezone:   claims.Timezone,
	}

	token := lib.NewWithClaims(p.algorithm, custom)
	token.Header["typ"] = string(kind)

	return token.SignedString(p.signKey)
}

func (p *Provider) Create(
	sessionID string,
	claims jwt.Claims,
	optionalIncludeRefresh ...bool,
) (*jwt.Token, error) {
	includeRefresh := len(optionalIncludeRefresh) > 0 && optionalIncludeRefresh[0]
	now := time.Now()

	accessToken, err := p.createToken(sessionID, claims, jwt.Kind_ACCESS, now, p.accessTTL)
	if err != nil {
		return nil, err
	}

	token := &jwt.Token{
		TokenType:   "Bearer",
		AccessToken: accessToken,
		ExpiresIn:   int64(p.accessTTL.Seconds()),
	}

	if includeRefresh {
		refreshToken, err := p.createToken(sessionID, claims, jwt.Kind_REFRESH, now, p.refreshTTL)
		if err != nil {
			return nil, err
		}
		token.RefreshToken = refreshToken
	}

	return token, nil
}

func (p *Provider) Decode(
	tokenString string,
) (*jwt.Decoded, error) {
	token, err := lib.ParseWithClaims(tokenString, &customClaims{}, func(t *lib.Token) (any, error) {
		return p.verifyKey, nil
	}, lib.WithValidMethods([]string{p.algorithm.Alg()}),
		lib.WithIssuer(p.issuer),
		lib.WithAudience(p.audience))

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*customClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	kind := jwt.Kind_ACCESS
	if typ, ok := token.Header["typ"].(string); ok {
		kind = jwt.Kind(typ)
	}

	return &jwt.Decoded{
		Kind:      kind,
		SessionID: claims.ID,
		Claims: jwt.Claims{
			Subject:    claims.Subject,
			Email:      claims.Email,
			GivenName:  claims.GivenName,
			FamilyName: claims.FamilyName,
			Picture:    claims.Picture,
			Language:   claims.Language,
			Theme:      claims.Theme,
			Timezone:   claims.Timezone,
		},
		IssuedAt:  claims.IssuedAt.Time,
		ExpiresAt: claims.ExpiresAt.Time,
	}, nil
}
