package kick_test

import (
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk/v2"
	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
	"github.com/henrikah/kick-go-sdk/v2/kickoauthtypes"
)

func Test_NewOAuthClientMissingClientID_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	config := kickoauthtypes.OAuthClientConfig{
		ClientID:     "",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}

	// Act
	client, err := kick.NewOAuthClient(config)

	// Assert
	if client != nil {
		t.Fatal("Expected client to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationError := kickerrors.IsValidationError(err)

	if validationError == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "ClientID" {
		t.Fatalf("Expected error on field 'ClientID', got '%s'", validationError.Field)
	}
}

func Test_NewOAuthClientMissingClientSecret_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	config := kickoauthtypes.OAuthClientConfig{
		ClientID:     "test-id",
		ClientSecret: "",
		HTTPClient:   httpClient,
	}

	// Act
	client, err := kick.NewOAuthClient(config)

	// Assert
	if client != nil {
		t.Fatal("Expected client to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationError := kickerrors.IsValidationError(err)

	if validationError == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "ClientSecret" {
		t.Fatalf("Expected error on field 'ClientSecret', got '%s'", validationError.Field)
	}
}

func Test_NewOAuthClientMissingHTTPClient_Error(t *testing.T) {
	// Arrange
	config := kickoauthtypes.OAuthClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   nil,
	}

	// Act
	client, err := kick.NewOAuthClient(config)

	// Assert
	if client != nil {
		t.Fatal("Expected client to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationError := kickerrors.IsValidationError(err)

	if validationError == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "httpClient" {
		t.Fatalf("Expected error on field 'httpClient', got '%s'", validationError.Field)
	}
}

func Test_NewOAuthClient_Success(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	config := kickoauthtypes.OAuthClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}

	// Act
	client, err := kick.NewOAuthClient(config)

	// Assert
	if client == nil {
		t.Fatal("Expected client to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}
}
