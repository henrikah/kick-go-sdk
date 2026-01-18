package kick_test

import (
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/tests/mocks"
)

func TestGetKicksLeaderboardSuccess(t *testing.T) {
	// Arrange
	accessToken := "access-token"

	expectedJSON := `{
		"data": {
			"lifetime": [{
				"gifted_amount": 100,
				"rank": 1,
				"user_id": 321,
				"username": "TestUser"
			}],
			"month": [{
				"gifted_amount": 100,
				"rank": 1,
				"user_id": 321,
				"username": "TestUser"
			}],
			"week": [{
				"gifted_amount": 100,
				"rank": 1,
				"user_id": 321,
				"username": "TestUser"
			}]
		},
		"message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/kicks/leaderboard" {
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
		HTTPClient:   httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act

	kicksLeaderboardData, err := client.Kicks().GetKicksLeaderboard(t.Context(), accessToken, nil)

	// Assert
	if kicksLeaderboardData == nil {
		t.Fatal("Expected kicksLeaderboardData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if len(kicksLeaderboardData.Data.Lifetime) != 1 {
		t.Fatalf("Expected Data.Lifetime to be 1 slice, got %d", len(kicksLeaderboardData.Data.Lifetime))
	}

	if len(kicksLeaderboardData.Data.Month) != 1 {
		t.Fatalf("Expected Data.Month to be 1 slice, got %d", len(kicksLeaderboardData.Data.Month))
	}

	if len(kicksLeaderboardData.Data.Week) != 1 {
		t.Fatalf("Expected Data.Week to be 1 slice, got %d", len(kicksLeaderboardData.Data.Week))
	}

	if kicksLeaderboardData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", kicksLeaderboardData.Message)
	}
}
