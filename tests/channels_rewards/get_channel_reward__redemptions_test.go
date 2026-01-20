package kick_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk/v2"
	"github.com/henrikah/kick-go-sdk/v2/enums/kickchannelrewardstatus"
	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
	"github.com/henrikah/kick-go-sdk/v2/kickfilters"
	"github.com/henrikah/kick-go-sdk/v2/tests/mocks"
)

func Test_GetChannelRewardRedemptionsMissingAccessToken_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := ""

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	channelRewardRedemptionsData, err := client.ChannelReward().GetChannelRewardRedemptions(t.Context(), accessToken, nil)

	// Assert
	if channelRewardRedemptionsData != nil {
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

func Test_GetChannelRewardRedemptionsUnAuthorized_Error(t *testing.T) {
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
	channelRewardRedemptionsData, err := client.ChannelReward().GetChannelRewardRedemptions(t.Context(), accessToken, nil)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if channelRewardRedemptionsData != nil {
		t.Fatal("Expected channelsData to be nil on error")
	}

	if !errors.As(err, &apiError) {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_GetChannelRewardRedemptions_Success(t *testing.T) {
	// Arrange
	accessToken := "access-token"

	expectedJSON := `{
		"data": [{
			"redemptions": [{
				"id": "redemption-id-1",
				"redeemed_at": "2025-12-19T06:26:36Z",
				"redeemer": {
					"broadcaster_user_id": 42
				},
				"status": "pending",
				"user_input": "Some message"
			}],
			"reward": {
				"can_manage": true,
				"cost": 101,
				"description": "description-1",
				"id": "reward-id-1",
				"is_deleted": false,
				"title": "Title-1"
			}
		}],
		"message": "test-message",
		"pagination": {
			"next_cursor": "some-cursor"
		}
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/channels/rewards/redemptions" {
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

	channelRewardRedemptionsData, err := client.ChannelReward().GetChannelRewardRedemptions(t.Context(), accessToken, nil)

	// Assert
	if channelRewardRedemptionsData == nil {
		t.Fatal("Expected channelRewardsData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if len(channelRewardRedemptionsData.Data) != 1 {
		t.Fatalf("Expected Data to be 1 slice, got %d", len(channelRewardRedemptionsData.Data))
	}

	if len(channelRewardRedemptionsData.Data[0].Redemptions) != 1 {
		t.Fatalf("Expected Data to be 1 slice, got %d", len(channelRewardRedemptionsData.Data[0].Redemptions))
	}

	if channelRewardRedemptionsData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", channelRewardRedemptionsData.Message)
	}
}

func Test_GetChannelRewardRedemptionsWithFilters_Success(t *testing.T) {
	// Arrange
	accessToken := "access-token"

	expectedJSON := `{
		"data": [{
			"redemptions": [{
				"id": "redemption-id-1",
				"redeemed_at": "2025-12-19T06:26:36Z",
				"redeemer": {
					"broadcaster_user_id": 42
				},
				"status": "pending",
				"user_input": "Some message"
			}],
			"reward": {
				"can_manage": true,
				"cost": 101,
				"description": "description-1",
				"id": "reward-id-1",
				"is_deleted": false,
				"title": "Title-1"
			}
		}],
		"message": "test-message",
		"pagination": {
			"next_cursor": "some-cursor"
		}
	}`

	filters := kickfilters.NewRewardRedemptionsFilter().WithCursor("test").WithRewardID("123").WithRewardIDs([]string{"111", "222"}).WithStatus(kickchannelrewardstatus.Pending)

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/channels/rewards/redemptions?cursor=test&id=111&id=222&reward_id=123&status=pending" {
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

	channelRewardRedemptionsData, err := client.ChannelReward().GetChannelRewardRedemptions(t.Context(), accessToken, filters)

	// Assert
	if channelRewardRedemptionsData == nil {
		t.Fatal("Expected channelRewardsData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if len(channelRewardRedemptionsData.Data) != 1 {
		t.Fatalf("Expected Data to be 1 slice, got %d", len(channelRewardRedemptionsData.Data))
	}

	if len(channelRewardRedemptionsData.Data[0].Redemptions) != 1 {
		t.Fatalf("Expected Data to be 1 slice, got %d", len(channelRewardRedemptionsData.Data[0].Redemptions))
	}

	if channelRewardRedemptionsData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", channelRewardRedemptionsData.Message)
	}
}
