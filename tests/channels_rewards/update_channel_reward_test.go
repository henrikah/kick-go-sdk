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

func Test_UpdateChannelRewardMissingAccessToken_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := ""

	channelRewardID := "test-reward-id"
	channelRewardData := kickapitypes.UpdateChannelReward{}

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	updateChannelRewardData, err := client.ChannelReward().UpdateChannelReward(t.Context(), accessToken, channelRewardID, channelRewardData)

	// Assert
	if updateChannelRewardData != nil {
		t.Fatal("Expected updateChannelRewardData to be nil")
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

func Test_UpdateChannelRewardCostLessThanOne_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := "access-token"

	channelRewardID := "test-reward-id"
	cost := 0
	channelRewardData := kickapitypes.UpdateChannelReward{
		Cost: &cost,
	}

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	updateChannelRewardData, err := client.ChannelReward().UpdateChannelReward(t.Context(), accessToken, channelRewardID, channelRewardData)

	// Assert
	if updateChannelRewardData != nil {
		t.Fatal("Expected updateChannelRewardData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &validationError) {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "Cost" {
		t.Fatalf("Expected error on field 'Cost', got '%s'", validationError.Field)
	}
}

func Test_UpdateChannelRewardDescriptionTooLong_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := "access-token"

	channelRewardID := "test-reward-id"
	description := "123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901"
	channelRewardData := kickapitypes.UpdateChannelReward{
		Description: &description,
	}

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	updateChannelRewardData, err := client.ChannelReward().UpdateChannelReward(t.Context(), accessToken, channelRewardID, channelRewardData)

	// Assert
	if updateChannelRewardData != nil {
		t.Fatal("Expected updateChannelRewardData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &validationError) {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "Description" {
		t.Fatalf("Expected error on field 'Description', got '%s'", validationError.Field)
	}
}

func Test_UpdateChannelRewardTitleTooLong_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := "access-token"

	channelRewardID := "test-reward-id"
	title := "012345678901234567890123456789012345678901234567891"
	channelRewardData := kickapitypes.UpdateChannelReward{
		Title: &title,
	}

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	updateChannelRewardData, err := client.ChannelReward().UpdateChannelReward(t.Context(), accessToken, channelRewardID, channelRewardData)

	// Assert
	if updateChannelRewardData != nil {
		t.Fatal("Expected updateChannelRewardData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &validationError) {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "Title" {
		t.Fatalf("Expected error on field 'Title', got '%s'", validationError.Field)
	}
}

func Test_UpdateChannelRewardUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"message": "Invalid request"}`

	accessToken := "access-token"

	channelRewardID := "test-reward-id"
	channelRewardData := kickapitypes.UpdateChannelReward{}

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
	updateChannelRewardData, err := client.ChannelReward().UpdateChannelReward(t.Context(), accessToken, channelRewardID, channelRewardData)

	// Assert
	if updateChannelRewardData != nil {
		t.Fatal("Expected updateChannelRewardData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &apiError) {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_UpdateChannelRewardWithMissingFields_Success(t *testing.T) {
	// Arrange
	accessToken := "access-token"

	channelRewardID := "test-reward-id"
	channelRewardData := kickapitypes.UpdateChannelReward{}

	expectedJSON := `{
		"data": {
			"background_color": "#123456",
			"cost": 100,
			"description": "description",
			"id": "generated-id",
			"is_enabled": true,
			"is_paused": false,
			"is_user_input_required": false,
			"should_redemptions_skip_request_queue": false,
			"title": "test-title"
		},
		"message": "test"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/channels/rewards/test-reward-id" {
				t.Fatalf("Unexpected request URL: %s", req.URL.String())
			}

			if req.Method != "PATCH" {
				t.Fatalf("Unexpected request method: %s", req.Method)
			}

			if req.Header.Get("Accept") != "application/json" {
				t.Fatal("Missing Accept header")
			}

			if req.Header.Get("Content-Type") != "application/json" {
				t.Fatal("Missing Content-Type header")
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

			if _, ok := updateData["cost"]; ok {
				t.Fatal("Expected cost to be omitted")
			}

			if _, ok := updateData["title"]; ok {
				t.Fatal("Expected title to be omitted")
			}

			return mocks.NewMockResponse(http.StatusOK, expectedJSON), nil
		},
	}

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act

	updateChannelRewardData, err := client.ChannelReward().UpdateChannelReward(t.Context(), accessToken, channelRewardID, channelRewardData)

	// Assert
	if updateChannelRewardData == nil {
		t.Fatal("Expected updateChannelRewardData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}
}

func Test_UpdateChannel_Success(t *testing.T) {
	// Arrange
	accessToken := "access-token"

	channelRewardID := "test-reward-id"
	cost := 100
	title := "test-title"
	channelRewardData := kickapitypes.UpdateChannelReward{
		Cost:  &cost,
		Title: &title,
	}

	expectedJSON := `{
		"data": {
			"background_color": "#123456",
			"cost": 100,
			"description": "description",
			"id": "generated-id",
			"is_enabled": true,
			"is_paused": false,
			"is_user_input_required": false,
			"should_redemptions_skip_request_queue": false,
			"title": "test-title"
		},
		"message": "test"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/channels/rewards/test-reward-id" {
				t.Fatalf("Unexpected request URL: %s", req.URL.String())
			}

			if req.Method != "PATCH" {
				t.Fatalf("Unexpected request method: %s", req.Method)
			}

			if req.Header.Get("Accept") != "application/json" {
				t.Fatal("Missing Accept header")
			}

			if req.Header.Get("Content-Type") != "application/json" {
				t.Fatal("Missing Content-Type header")
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

			var updateData kickapitypes.UpdateChannelReward

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

	updateChannelRewardData, err := client.ChannelReward().UpdateChannelReward(t.Context(), accessToken, channelRewardID, channelRewardData)

	// Assert

	if updateChannelRewardData == nil {
		t.Fatal("Expected updateChannelRewardData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}
}
