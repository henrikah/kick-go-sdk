package kickcontracts

import (
	"context"

	"github.com/henrikah/kick-go-sdk/v2/enums/kickscopes"
	"github.com/henrikah/kick-go-sdk/v2/kickoauthtypes"
)

// OAuth handles OAuth2 flows for the Kick API.
//
// All examples use context.TODO() as a placeholder. Replace with a proper
// context (with timeout/cancel) in production code.
type OAuth interface {
	// InitiateAuthorization starts the PKCE OAuth2 flow by generating a code verifier and challenge,
	// then constructing the authorization URL.
	//
	// Example:
	//
	//	oAuth, err := kick.NewOAuthClient(kickoauthtypes.OAuthClientConfig{...})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	authData, err := oAuth.InitiateAuthorization("http://localhost", "random-state", kickscopes.Scopes{kickscope.UserRead})
	//	if err != nil {
	//		var apiErr *kickerrors.APIError
	//		if errors.As(err, &apiErr) {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	InitiateAuthorization(redirectURI string, state string, scopes kickscopes.Scopes) (*kickoauthtypes.InitiateAuthorizationData, error)

	// ExchangeAuthorizationCode exchanges an authorization code for an access token.
	//
	// Example:
	//
	//	oAuth, err := kick.NewOAuthClient(kickoauthtypes.OAuthClientConfig{...})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	tokenData, err := oAuth.ExchangeAuthorizationCode(context.TODO(), "http://localhost", "auth-code", "pkce-verifier")
	//	if err != nil {
	//		var apiErr *kickerrors.APIError
	//		if errors.As(err, &apiErr) {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	ExchangeAuthorizationCode(ctx context.Context, redirectURI, authorizationCode, codeVerifier string) (*kickoauthtypes.CodeExchangeResponse, error)

	// GetAppAccessToken requests an app access token for public data.
	//
	// Example:
	//
	//	oAuth, err := kick.NewOAuthClient(kickoauthtypes.OAuthClientConfig{...})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	tokenData, err := oAuth.GetAppAccessToken(context.TODO())
	//	if err != nil {
	//		var apiErr *kickerrors.APIError
	//		if errors.As(err, &apiErr) {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	GetAppAccessToken(ctx context.Context) (*kickoauthtypes.AppAccessTokenResponse, error)

	// RevokeAccessToken revokes a user's access token.
	//
	// Example:
	//
	//	oAuth, err := kick.NewOAuthClient(kickoauthtypes.OAuthClientConfig{...})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	err := oAuth.RevokeAccessToken(context.TODO(), accessToken)
	//	if err != nil {
	//		var apiErr *kickerrors.APIError
	//		if errors.As(err, &apiErr) {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	RevokeAccessToken(ctx context.Context, accessToken string) error

	// RevokeRefreshToken revokes a user's refresh token and all access tokens associated with it.
	//
	// Example:
	//
	//	oAuth, err := kick.NewOAuthClient(kickoauthtypes.OAuthClientConfig{...})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	err := oAuth.RevokeRefreshToken(context.TODO(), refreshToken)
	//	if err != nil {
	//		var apiErr *kickerrors.APIError
	//		if errors.As(err, &apiErr) {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	RevokeRefreshToken(ctx context.Context, refreshToken string) error

	// TokenIntrospect validates an access token and returns metadata such as
	// validity, expiration, and associated user/client.
	//
	// Example:
	//
	//	oAuth, err := kick.NewOAuthClient(kickoauthtypes.OAuthClientConfig{...})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	introspection, err := oAuth.TokenIntrospect(context.TODO(), accessToken)
	//	if err != nil {
	//		var apiErr *kickerrors.APIError
	//		if errors.As(err, &apiErr) {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	TokenIntrospect(ctx context.Context, accessToken string) (*kickoauthtypes.TokenIntrospect, error)
}
