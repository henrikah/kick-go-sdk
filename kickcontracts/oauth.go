package kickcontracts

import (
	"context"

	"github.com/henrikah/kick-go-sdk/enums/kickscopes"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
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
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{...})
	//	if err != nil {
	//	    log.Printf("could not create APIClient: %v", err)
	//	    return nil, err
	//	}
	//
	//	authData, err := client.OAuth().InitiateAuthorization("http://localhost", "random-state", kickscopes.Scopes{kickscope.UserRead})
	//	if err != nil {
	//	    log.Printf("could not initiate authorization: %v", err)
	//	    return nil, err
	//	}
	InitiateAuthorization(redirectURI string, state string, scopes kickscopes.Scopes) (*kickapitypes.InitiateAuthorizationData, error)

	// ExchangeAuthorizationCode exchanges an authorization code for an access token.
	//
	// Example:
	//
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{...})
	//	if err != nil {
	//	    log.Printf("could not create APIClient: %v", err)
	//	    return nil, err
	//	}
	//
	//	tokenData, err := client.OAuth().ExchangeAuthorizationCode(context.TODO(), "http://localhost", "auth-code", "pkce-verifier")
	//	if err != nil {
	//	    log.Printf("could not exchange authorization code: %v", err)
	//	    return nil, err
	//	}
	ExchangeAuthorizationCode(ctx context.Context, redirectURI, authorizationCode, codeVerifier string) (*kickapitypes.CodeExchangeResponse, error)

	// GetAppAccessToken requests an app access token for public data.
	//
	// Example:
	//
	//	tokenData, err := client.OAuth().GetAppAccessToken(context.TODO())
	//	if err != nil {
	//	    log.Printf("could not get app access token: %v", err)
	//	    return nil, err
	//	}
	GetAppAccessToken(ctx context.Context) (*kickapitypes.AppAccessTokenResponse, error)

	// RevokeAccessToken revokes a user's access token.
	//
	// Example:
	//
	//	err := client.OAuth().RevokeAccessToken(context.TODO(), accessToken)
	//	if err != nil {
	//	    log.Printf("could not revoke access token: %v", err)
	//	}
	RevokeAccessToken(ctx context.Context, accessToken string) error

	// RevokeRefreshToken revokes a user's refresh token and all access tokens associated with it.
	//
	// Example:
	//
	//	err := client.OAuth().RevokeRefreshToken(context.TODO(), refreshToken)
	//	if err != nil {
	//	    log.Printf("could not revoke refresh token: %v", err)
	//	}
	RevokeRefreshToken(ctx context.Context, refreshToken string) error
}
