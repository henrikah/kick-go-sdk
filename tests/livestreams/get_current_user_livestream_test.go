package kick_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickerrors"
	"github.com/henrikah/kick-go-sdk/tests/mocks"
)

func Test_GetCurrentUserLivestreamMissingAccessToken_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := ""

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	livestreamData, err := client.Livestream().GetCurrentUserLivestream(t.Context(), accessToken)

	// Assert
	if livestreamData != nil {
		t.Fatal("Expected livestreamData to be nil")
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

func Test_GetCurrentUserLivestreamUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"message": "Invalid request"}`

	accessToken := "access-token"

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
	livestreamData, err := client.Livestream().GetCurrentUserLivestream(t.Context(), accessToken)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if livestreamData != nil {
		t.Fatal("Expected livestreamData to be nil on error")
	}

	if !errors.As(err, &apiError) {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_GetCurrentUserLivestream_Success(t *testing.T) {
	// Arrange
	accessToken := "access-token"

	expectedJSON := `{
		"data": [{
			"broadcaster_user_id": 1,
			"category": {
				"id": 42,
				"name": "category-42",
				"thumbnail": "https://category-42"
			},
			"channel_id": 1,
			"custom_tags": ["some-tag"],
			"has_mature_content": true,
			"language": "en",
			"slug": "slug-1",
			"started_at": "started-at-1",
			"stream_title": "stream-title-1",
			"thumbnail": "thumbnail-1",
			"viewer_count": 5
		}],
		"message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/livestreams/stats" {
				t.Fatalf("Unexpected request URL: %s", req.URL.String())
			}

			if req.Method != "GET" {
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
	livestreamData, err := client.Livestream().GetCurrentUserLivestream(t.Context(), accessToken)

	// Assert
	if livestreamData == nil {
		t.Fatal("Expected livestreamData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if len(livestreamData.Data) != 1 {
		t.Fatalf("Expected Data to be 1 slice, got %d", len(livestreamData.Data))
	}

	if livestreamData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", livestreamData.Message)
	}
}
