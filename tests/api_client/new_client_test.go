package kick_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk/v2"
	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
)

func Test_NewAPIClientMissingHTTPClient_Error(t *testing.T) {
	// Arrange
	config := kickapitypes.APIClientConfig{
		HTTPClient: nil,
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

	if validationError.Field != "httpClient" {
		t.Fatalf("Expected error on field 'httpClient', got '%s'", validationError.Field)
	}
}

func Test_NewAPIClient_Success(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
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
