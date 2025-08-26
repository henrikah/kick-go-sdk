package kickapitypes

import "github.com/henrikah/kick-go-sdk/enums/kickscopes"

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
