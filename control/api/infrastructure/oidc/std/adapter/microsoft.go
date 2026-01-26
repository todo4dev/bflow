// infrastructure/oidc/std/adapter/microsoft.go
package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"src/port/oidc"
)

func GetMapValueWithDefault[T any](m map[string]any, defaultValue T, keys ...string) T {
	for _, key := range keys {
		if v, ok := m[key]; ok {
			return v.(T)
		}
	}
	return defaultValue
}

type HttpMicrosoft struct {
	httpBase
}

var _ oidc.Adapter = (*HttpMicrosoft)(nil)

func NewHttpMicrosoft(
	baseURI string,
	clientID string,
	clientSecret string,
	callbackURI string,
) *HttpMicrosoft {
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

func (m *HttpMicrosoft) GetAuthURL(state map[string]any) string {
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

	var data map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	claims := &oidc.Claims{
		Subject:    GetMapValueWithDefault(data, "", "sub"),
		Email:      GetMapValueWithDefault(data, "", "email"),
		GivenName:  GetMapValueWithDefault(data, "", "given_name", "givenname", "givenName"),
		FamilyName: GetMapValueWithDefault(data, "", "family_name", "familyname", "familyName"),
	}

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
