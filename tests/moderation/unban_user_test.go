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

func Test_UnbanUserMissingAccessToken_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := ""
	broadcasterUserID := 1
	userID := 2

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	// Act
	unbanUserData, err := client.Moderation().UnbanUser(t.Context(), accessToken, broadcasterUserID, userID)

	// Assert
	if unbanUserData != nil {
		t.Fatal("Expected unbanUserData to be nil")
	}

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

func Test_UnbanUserInvalidBroadcasterUserID_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := "access-token"
	broadcasterUserID := 0
	userID := 2

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	// Act
	unbanUserData, err := client.Moderation().UnbanUser(t.Context(), accessToken, broadcasterUserID, userID)

	// Assert
	if unbanUserData != nil {
		t.Fatal("Expected unbanUserData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationErr := kickerrors.IsValidationError(err)

	if validationErr == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationErr.Field != "broadcasterUserID" {
		t.Fatalf("Expected error on field 'broadcasterUserID', got '%s'", validationErr.Field)
	}
}

func Test_UnbanUserInvalidUserID_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := "access-token"
	broadcasterUserID := 1
	userID := 0

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	// Act
	unbanUserData, err := client.Moderation().UnbanUser(t.Context(), accessToken, broadcasterUserID, userID)

	// Assert
	if unbanUserData != nil {
		t.Fatal("Expected unbanUserData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationErr := kickerrors.IsValidationError(err)

	if validationErr == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationErr.Field != "userID" {
		t.Fatalf("Expected error on field 'userID', got '%s'", validationErr.Field)
	}
}

func Test_UnbanUserUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"message": "Invalid request"}`

	accessToken := "access-token"
	broadcasterUserID := 1
	userID := 2

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
	unbanUserData, err := client.Moderation().UnbanUser(t.Context(), accessToken, broadcasterUserID, userID)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if unbanUserData != nil {
		t.Fatal("Expected unbanUserData to be nil on error")
	}

	apiErr := kickerrors.IsAPIError(err)

	if apiErr == nil {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_UnbanUser_Success(t *testing.T) {
	// Arrange
	accessToken := "access-token"
	broadcasterUserID := 1
	userID := 2

	expectedJSON := `{
		"data": {},
		"message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/moderation/bans" {
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

			return mocks.NewMockResponse(http.StatusOK, expectedJSON), nil
		},
	}

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act

	unbanUserData, err := client.Moderation().UnbanUser(t.Context(), accessToken, broadcasterUserID, userID)

	// Assert
	if unbanUserData == nil {
		t.Fatal("Expected unbanUserData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if unbanUserData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", unbanUserData.Message)
	}
}
