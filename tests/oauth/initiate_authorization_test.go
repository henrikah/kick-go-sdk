package kick_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk"
	"github.com/henrikah/kick-go-sdk/enums/kickscopes"
	"github.com/henrikah/kick-go-sdk/kickoauthtypes"
	"github.com/henrikah/kick-go-sdk/kickerrors"
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

	var validationError *kickerrors.ValidationError
	// Act

	authorizationData, err := client.InitiateAuthorization(redirectURI, state, scopes)

	// Assert
	if authorizationData != nil {
		t.Fatal("Expected authorizationData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &validationError) {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "redirectURI" {
		t.Fatalf("Expected error on field 'redirectURI', got '%s'", validationError.Field)
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

	var validationError *kickerrors.ValidationError
	// Act

	authorizationData, err := client.InitiateAuthorization(redirectURI, state, scopes)

	// Assert
	if authorizationData != nil {
		t.Fatal("Expected authorizationData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &validationError) {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "state" {
		t.Fatalf("Expected error on field 'state', got '%s'", validationError.Field)
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

	var validationError *kickerrors.ValidationError
	// Act

	authorizationData, err := client.InitiateAuthorization(redirectURI, state, scopes)

	// Assert
	if authorizationData != nil {
		t.Fatal("Expected authorizationData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &validationError) {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "scopes" {
		t.Fatalf("Expected error on field 'scopes', got '%s'", validationError.Field)
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
