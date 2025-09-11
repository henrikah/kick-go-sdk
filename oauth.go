package kick

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/henrikah/kick-go-sdk/enums/kickscopes"
	"github.com/henrikah/kick-go-sdk/internal/auth"
	"github.com/henrikah/kick-go-sdk/internal/endpoints"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickcontracts"
	"github.com/henrikah/kick-go-sdk/kickerrors"
)

type oAuthClient struct {
	client *apiClient
}

func newOAuthClient(client *apiClient) kickcontracts.OAuth {
	return &oAuthClient{
		client: client,
	}
}
func (c *oAuthClient) InitiateAuthorization(redirectURI string, state string, scopes kickscopes.Scopes) (*kickapitypes.InitiateAuthorizationData, error) {
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
	queryParams.Set("client_id", c.client.clientID)
	queryParams.Set("response_type", "code")
	queryParams.Set("redirect_uri", redirectURI)
	queryParams.Set("state", state)
	queryParams.Set("scope", scopes.Join(" "))
	queryParams.Set("code_challenge", pkce.CodeChallenge)
	queryParams.Set("code_challenge_method", pkce.Method)

	authorizationURL.RawQuery = queryParams.Encode()

	return &kickapitypes.InitiateAuthorizationData{
		AuthorizationURL: authorizationURL.String(),
		PKCEVerifier:     pkce.CodeVerifier,
	}, nil
}

func (c *oAuthClient) ExchangeAuthorizationCode(ctx context.Context, redirectURI string, authorizationCode string, codeVerifier string) (*kickapitypes.CodeExchangeResponse, error) {
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
	codeExchangeData.Set("client_id", c.client.clientID)
	codeExchangeData.Set("client_secret", c.client.clientSecret)
	codeExchangeData.Set("redirect_uri", redirectURI)
	codeExchangeData.Set("grant_type", "authorization_code")
	codeExchangeData.Set("code_verifier", codeVerifier)

	var tokenData kickapitypes.CodeExchangeResponse

	if err := c.client.makeFormRequest(ctx, http.MethodPost, endpoints.CodeExchangeURL(), strings.NewReader(codeExchangeData.Encode()), nil, &tokenData); err != nil {
		return nil, err
	}

	return &tokenData, nil
}

func (c *oAuthClient) GetAppAccessToken(ctx context.Context) (*kickapitypes.AppAccessTokenResponse, error) {
	appAccessTokenData := url.Values{}
	appAccessTokenData.Set("client_id", c.client.clientID)
	appAccessTokenData.Set("client_secret", c.client.clientSecret)
	appAccessTokenData.Set("grant_type", "client_credentials")

	var tokenData kickapitypes.AppAccessTokenResponse

	if err := c.client.makeFormRequest(ctx, http.MethodPost, endpoints.CodeExchangeURL(), strings.NewReader(appAccessTokenData.Encode()), nil, &tokenData); err != nil {
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

	if err := c.client.makeFormRequest(ctx, http.MethodPost, endpoints.RevokeTokenURL(), strings.NewReader(tokenData.Encode()), nil, nil); err != nil {
		return err
	}
	return nil
}
