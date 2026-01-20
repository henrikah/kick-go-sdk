package kick_test

import (
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk/v2"
	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
	"github.com/henrikah/kick-go-sdk/v2/tests/mocks"
)

func Test_GetCurrentBroadcasterChannel_Success(t *testing.T) {
	// Arrange
	accessToken := "access-token"

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
		}],
		"message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/channels" {
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

	channelData, err := client.Channel().GetCurrentBroadcasterChannel(t.Context(), accessToken)

	// Assert
	if channelData == nil {
		t.Fatal("Expected channelsData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if len(channelData.Data) != 1 {
		t.Fatalf("Expected Data to be 1 slice, got %d", len(channelData.Data))
	}

	if channelData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", channelData.Message)
	}
}
