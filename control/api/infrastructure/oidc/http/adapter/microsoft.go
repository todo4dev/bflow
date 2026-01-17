// infrastructure/oidc/adapter/microsoft.go
package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"src/port/oidc"
)

type HttpMicrosoft struct {
	httpBase
}

func NewHttpMicrosoft(baseURI string, clientID string, clientSecret string, callbackURI string) *HttpMicrosoft {
	return &HttpMicrosoft{
		httpBase: httpBase{
			clientID:     clientID,
			clientSecret: clientSecret,
			redirectURI:  baseURI + callbackURI,
			authorizeURL: "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
			tokenURL:     "https://login.microsoftonline.com/common/oauth2/v2.0/token",
			userInfoURL:  "https://graph.microsoft.com/oidc/userinfo",
			scopes:       "openid profile email offline_access User.Read",
		},
	}
}

func (m *HttpMicrosoft) GetAuthURL(state string, _ []string) string {
	return m.httpBase.GetAuthURL(state, map[string]string{"response_mode": "query"})
}

func (m *HttpMicrosoft) GetInfo(ctx context.Context, accessToken string) (*oidc.Claims, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", m.userInfoURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Microsoft custom fields variations
	var data map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	claims := &oidc.Claims{
		Custom: make(map[string]any),
	}

	if v, ok := data["sub"].(string); ok {
		claims.Subject = v
	}
	if v, ok := data["email"].(string); ok {
		claims.Email = v
	}
	if v, ok := data["name"].(string); ok {
		claims.Name = v
	}
	if v, ok := data["picture"].(string); ok {
		claims.Picture = v
	}

	// Microsoft specific fallbacks
	givenName := ""
	if v, ok := data["given_name"].(string); ok {
		givenName = v
	} else if v, ok := data["givenname"].(string); ok {
		givenName = v
	} else if v, ok := data["givenName"].(string); ok {
		givenName = v
	}
	claims.Custom["given_name"] = givenName

	familyName := ""
	if v, ok := data["family_name"].(string); ok {
		familyName = v
	} else if v, ok := data["familyname"].(string); ok {
		familyName = v
	} else if v, ok := data["familyName"].(string); ok {
		familyName = v
	}
	claims.Custom["family_name"] = familyName

	return claims, nil
}

func (m *HttpMicrosoft) GetPicture(ctx context.Context, accessToken string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://graph.microsoft.com/v1.0/me/photo/$value", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("failed to get picture: status %d", resp.StatusCode)
	}

	return resp.Body, nil
}
