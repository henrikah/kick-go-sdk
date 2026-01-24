package kick_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk/v2"
	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
	"github.com/henrikah/kick-go-sdk/v2/tests/mocks"
)

func Test_TimeOutUserMissingAccessToken_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := ""
	broadcasterUserID := 1
	duration := 60
	userID := 2
	reason := "some-reason"

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	// Act
	timeOutUserData, err := client.Moderation().TimeOutUser(t.Context(), accessToken, broadcasterUserID, userID, duration, &reason)

	// Assert
	if timeOutUserData != nil {
		t.Fatal("Expected timeOutUserData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationError := kickerrors.IsValidationError(err)

	if validationError == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "accessToken" {
		t.Fatalf("Expected error on field 'accessToken', got '%s'", validationError.Field)
	}
}

func Test_TimeOutUserInvalidBroadcasterUserID_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := "access-token"
	broadcasterUserID := 0
	duration := 60
	userID := 2
	reason := "some-reason"

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	// Act
	timeOutUserData, err := client.Moderation().TimeOutUser(t.Context(), accessToken, broadcasterUserID, userID, duration, &reason)

	// Assert
	if timeOutUserData != nil {
		t.Fatal("Expected timeOutUserData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationError := kickerrors.IsValidationError(err)

	if validationError == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "broadcasterUserID" {
		t.Fatalf("Expected error on field 'broadcasterUserID', got '%s'", validationError.Field)
	}
}

func Test_TimeOutUserDurationTooLow_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := "access-token"
	broadcasterUserID := 1
	duration := 0
	userID := 2
	reason := "some-reason"

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	// Act
	timeOutUserData, err := client.Moderation().TimeOutUser(t.Context(), accessToken, broadcasterUserID, userID, duration, &reason)

	// Assert
	if timeOutUserData != nil {
		t.Fatal("Expected timeOutUserData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationError := kickerrors.IsValidationError(err)

	if validationError == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "duration" {
		t.Fatalf("Expected error on field 'duration', got '%s'", validationError.Field)
	}
}

func Test_TimeOutUserDurationTooHigh_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := "access-token"
	broadcasterUserID := 1
	duration := 10081
	userID := 2
	reason := "some-reason"

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	// Act
	timeOutUserData, err := client.Moderation().TimeOutUser(t.Context(), accessToken, broadcasterUserID, userID, duration, &reason)

	// Assert
	if timeOutUserData != nil {
		t.Fatal("Expected timeOutUserData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationError := kickerrors.IsValidationError(err)

	if validationError == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "duration" {
		t.Fatalf("Expected error on field 'duration', got '%s'", validationError.Field)
	}
}

func Test_TimeOutUserInvalidUserID_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := "access-token"
	broadcasterUserID := 1
	duration := 60
	userID := 0
	reason := "some-reason"

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	// Act
	timeOutUserData, err := client.Moderation().TimeOutUser(t.Context(), accessToken, broadcasterUserID, userID, duration, &reason)

	// Assert
	if timeOutUserData != nil {
		t.Fatal("Expected timeOutUserData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationError := kickerrors.IsValidationError(err)

	if validationError == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "userID" {
		t.Fatalf("Expected error on field 'userID', got '%s'", validationError.Field)
	}
}

func Test_TimeOutUserUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"message": "Invalid request"}`

	accessToken := "access-token"
	broadcasterUserID := 1
	duration := 60
	userID := 2
	reason := "some-reason"

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
	timeOutUserData, err := client.Moderation().TimeOutUser(t.Context(), accessToken, broadcasterUserID, userID, duration, &reason)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if timeOutUserData != nil {
		t.Fatal("Expected timeOutUserData to be nil on error")
	}

	apiError := kickerrors.IsAPIError(err)

	if apiError == nil {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_TimeOutUser_Success(t *testing.T) {
	// Arrange
	accessToken := "access-token"
	broadcasterUserID := 1
	duration := 60
	userID := 2
	var reason *string = nil

	expectedJSON := `{
		"data": {},
		"message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/moderation/bans" {
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

			if _, ok := updateData["reason"]; ok {
				t.Fatal("Expected reason to be omitted")
			}

			return mocks.NewMockResponse(http.StatusOK, expectedJSON), nil
		},
	}

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act

	timeOutUserData, err := client.Moderation().TimeOutUser(t.Context(), accessToken, broadcasterUserID, userID, duration, reason)

	// Assert
	if timeOutUserData == nil {
		t.Fatal("Expected timeOutUserData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if timeOutUserData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", timeOutUserData.Message)
	}
}

func Test_TimeOutUserWithReason_Success(t *testing.T) {
	// Arrange
	accessToken := "access-token"
	broadcasterUserID := 1
	duration := 60
	userID := 2
	reason := "test-reason"

	expectedJSON := `{
		"data": {},
		"message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/moderation/bans" {
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

	timeOutUserData, err := client.Moderation().TimeOutUser(t.Context(), accessToken, broadcasterUserID, userID, duration, &reason)

	// Assert
	if timeOutUserData == nil {
		t.Fatal("Expected timeOutUserData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if timeOutUserData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", timeOutUserData.Message)
	}
}
