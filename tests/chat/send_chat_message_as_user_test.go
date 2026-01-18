package kick_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickerrors"
	"github.com/henrikah/kick-go-sdk/tests/mocks"
)

func Test_SendChatMessageAsUserMissingAccessToken_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := ""
	broadcasterUserID := 1
	replyToMessageID := "test-reply-message"
	message := "test-message"

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	sendChatMessageData, err := client.Chat().SendChatMessageAsUser(t.Context(), accessToken, broadcasterUserID, &replyToMessageID, message)

	// Assert
	if sendChatMessageData != nil {
		t.Fatal("Expected sendChatMessageData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &validationError) {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "accessToken" {
		t.Fatalf("Expected error on field 'accessToken', got '%s'", validationError.Field)
	}
}

func Test_SendChatMessageAsUserInvalidBroadcasterUserID_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := "access-token"
	broadcasterUserID := 0
	replyToMessageID := "test-reply-message"
	message := "test-message"

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	sendChatMessageData, err := client.Chat().SendChatMessageAsUser(t.Context(), accessToken, broadcasterUserID, &replyToMessageID, message)

	// Assert
	if sendChatMessageData != nil {
		t.Fatal("Expected sendChatMessageData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &validationError) {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "broadcasterUserID" {
		t.Fatalf("Expected error on field 'broadcasterUserID', got '%s'", validationError.Field)
	}
}

func Test_SendChatMessageAsUserEmptyMessage_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := "access-token"
	broadcasterUserID := 1
	replyToMessageID := "test-reply-message"
	message := ""

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	sendChatMessageData, err := client.Chat().SendChatMessageAsUser(t.Context(), accessToken, broadcasterUserID, &replyToMessageID, message)

	// Assert
	if sendChatMessageData != nil {
		t.Fatal("Expected sendChatMessageData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &validationError) {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "message" {
		t.Fatalf("Expected error on field 'message', got '%s'", validationError.Field)
	}
}

func Test_SendChatMessageAsUserUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"message": "Invalid request"}`

	accessToken := "access-token"
	broadcasterUserID := 1
	replyToMessageID := "test-reply-message"
	message := "test-message"

	mockClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return mocks.NewMockResponse(http.StatusUnauthorized, errorJSON), nil
		},
	}

	config := kickapitypes.APIClientConfig{
		HTTPClient: mockClient,
	}
	client, _ := kick.NewAPIClient(config)

	var apiError *kickerrors.APIError
	// Act
	sendChatMessageData, err := client.Chat().SendChatMessageAsUser(t.Context(), accessToken, broadcasterUserID, &replyToMessageID, message)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if sendChatMessageData != nil {
		t.Fatal("Expected sendChatMessageData to be nil on error")
	}

	if !errors.As(err, &apiError) {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_SendChatMessageAsUser_Success(t *testing.T) {
	// Arrange
	accessToken := "access-token"
	broadcasterUserID := 1
	var replyToMessageID *string = nil
	message := "test-message"

	expectedJSON := `{
		"data": {
			"is_sent": true,
			"message_id": "test-message-id"
		},
		"message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/chat" {
				t.Fatalf("Unexpected request URL: %s", req.URL.String())
			}

			if req.Method != "POST" {
				t.Fatalf("Unexpected request method: %s", req.Method)
			}

			if req.Header.Get("Accept") != "application/json" {
				t.Fatal("Missing Accept header")
			}

			if req.Header.Get("Authorization") != "Bearer "+accessToken {
				t.Fatal("Missing Authorization header")
			}
			t.Cleanup(func() {
				defer func() {
					if err := req.Body.Close(); err != nil {
						t.Logf("failed to close request body: %v", err)
					}
				}()
			})

			bodyDecoder := json.NewDecoder(req.Body)

			var updateData map[string]any

			err := bodyDecoder.Decode(&updateData)

			if err != nil {
				return mocks.NewMockResponse(http.StatusInternalServerError, ""), nil
			}

			if _, ok := updateData["reply_to_message_id"]; ok {
				t.Fatal("Expected reply_to_message_id to be omitted")
			}

			return mocks.NewMockResponse(http.StatusOK, expectedJSON), nil
		},
	}

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act

	sendChatMessageData, err := client.Chat().SendChatMessageAsUser(t.Context(), accessToken, broadcasterUserID, replyToMessageID, message)

	// Assert
	if sendChatMessageData == nil {
		t.Fatal("Expected sendChatMessageData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if !sendChatMessageData.Data.IsSent {
		t.Fatalf("Expected IsSent to be true, got %t", sendChatMessageData.Data.IsSent)
	}

	if sendChatMessageData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", sendChatMessageData.Message)
	}
}

func Test_SendChatMessageAsUserWithReplyToMessageID_Success(t *testing.T) {
	// Arrange
	accessToken := "access-token"
	broadcasterUserID := 1
	replyToMessageID := "test-reply-message"
	message := "test-message"

	expectedJSON := `{
		"data": {
			"is_sent": true,
			"message_id": "test-message-id"
		},
		"message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/chat" {
				t.Fatalf("Unexpected request URL: %s", req.URL.String())
			}

			if req.Method != "POST" {
				t.Fatalf("Unexpected request method: %s", req.Method)
			}

			if req.Header.Get("Accept") != "application/json" {
				t.Fatal("Missing Accept header")
			}

			if req.Header.Get("Authorization") != "Bearer "+accessToken {
				t.Fatal("Missing Authorization header")
			}

			return mocks.NewMockResponse(http.StatusOK, expectedJSON), nil
		},
	}

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act

	sendChatMessageData, err := client.Chat().SendChatMessageAsUser(t.Context(), accessToken, broadcasterUserID, &replyToMessageID, message)

	// Assert
	if sendChatMessageData == nil {
		t.Fatal("Expected sendChatMessageData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if !sendChatMessageData.Data.IsSent {
		t.Fatalf("Expected IsSent to be true, got %t", sendChatMessageData.Data.IsSent)
	}

	if sendChatMessageData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", sendChatMessageData.Message)
	}
}
