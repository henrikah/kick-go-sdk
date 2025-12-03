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

func Test_GetChannelRewardsMissingAccessToken_Error(t *testing.T) {
	// Arrange
	ctx := t.Context()
	httpClient := http.DefaultClient

	accessToken := ""

	config := kickapitypes.APIClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	channelRewardsData, err := client.ChannelReward().GetChannelRewards(ctx, accessToken)

	// Assert
	if channelRewardsData != nil {
		t.Fatal("Expected channelRewardsData to be nil")
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

func Test_GetChannelRewardsUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"message": "Invalid request"}`

	ctx := t.Context()

	accessToken := "access-token"

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
	channelRewardsData, err := client.ChannelReward().GetChannelRewards(ctx, accessToken)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if channelRewardsData != nil {
		t.Fatal("Expected channelsData to be nil on error")
	}

	if !errors.As(err, &apiError) {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_GetChannelsByBroadcasterUserID_Success(t *testing.T) {
	// Arrange
	ctx := t.Context()
	clientID := "test-id"
	clientSecret := "test-secret"

	accessToken := "access-token"

	expectedJSON := `{
		"data": [{
			"background_color": "#123456",
			"cost": 100,
			"description": "description-1",
			"id": "reward-id-1",
			"is_enabled": true,
			"is_paused": false,
			"is_user_input_required":  true,
			"should_redemptions_skip_request_queue": true,
			"title": "Title-1"
		},
		{
			"background_color": "#654321",
			"cost": 101,
			"description": "description-2",
			"id": "reward-id-2",
			"is_enabled": true,
			"is_paused": false,
			"is_user_input_required":  true,
			"should_redemptions_skip_request_queue": true,
			"title": "Title-2"
		}],
		"message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/channels/rewards" {
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

	channelRewardsData, err := client.ChannelReward().GetChannelRewards(ctx, accessToken)

	// Assert
	if channelRewardsData == nil {
		t.Fatal("Expected channelRewardsData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if len(channelRewardsData.Data) != 2 {
		t.Fatalf("Expected Data to be 2 slices, got %d", len(channelRewardsData.Data))
	}

	if channelRewardsData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", channelRewardsData.Message)
	}
}
