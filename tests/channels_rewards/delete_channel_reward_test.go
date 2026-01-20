package kick_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk/v2"
	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
	"github.com/henrikah/kick-go-sdk/v2/tests/mocks"
)

func Test_DeleteChannelRewardMissingAccessToken_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := ""
	rewardID := "test-reward-id"

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	err := client.ChannelReward().DeleteChannelReward(t.Context(), accessToken, rewardID)

	// Assert
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

func Test_DeleteChannelRewardEmptyRewardID_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := "access-token"
	rewardID := ""

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	err := client.ChannelReward().DeleteChannelReward(t.Context(), accessToken, rewardID)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &validationError) {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "rewardID" {
		t.Fatalf("Expected error on field 'rewardID', got '%s'", validationError.Field)
	}
}

func Test_DeleteChannelRewardUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := ``

	accessToken := "access-token"
	rewardID := "test-reward-id"

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
	err := client.ChannelReward().DeleteChannelReward(t.Context(), accessToken, rewardID)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &apiError) {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_DeleteChannelReward_Success(t *testing.T) {
	// Arrange
	accessToken := "access-token"
	rewardID := "test-reward-id"

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/channels/rewards/test-reward-id" {
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

			return mocks.NewMockResponse(http.StatusNoContent, ""), nil
		},
	}

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act

	err := client.ChannelReward().DeleteChannelReward(t.Context(), accessToken, rewardID)

	// Assert
	if err != nil {
		t.Fatal("Expected error to be nil")
	}
}
