package kick_test

import (
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk/v2"
	"github.com/henrikah/kick-go-sdk/v2/enums/kickscopes"
	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
	"github.com/henrikah/kick-go-sdk/v2/kickoauthtypes"
)

func Test_InitiateAuthorizationMissingRedirectURI_Error(t *testing.T) {
	// Arrange

	redirectURI := ""
	state := "unique-state"
	scopes := kickscopes.Scopes{kickscopes.UserRead, kickscopes.ChannelRead}

	httpClient := http.DefaultClient

	config := kickoauthtypes.OAuthClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}

	client, _ := kick.NewOAuthClient(config)

	// Act

	authorizationData, err := client.InitiateAuthorization(redirectURI, state, scopes)

	// Assert
	if authorizationData != nil {
		t.Fatal("Expected authorizationData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationErr := kickerrors.IsValidationError(err)

	if validationErr == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationErr.Field != "redirectURI" {
		t.Fatalf("Expected error on field 'redirectURI', got '%s'", validationErr.Field)
	}
}
func Test_InitiateAuthorizationMissingState_Error(t *testing.T) {
	// Arrange

	redirectURI := "http://localhost"
	state := ""
	scopes := kickscopes.Scopes{kickscopes.UserRead, kickscopes.ChannelRead}

	httpClient := http.DefaultClient

	config := kickoauthtypes.OAuthClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}

	client, _ := kick.NewOAuthClient(config)

	// Act

	authorizationData, err := client.InitiateAuthorization(redirectURI, state, scopes)

	// Assert
	if authorizationData != nil {
		t.Fatal("Expected authorizationData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationErr := kickerrors.IsValidationError(err)

	if validationErr == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationErr.Field != "state" {
		t.Fatalf("Expected error on field 'state', got '%s'", validationErr.Field)
	}
}
func Test_InitiateAuthorizationMissingScopes_Error(t *testing.T) {
	// Arrange

	redirectURI := "http://localhost"
	state := "unique-state"
	scopes := kickscopes.Scopes{}

	httpClient := http.DefaultClient

	config := kickoauthtypes.OAuthClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}

	client, _ := kick.NewOAuthClient(config)

	// Act

	authorizationData, err := client.InitiateAuthorization(redirectURI, state, scopes)

	// Assert
	if authorizationData != nil {
		t.Fatal("Expected authorizationData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationErr := kickerrors.IsValidationError(err)

	if validationErr == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationErr.Field != "scopes" {
		t.Fatalf("Expected error on field 'scopes', got '%s'", validationErr.Field)
	}
}
func Test_InitiateAuthorization_Success(t *testing.T) {
	// Arrange

	redirectURI := "http://localhost"
	state := "unique-state"
	scopes := kickscopes.Scopes{kickscopes.UserRead, kickscopes.ChannelRead}

	httpClient := http.DefaultClient

	config := kickoauthtypes.OAuthClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}

	client, _ := kick.NewOAuthClient(config)

	// Act

	authorizationData, err := client.InitiateAuthorization(redirectURI, state, scopes)

	// Assert
	if authorizationData == nil {
		t.Fatal("Expected authorizationData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if authorizationData.AuthorizationURL == "" {
		t.Fatal("Expected AuthorizationURL to not be empty")
	}

	if authorizationData.PKCEVerifier == "" {
		t.Fatal("Expected PKCEVerifier to not be empty")
	}
}
