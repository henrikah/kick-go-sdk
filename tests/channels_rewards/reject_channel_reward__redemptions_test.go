package kick_test

import (
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk/v2"
	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
	"github.com/henrikah/kick-go-sdk/v2/tests/mocks"
)

func Test_RejectChannelRewardRedemptionsMissingAccessToken_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := ""

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	// Act
	rejectChannelRewardRedemptionsData, err := client.ChannelReward().RejectRewardRedemption(t.Context(), accessToken, nil)

	// Assert
	if rejectChannelRewardRedemptionsData != nil {
		t.Fatal("Expected rejectChannelRewardRedemptionsData to be nil")
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

func Test_RejectChannelRewardRedemptionsUnAuthorized_Error(t *testing.T) {
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

	// Act
	rejectChannelRewardRedemptionsData, err := client.ChannelReward().RejectRewardRedemption(t.Context(), accessToken, ids)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if rejectChannelRewardRedemptionsData != nil {
		t.Fatal("Expected rejectChannelRewardRedemptionsData to be nil on error")
	}

	apiError := kickerrors.IsAPIError(err)

	if apiError == nil {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_RejectChannelRewardRedemptions_Success(t *testing.T) {
	// Arrange
	accessToken := "access-token"
	ids := []string{"id-1"}
	expectedJSON := `{
		"data": [],
		"message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/channels/rewards/redemptions/reject" {
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

	rejectChannelRewardRedemptionsData, err := client.ChannelReward().RejectRewardRedemption(t.Context(), accessToken, ids)

	// Assert
	if rejectChannelRewardRedemptionsData == nil {
		t.Fatal("Expected rejectChannelRewardRedemptionsData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if len(rejectChannelRewardRedemptionsData.Data) != 0 {
		t.Fatalf("Expected Data to be 0 slice, got %d", len(rejectChannelRewardRedemptionsData.Data))
	}
}

func Test_RejectChannelRewardRedemptionsWithFailedRedemptions_Success(t *testing.T) {
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
			if req.URL.String() != "https://api.kick.com/public/v1/channels/rewards/redemptions/reject" {
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

	rejectChannelRewardRedemptionsData, err := client.ChannelReward().RejectRewardRedemption(t.Context(), accessToken, ids)

	// Assert
	if rejectChannelRewardRedemptionsData == nil {
		t.Fatal("Expected rejectChannelRewardRedemptionsData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if len(rejectChannelRewardRedemptionsData.Data) != 1 {
		t.Fatalf("Expected Data to be 1 slice, got %d", len(rejectChannelRewardRedemptionsData.Data))
	}
}
