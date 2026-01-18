package kick_test

import (
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/tests/mocks"
)

func Test_GetUserByID_Success(t *testing.T) {
	// Arrange
	accessToken := "access-token"
	var user int64 = 2

	expectedJSON := `{
		"data": [{
			"email": "test-user-2@email.com",
			"name": "test-user-2",
			"profile_picture": "https://test-user-2",
			"user_id": 2
		}],
		"message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/users?id=2" {
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

	userData, err := client.User().GetUserByID(t.Context(), accessToken, user)

	// Assert
	if userData == nil {
		t.Fatal("Expected userData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if len(userData.Data) != 1 {
		t.Fatalf("Expected Data to be 1 slice, got %d", len(userData.Data))
	}

	if userData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", userData.Message)
	}
}
