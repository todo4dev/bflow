package adapter

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"src/port/oidc"
)

type HttpGoogle struct {
	httpBase
}

func NewHttpGoogle(baseURI string, clientID string, clientSecret string, callbackURI string) *HttpGoogle {
	return &HttpGoogle{
		httpBase: httpBase{
			clientID:     clientID,
			clientSecret: clientSecret,
			redirectURI:  baseURI + callbackURI,
			authorizeURL: "https://accounts.google.com/o/oauth2/v2/auth",
			tokenURL:     "https://oauth2.googleapis.com/token",
			userInfoURL:  "https://openidconnect.googleapis.com/v1/userinfo",
			scopes:       "openid profile email",
		},
	}
}

func (g *HttpGoogle) GetAuthURL(state string, _ []string) string {
	return g.httpBase.GetAuthURL(state, map[string]string{
		"access_type":            "offline",
		"include_granted_scopes": "false",
		"prompt":                 "consent",
	})
}

func (g *HttpGoogle) GetInfo(ctx context.Context, accessToken string) (*oidc.Claims, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", g.userInfoURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Sub           string `json:"sub"`
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
		Name          string `json:"name"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Picture       string `json:"picture"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &oidc.Claims{
		Subject:       data.Sub,
		Email:         data.Email,
		EmailVerified: data.EmailVerified,
		Name:          data.Name,
		Picture:       data.Picture,
		Custom: map[string]any{
			"given_name":  data.GivenName,
			"family_name": data.FamilyName,
		},
	}, nil
}

func (g *HttpGoogle) GetPicture(ctx context.Context, accessToken string) (io.ReadCloser, error) {
	claims, err := g.GetInfo(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "GET", claims.Picture, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
