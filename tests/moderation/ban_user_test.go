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

func Test_BanUserMissingAccessToken_Error(t *testing.T) {
	// Arrange
	ctx := t.Context()
	httpClient := http.DefaultClient

	accessToken := ""
	broadcasterUserID := 1
	userID := 2
	reason := "some-reason"

	config := kickapitypes.APIClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	banUserData, err := client.Moderation().BanUser(ctx, accessToken, broadcasterUserID, userID, &reason)

	// Assert
	if banUserData != nil {
		t.Fatal("Expected banUserData to be nil")
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

func Test_BanUserInvalidBroadcasterUserID_Error(t *testing.T) {
	// Arrange
	ctx := t.Context()
	httpClient := http.DefaultClient

	accessToken := "access-token"
	broadcasterUserID := 0
	userID := 2
	reason := "some-reason"

	config := kickapitypes.APIClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	banUserData, err := client.Moderation().BanUser(ctx, accessToken, broadcasterUserID, userID, &reason)

	// Assert
	if banUserData != nil {
		t.Fatal("Expected banUserData to be nil")
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

func Test_BanUserInvalidUserID_Error(t *testing.T) {
	// Arrange
	ctx := t.Context()
	httpClient := http.DefaultClient

	accessToken := "access-token"
	broadcasterUserID := 1
	userID := 0
	reason := "some-reason"

	config := kickapitypes.APIClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	banUserData, err := client.Moderation().BanUser(ctx, accessToken, broadcasterUserID, userID, &reason)

	// Assert
	if banUserData != nil {
		t.Fatal("Expected banUserData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &validationError) {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "userID" {
		t.Fatalf("Expected error on field 'userID', got '%s'", validationError.Field)
	}
}

func Test_BanUserUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"message": "Invalid request"}`

	ctx := t.Context()

	accessToken := "access-token"
	broadcasterUserID := 1
	userID := 2
	reason := "some-reason"

	mockClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return mocks.NewMockResponse(http.StatusUnauthorized, errorJSON), nil
		},
	}

	config := kickapitypes.APIClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   mockClient,
	}
	client, _ := kick.NewAPIClient(config)

	var apiError *kickerrors.APIError
	// Act
	banUserData, err := client.Moderation().BanUser(ctx, accessToken, broadcasterUserID, userID, &reason)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if banUserData != nil {
		t.Fatal("Expected banUserData to be nil on error")
	}

	if !errors.As(err, &apiError) {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_BanUser_Success(t *testing.T) {
	// Arrange
	ctx := t.Context()
	clientID := "test-id"
	clientSecret := "test-secret"

	accessToken := "access-token"
	broadcasterUserID := 1
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

			if _, ok := updateData["duration"]; ok {
				t.Fatal("Expected duration to be omitted")
			}

			return mocks.NewMockResponse(http.StatusOK, expectedJSON), nil
		},
	}

	config := kickapitypes.APIClientConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HTTPClient:   httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act

	banUserData, err := client.Moderation().BanUser(ctx, accessToken, broadcasterUserID, userID, reason)

	// Assert
	if banUserData == nil {
		t.Fatal("Expected banUserData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if banUserData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", banUserData.Message)
	}
}

func Test_BanUserWithReason_Success(t *testing.T) {
	// Arrange
	ctx := t.Context()
	clientID := "test-id"
	clientSecret := "test-secret"

	accessToken := "access-token"
	broadcasterUserID := 1
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
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HTTPClient:   httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act

	banUserData, err := client.Moderation().BanUser(ctx, accessToken, broadcasterUserID, userID, &reason)

	// Assert
	if banUserData == nil {
		t.Fatal("Expected banUserData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if banUserData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", banUserData.Message)
	}
}
