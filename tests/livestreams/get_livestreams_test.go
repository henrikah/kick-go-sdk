package kick_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk"
	"github.com/henrikah/kick-go-sdk/enums/kicksortbyenum"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickerrors"
	"github.com/henrikah/kick-go-sdk/kickfilters"
	"github.com/henrikah/kick-go-sdk/tests/mocks"
)

func Test_GetLivestreamsMissingAccessToken_Error(t *testing.T) {
	// Arrange
	ctx := t.Context()
	httpClient := http.DefaultClient

	accessToken := ""
	filter := kickfilters.NewLivestreamFilterBuilder().WithBroadcasterUserIDs([]int64{1, 2}).WithCategoryID(42).WithLanguage("en").WithLimit(69).WithSortBy(kicksortbyenum.SortByViewerCount)

	config := kickapitypes.APIClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	livestreamsData, err := client.Livestream.SearchLivestreams(ctx, accessToken, filter)

	// Assert
	if livestreamsData != nil {
		t.Fatal("Expected livestreamsData to be nil")
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

func Test_GetLivestreamsUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"message": "Invalid request"}`

	ctx := t.Context()

	accessToken := "access-token"
	filter := kickfilters.NewLivestreamFilterBuilder().WithBroadcasterUserIDs([]int64{1, 2}).WithCategoryID(42).WithLanguage("en").WithLimit(69).WithSortBy(kicksortbyenum.SortByViewerCount)

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
	livestreamsData, err := client.Livestream.SearchLivestreams(ctx, accessToken, filter)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if livestreamsData != nil {
		t.Fatal("Expected livestreamsData to be nil on error")
	}

	if !errors.As(err, &apiError) {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_GetLivestreams_Success(t *testing.T) {
	// Arrange
	ctx := t.Context()
	clientID := "test-id"
	clientSecret := "test-secret"

	accessToken := "access-token"
	filter := kickfilters.NewLivestreamFilterBuilder().WithBroadcasterUserIDs([]int64{1, 2}).WithCategoryID(42).WithLanguage("en").WithLimit(69).WithSortBy(kicksortbyenum.SortByViewerCount)

	expectedJSON := `{
		"data": [{
			"broadcaster_user_id": 1,
			"category": {
				"id": 42,
				"name": "category-42",
				"thumbnail": "https://category-42"
			},
			"channel_id": 1,
			"has_mature_content": true,
			"language": "en",
			"slug": "slug-1",
			"started_at": "started-at-1",
			"stream_title": "stream-title-1",
			"thumbnail": "thumbnail-1",
			"viewer_count": 5
		},
		{
			"broadcaster_user_id": 2,
			"category": {
				"id": 42,
				"name": "category-42",
				"thumbnail": "https://category-42"
			},
			"channel_id": 2,
			"has_mature_content": true,
			"language": "en",
			"slug": "slug-2",
			"started_at": "started-at-2",
			"stream_title": "stream-title-2",
			"thumbnail": "thumbnail-2",
			"viewer_count": 7
		}],
		"message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/livestreams?broadcaster_user_id=1&broadcaster_user_id=2&category_id=42&language=en&limit=69&sort=viewer_count" {
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
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HTTPClient:   httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act

	livestreamsData, err := client.Livestream.SearchLivestreams(ctx, accessToken, filter)

	// Assert
	if livestreamsData == nil {
		t.Fatal("Expected livestreamsData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if len(livestreamsData.Data) != 2 {
		t.Fatalf("Expected Data to be 2 slices, got %d", len(livestreamsData.Data))
	}

	if livestreamsData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", livestreamsData.Message)
	}
}

func Test_GetLivestreamsWithoutFilter_Success(t *testing.T) {
	// Arrange
	ctx := t.Context()
	clientID := "test-id"
	clientSecret := "test-secret"

	accessToken := "access-token"
	var filter kickfilters.LivestreamFilterBuilder = nil

	expectedJSON := `{
		"data": [{
			"broadcaster_user_id": 1,
			"category": {
				"id": 42,
				"name": "category-42",
				"thumbnail": "https://category-42"
			},
			"channel_id": 1,
			"has_mature_content": true,
			"language": "en",
			"slug": "slug-1",
			"started_at": "started-at-1",
			"stream_title": "stream-title-1",
			"thumbnail": "thumbnail-1",
			"viewer_count": 5
		},
		{
			"broadcaster_user_id": 2,
			"category": {
				"id": 42,
				"name": "category-42",
				"thumbnail": "https://category-42"
			},
			"channel_id": 2,
			"has_mature_content": true,
			"language": "en",
			"slug": "slug-2",
			"started_at": "started-at-2",
			"stream_title": "stream-title-2",
			"thumbnail": "thumbnail-2",
			"viewer_count": 7
		}],
		"message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/livestreams" {
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
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HTTPClient:   httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act

	livestreamsData, err := client.Livestream.SearchLivestreams(ctx, accessToken, filter)

	// Assert
	if livestreamsData == nil {
		t.Fatal("Expected livestreamsData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if len(livestreamsData.Data) != 2 {
		t.Fatalf("Expected Data to be 2 slices, got %d", len(livestreamsData.Data))
	}

	if livestreamsData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", livestreamsData.Message)
	}
}
