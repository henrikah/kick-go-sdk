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

func Test_CreateChannelRewardMissingAccessToken_Error(t *testing.T) {
	// Arrange
	ctx := t.Context()
	httpClient := http.DefaultClient

	accessToken := ""

	channelRewardData := kickapitypes.CreateChannelReward{
		Cost:  100,
		Title: "Test",
	}

	config := kickapitypes.APIClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act

	createChannelRewardResult, err := client.ChannelReward().CreateChannelReward(ctx, accessToken, channelRewardData)

	// Assert

	if createChannelRewardResult != nil {
		t.Fatal("Expected createChannelRewardResult to be nil")
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

func Test_CreateChannelRewardLessThanOneCost_Error(t *testing.T) {
	// Arrange
	ctx := t.Context()
	httpClient := http.DefaultClient

	accessToken := "access-token"

	channelRewardData := kickapitypes.CreateChannelReward{
		Cost:  0,
		Title: "Test",
	}

	config := kickapitypes.APIClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act

	createChannelRewardResult, err := client.ChannelReward().CreateChannelReward(ctx, accessToken, channelRewardData)

	// Assert

	if createChannelRewardResult != nil {
		t.Fatal("Expected createChannelRewardResult to be nil")
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

func Test_CreateChannelRewardUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"message": "Invalid request"}`

	ctx := t.Context()

	accessToken := "access-token"

	channelRewardData := kickapitypes.CreateChannelReward{
		Cost:  100,
		Title: "Test",
	}

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

	createChannelRewardResult, err := client.ChannelReward().CreateChannelReward(ctx, accessToken, channelRewardData)

	// Assert

	if createChannelRewardResult != nil {
		t.Fatal("Expected createChannelRewardResult to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &apiError) {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_CreateChannelRewardWithMissingFields_Success(t *testing.T) {
	// Arrange
	ctx := t.Context()
	clientID := "test-id"
	clientSecret := "test-secret"

	accessToken := "access-token"

	channelRewardData := kickapitypes.CreateChannelReward{
		Cost:  100,
		Title: "Test",
	}

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/channels/rewards" {
				t.Fatalf("Unexpected request URL: %s", req.URL.String())
			}

			if req.Method != "POST" {
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

			if _, ok := updateData["background_color"]; ok {
				t.Fatal("Expected background_color to be omitted")
			}

			if _, ok := updateData["is_enabled"]; ok {
				t.Fatal("Expected is_enabled to be omitted")
			}

			return mocks.NewMockResponse(http.StatusOK, `{ "data": { "cost": 100, "title": "Test" }, "message": "text" }`), nil
		},
	}

	config := kickapitypes.APIClientConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HTTPClient:   httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act

	createChannelRewardResult, err := client.ChannelReward().CreateChannelReward(ctx, accessToken, channelRewardData)

	// Assert

	if createChannelRewardResult == nil {
		t.Fatal("Expected createChannelRewardResult to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}
}

func Test_CreateChannelReward_Success(t *testing.T) {
	// Arrange
	ctx := t.Context()
	clientID := "test-id"
	clientSecret := "test-secret"

	accessToken := "access-token"

	backgroundColor := "#123456"
	description := "description"
	isEnabled := true
	isUserInputRequired := false
	shouldRedemptionsSkipRequestQueue := false

	channelRewardData := kickapitypes.CreateChannelReward{
		BackgroundColor:                   &backgroundColor,
		Cost:                              100,
		Description:                       &description,
		IsEnabled:                         &isEnabled,
		IsUserInputRequired:               &isUserInputRequired,
		ShouldRedemptionsSkipRequestQueue: &shouldRedemptionsSkipRequestQueue,
		Title:                             "Test",
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
			"title": "Test"
		},
		"message": "test"
	}`
	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/channels/rewards" {
				t.Fatalf("Unexpected request URL: %s", req.URL.String())
			}

			if req.Method != "POST" {
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

			var updateData kickapitypes.UpdateChannelRequest

			err := bodyDecoder.Decode(&updateData)

			if err != nil {
				return mocks.NewMockResponse(http.StatusInternalServerError, ""), nil
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

	createChannelRewardResult, err := client.ChannelReward().CreateChannelReward(ctx, accessToken, channelRewardData)

	// Assert

	if createChannelRewardResult == nil {
		t.Fatal("Expected createChannelRewardResult to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}
}
