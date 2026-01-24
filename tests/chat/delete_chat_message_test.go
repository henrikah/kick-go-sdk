package kick_test

import (
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk/v2"
	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
	"github.com/henrikah/kick-go-sdk/v2/tests/mocks"
)

func Test_DeleteChatMessageMissingAccessToken_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := ""
	messageID := "test-message-id"

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	// Act
	err := client.Chat().DeleteChatMessage(t.Context(), accessToken, messageID)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationErr := kickerrors.IsValidationError(err)

	if validationErr == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationErr.Field != "accessToken" {
		t.Fatalf("Expected error on field 'accessToken', got '%s'", validationErr.Field)
	}
}

func Test_DeleteChatMessageEmptyMessageID_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := "access-token"
	messageID := ""

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	// Act
	err := client.Chat().DeleteChatMessage(t.Context(), accessToken, messageID)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationErr := kickerrors.IsValidationError(err)

	if validationErr == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationErr.Field != "messageID" {
		t.Fatalf("Expected error on field 'messageID', got '%s'", validationErr.Field)
	}
}

func Test_DeleteChatMessageUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := ``

	accessToken := "access-token"
	messageID := "test-message-id"

	mockClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return mocks.NewMockResponse(http.StatusUnauthorized, errorJSON), nil
		},
	}

	config := kickapitypes.APIClientConfig{
		HTTPClient: mockClient,
	}
	client, _ := kick.NewAPIClient(config)

	// Act
	err := client.Chat().DeleteChatMessage(t.Context(), accessToken, messageID)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	apiErr := kickerrors.IsAPIError(err)

	if apiErr == nil {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_DeleteChatMessage_Success(t *testing.T) {
	// Arrange
	accessToken := "access-token"
	messageID := "test-message-id"

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/chat/test-message-id" {
				t.Fatalf("Unexpected request URL: %s", req.URL.String())
			}

			if req.Method != "DELETE" {
				t.Fatalf("Unexpected request method: %s", req.Method)
			}

			if req.Header.Get("Accept") != "application/json" {
				t.Fatal("Missing Accept header")
			}

			if req.Header.Get("Authorization") != "Bearer "+accessToken {
				t.Fatal("Missing Authorization header")
			}

			return mocks.NewMockResponse(http.StatusNoContent, ""), nil
		},
	}

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act

	err := client.Chat().DeleteChatMessage(t.Context(), accessToken, messageID)

	// Assert
	if err != nil {
		t.Fatal("Expected error to be nil")
	}
}
