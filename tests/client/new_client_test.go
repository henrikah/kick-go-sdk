package kick_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickerrors"
)

func Test_NewAPIClientMissingClientID_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	config := kickapitypes.APIClientConfig{
		ClientID:     "",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}

	var validationError *kickerrors.ValidationError
	// Act
	client, err := kick.NewAPIClient(config)

	// Assert
	if client != nil {
		t.Fatal("Expected client to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &validationError) {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "ClientID" {
		t.Fatalf("Expected error on field 'ClientID', got '%s'", validationError.Field)
	}
}

func Test_NewAPIClientMissingClientSecret_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	config := kickapitypes.APIClientConfig{
		ClientID:     "test-id",
		ClientSecret: "",
		HTTPClient:   httpClient,
	}

	var validationError *kickerrors.ValidationError
	// Act
	client, err := kick.NewAPIClient(config)

	// Assert
	if client != nil {
		t.Fatal("Expected client to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &validationError) {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "ClientSecret" {
		t.Fatalf("Expected error on field 'ClientSecret', got '%s'", validationError.Field)
	}
}

func Test_NewAPIClientMissingHTTPClient_Error(t *testing.T) {
	// Arrange
	config := kickapitypes.APIClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   nil,
	}

	var validationError *kickerrors.ValidationError
	// Act
	client, err := kick.NewAPIClient(config)

	// Assert
	if client != nil {
		t.Fatal("Expected client to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &validationError) {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "HTTPClient" {
		t.Fatalf("Expected error on field 'HTTPClient', got '%s'", validationError.Field)
	}
}

func Test_NewAPIClient_Success(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	config := kickapitypes.APIClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}

	// Act
	client, err := kick.NewAPIClient(config)

	// Assert
	if client == nil {
		t.Fatal("Expected client to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}
}
