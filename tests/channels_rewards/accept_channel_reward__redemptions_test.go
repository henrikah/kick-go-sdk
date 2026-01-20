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

func Test_AcceptChannelRewardRedemptionsMissingAccessToken_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := ""

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	acceptChannelRewardRedemptionsData, err := client.ChannelReward().AcceptRewardRedemption(t.Context(), accessToken, nil)

	// Assert
	if acceptChannelRewardRedemptionsData != nil {
		t.Fatal("Expected acceptChannelRewardRedemptionsData to be nil")
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

func Test_AcceptChannelRewardRedemptionsUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"message": "Invalid request"}`

	accessToken := "access-token"
	ids := []string{"id-1"}

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
	acceptChannelRewardRedemptionsData, err := client.ChannelReward().AcceptRewardRedemption(t.Context(), accessToken, ids)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if acceptChannelRewardRedemptionsData != nil {
		t.Fatal("Expected acceptChannelRewardRedemptionsData to be nil on error")
	}

	if !errors.As(err, &apiError) {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_AcceptChannelRewardRedemptions_Success(t *testing.T) {
	// Arrange
	accessToken := "access-token"
	ids := []string{"id-1"}
	expectedJSON := `{
		"data": [],
		"message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/channels/rewards/redemptions/accept" {
				t.Fatalf("Unexpected request URL: %s", req.URL.String())
			}

			if req.Method != "POST" {
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

	acceptChannelRewardRedemptionsData, err := client.ChannelReward().AcceptRewardRedemption(t.Context(), accessToken, ids)

	// Assert
	if acceptChannelRewardRedemptionsData == nil {
		t.Fatal("Expected acceptChannelRewardRedemptionsData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if len(acceptChannelRewardRedemptionsData.Data) != 0 {
		t.Fatalf("Expected Data to be 0 slice, got %d", len(acceptChannelRewardRedemptionsData.Data))
	}
}

func Test_AcceptChannelRewardRedemptionsWithFailedRedemptions_Success(t *testing.T) {
	// Arrange
	accessToken := "access-token"
	ids := []string{"id-1"}
	expectedJSON := `{
		"data": [{
			"id": "id-1",
			"reason": "NOT_OWNED"
		}],
		"message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/channels/rewards/redemptions/accept" {
				t.Fatalf("Unexpected request URL: %s", req.URL.String())
			}

			if req.Method != "POST" {
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

	acceptChannelRewardRedemptionsData, err := client.ChannelReward().AcceptRewardRedemption(t.Context(), accessToken, ids)

	// Assert
	if acceptChannelRewardRedemptionsData == nil {
		t.Fatal("Expected acceptChannelRewardRedemptionsData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if len(acceptChannelRewardRedemptionsData.Data) != 1 {
		t.Fatalf("Expected Data to be 1 slice, got %d", len(acceptChannelRewardRedemptionsData.Data))
	}
}
