package kickoauthtypes

import "github.com/henrikah/kick-go-sdk/v2/enums/kickscopes"

type AppAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

type CodeExchangeResponse struct {
	AccessToken  string            `json:"access_token"`
	TokenType    string            `json:"token_type"`
	RefreshToken string            `json:"refresh_token"`
	ExpiresIn    int64             `json:"expires_in"`
	Scope        kickscopes.Scopes `json:"scope"`
}

type InitiateAuthorizationData struct {
	AuthorizationURL string
	PKCEVerifier     string
}

type TokenIntrospectData struct {
	Active    bool              `json:"active,omitempty"`
	ClientID  string            `json:"client_id,omitempty"`
	Expires   int               `json:"exp,omitempty"`
	Scopes    kickscopes.Scopes `json:"scope,omitempty"`
	TokenType string            `json:"token_type,omitempty"`
}

type TokenIntrospect struct {
	Data    TokenIntrospectData `json:"data,omitempty"`
	Message string              `json:"message,omitempty"`
}
