package kick

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/henrikah/kick-go-sdk/v2/enums/kickscopes"
	"github.com/henrikah/kick-go-sdk/v2/internal/auth"
	"github.com/henrikah/kick-go-sdk/v2/internal/endpoints"
	"github.com/henrikah/kick-go-sdk/v2/internal/transport"
	"github.com/henrikah/kick-go-sdk/v2/kickcontracts"
	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
	"github.com/henrikah/kick-go-sdk/v2/kickoauthtypes"
)

type oAuthClient struct {
	clientID     string
	clientSecret string
	requester    *transport.Requester
}

// NewOAuthClient creates a new NewOAuthClient instance with the provided configuration.
//
// Example:
//
//	oauthClient, err := kick.NewOAuthClient(kickapitypes.OAuthClientConfig{
//		ClientID:     "your-client-id",
//		ClientSecret: "your-client-secret",
//		HTTPClient:   http.DefaultClient,
//	})
//	if err != nil {
//		log.Fatalf("could not create OAuthClient: %v", err)
//	}
func NewOAuthClient(clientConfig kickoauthtypes.OAuthClientConfig) (kickcontracts.OAuth, error) {
	if err := kickerrors.ValidateNotEmpty("ClientID", clientConfig.ClientID); err != nil {
		return nil, err
	}
	if err := kickerrors.ValidateNotEmpty("ClientSecret", clientConfig.ClientSecret); err != nil {
		return nil, err
	}

	requester, err := transport.NewRequester(clientConfig.HTTPClient)
	if err != nil {
		return nil, err
	}

	client := &oAuthClient{
		clientID:     clientConfig.ClientID,
		clientSecret: clientConfig.ClientSecret,
		requester:    requester,
	}

	return client, nil
}

func (c *oAuthClient) InitiateAuthorization(redirectURI string, state string, scopes kickscopes.Scopes) (*kickoauthtypes.InitiateAuthorizationData, error) {
	if err := kickerrors.ValidateNotEmpty("redirectURI", redirectURI); err != nil {
		return nil, err
	}
	if err := kickerrors.ValidateNotEmpty("state", state); err != nil {
		return nil, err
	}
	if err := kickerrors.ValidateMinItems("scopes", scopes, 1); err != nil {
		return nil, err
	}

	pkce, err := auth.GeneratePKCE()
	if err != nil {
		return nil, err
	}

	authorizationURL, err := url.Parse(endpoints.UserAuthorizationURL())
	if err != nil {
		return nil, err
	}

	queryParams := authorizationURL.Query()
	queryParams.Set("client_id", c.clientID)
	queryParams.Set("response_type", "code")
	queryParams.Set("redirect_uri", redirectURI)
	queryParams.Set("state", state)
	queryParams.Set("scope", scopes.Join(" "))
	queryParams.Set("code_challenge", pkce.CodeChallenge)
	queryParams.Set("code_challenge_method", pkce.Method)

	authorizationURL.RawQuery = queryParams.Encode()

	return &kickoauthtypes.InitiateAuthorizationData{
		AuthorizationURL: authorizationURL.String(),
		PKCEVerifier:     pkce.CodeVerifier,
	}, nil
}

func (c *oAuthClient) ExchangeAuthorizationCode(ctx context.Context, redirectURI string, authorizationCode string, codeVerifier string) (*kickoauthtypes.CodeExchangeResponse, error) {
	if err := kickerrors.ValidateNotEmpty("redirectURI", redirectURI); err != nil {
		return nil, err
	}
	if err := kickerrors.ValidateNotEmpty("authorizationCode", authorizationCode); err != nil {
		return nil, err
	}
	if err := kickerrors.ValidateNotEmpty("codeVerifier", codeVerifier); err != nil {
		return nil, err
	}

	codeExchangeData := url.Values{}

	codeExchangeData.Set("code", authorizationCode)
	codeExchangeData.Set("client_id", c.clientID)
	codeExchangeData.Set("client_secret", c.clientSecret)
	codeExchangeData.Set("redirect_uri", redirectURI)
	codeExchangeData.Set("grant_type", "authorization_code")
	codeExchangeData.Set("code_verifier", codeVerifier)

	var tokenData kickoauthtypes.CodeExchangeResponse

	if err := c.requester.MakeFormRequest(ctx, http.MethodPost, endpoints.CodeExchangeURL(), strings.NewReader(codeExchangeData.Encode()), nil, &tokenData); err != nil {
		return nil, err
	}

	return &tokenData, nil
}

func (c *oAuthClient) GetAppAccessToken(ctx context.Context) (*kickoauthtypes.AppAccessTokenResponse, error) {
	appAccessTokenData := url.Values{}
	appAccessTokenData.Set("client_id", c.clientID)
	appAccessTokenData.Set("client_secret", c.clientSecret)
	appAccessTokenData.Set("grant_type", "client_credentials")

	var tokenData kickoauthtypes.AppAccessTokenResponse

	if err := c.requester.MakeFormRequest(ctx, http.MethodPost, endpoints.CodeExchangeURL(), strings.NewReader(appAccessTokenData.Encode()), nil, &tokenData); err != nil {
		return nil, err
	}

	return &tokenData, nil
}

func (c *oAuthClient) RevokeAccessToken(ctx context.Context, accessToken string) error {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return err
	}
	return c.revokeToken(ctx, accessToken, "access_token")
}
func (c *oAuthClient) RevokeRefreshToken(ctx context.Context, refreshToken string) error {
	if err := kickerrors.ValidateNotEmpty("refreshToken", refreshToken); err != nil {
		return err
	}
	return c.revokeToken(ctx, refreshToken, "refresh_token")
}

func (c *oAuthClient) revokeToken(ctx context.Context, token string, tokenType string) error {
	tokenData := url.Values{}
	tokenData.Set("token", token)
	tokenData.Set("token_hint_type", tokenType)

	if err := c.requester.MakeFormRequest(ctx, http.MethodPost, endpoints.RevokeTokenURL(), strings.NewReader(tokenData.Encode()), nil, nil); err != nil {
		return err
	}
	return nil
}

func (c *oAuthClient) TokenIntrospect(ctx context.Context, accessToken string) (*kickoauthtypes.TokenIntrospect, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}

	var tokenIntrospectData kickoauthtypes.TokenIntrospect

	if err := c.requester.MakeJSONRequest(ctx, http.MethodPost, endpoints.ViewTokenIntrospectURL(), nil, &accessToken, &tokenIntrospectData); err != nil {
		return nil, err
	}

	return &tokenIntrospectData, nil
}
