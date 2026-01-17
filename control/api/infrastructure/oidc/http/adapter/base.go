// infrastructure/oidc/adapter/_base.go
package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"src/port/oidc"
	"strings"
)

type httpBase struct {
	clientID     string
	clientSecret string
	redirectURI  string
	authorizeURL string
	tokenURL     string
	userInfoURL  string
	scopes       string
}

func (b *httpBase) GetAuthURL(state string, extraParams map[string]string) string {
	u, _ := url.Parse(b.authorizeURL)
	q := u.Query()
	q.Set("client_id", b.clientID)
	q.Set("redirect_uri", b.redirectURI)
	q.Set("response_type", "code")
	q.Set("scope", b.scopes)
	q.Set("state", state)

	for k, v := range extraParams {
		q.Set(k, v)
	}

	u.RawQuery = q.Encode()
	return u.String()
}

func (b *httpBase) exchangeToken(ctx context.Context, payload url.Values) (*oidc.Tokens, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", b.tokenURL, strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("oidc token error: %s", string(body))
	}

	var data struct {
		AccessToken  string `json:"access_token"`
		IDToken      string `json:"id_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int64  `json:"expires_in"`
		TokenType    string `json:"token_type"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &oidc.Tokens{
		AccessToken:  data.AccessToken,
		IDToken:      data.IDToken,
		RefreshToken: data.RefreshToken,
		ExpiresIn:    data.ExpiresIn,
	}, nil
}

func (b *httpBase) Exchange(ctx context.Context, code string) (*oidc.Tokens, error) {
	return b.exchangeToken(ctx, url.Values{
		"client_id":     {b.clientID},
		"client_secret": {b.clientSecret},
		"code":          {code},
		"grant_type":    {"authorization_code"},
		"redirect_uri":  {b.redirectURI},
	})
}

func (b *httpBase) GetToken(ctx context.Context, refreshToken string) (*oidc.Tokens, error) {
	return b.exchangeToken(ctx, url.Values{
		"client_id":     {b.clientID},
		"client_secret": {b.clientSecret},
		"refresh_token": {refreshToken},
		"grant_type":    {"refresh_token"},
	})
}
