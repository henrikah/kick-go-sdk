package kick_test

import (
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk/v2"
	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
	"github.com/henrikah/kick-go-sdk/v2/tests/mocks"
)

func Test_GetChannelsByBroadcasterSlugMissingAccessToken_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := ""
	slugs := []string{"slug-1", "slug-2"}

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	// Act
	channelsData, err := client.Channel().GetChannelsByBroadcasterSlug(t.Context(), accessToken, slugs)

	// Assert
	if channelsData != nil {
		t.Fatal("Expected channelsData to be nil")
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

func Test_GetChannelsByBroadcasterSlugUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"message": "Invalid request"}`

	accessToken := "access-token"
	slugs := []string{"slug-1", "slug-2"}

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
	channelsData, err := client.Channel().GetChannelsByBroadcasterSlug(t.Context(), accessToken, slugs)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if channelsData != nil {
		t.Fatal("Expected usersData to be nil on error")
	}

	apiError := kickerrors.IsAPIError(err)
	if apiError == nil {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_GetChannelsByBroadcasterSlug_Success(t *testing.T) {
	// Arrange
	accessToken := "access-token"
	slugs := []string{"slug-1", "slug-2"}

	expectedJSON := `{
		"data": [{
			"banner_picture": "https://banner-test-user-1",
			"broadcaster_user_id": 1,
			"category": {
				"id": 1,
				"name": "category-1",
				"thumbnail": "https://category-1"
			},
			"channel_description": "channel-description-1",
			"slug": "user-1-slug",
			"stream": {
				"is_live": true,
				"is_mature": true,
				"key": "key-1",
				"language": "language-1",
				"start_time": "start-time-1",
				"thumbnail": "https://stream-1",
				"url": "https://url-1",
				"viewer_count": 5
			},
			"stream_title": "title-1"
		},
		{
			"banner_picture": "https://banner-test-user-2",
			"broadcaster_user_id": 2,
			"category": {
				"id": 2,
				"name": "category-2",
				"thumbnail": "https://category-2"
			},
			"channel_description": "channel-description-2",
			"slug": "user-2-slug",
			"stream": {
				"is_live": true,
				"is_mature": true,
				"key": "key-2",
				"language": "language-2",
				"start_time": "start-time-2",
				"thumbnail": "https://stream-2",
				"url": "https://url-2",
				"viewer_count": 7
			},
			"stream_title": "title-2"
		}],
		"message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/channels?slug=slug-1&slug=slug-2" {
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

	channelsData, err := client.Channel().GetChannelsByBroadcasterSlug(t.Context(), accessToken, slugs)

	// Assert
	if channelsData == nil {
		t.Fatal("Expected channelsData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if len(channelsData.Data) != 2 {
		t.Fatalf("Expected Data to be 2 slices, got %d", len(channelsData.Data))
	}

	if channelsData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", channelsData.Message)
	}
}
