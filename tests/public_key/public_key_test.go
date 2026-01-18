package kick_test

import (
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/tests/mocks"
)

func Test_GetPublicKey_Success(t *testing.T) {
	// Arrange
	expectedJSON := `{
		"data": {
			"public_key": "public-key"
		},
		"message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/public-key" {
				t.Fatalf("Unexpected request URL: %s", req.URL.String())
			}

			if req.Method != "GET" {
				t.Fatalf("Unexpected request method: %s", req.Method)
			}

			if req.Header.Get("Accept") != "application/json" {
				t.Fatal("Missing Accept header")
			}

			return mocks.NewMockResponse(http.StatusOK, expectedJSON), nil
		},
	}

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act

	publicKeyData, err := client.PublicKey().GetWebhookPublicKey(t.Context())

	// Assert
	if publicKeyData == nil {
		t.Fatal("Expected publicKeyData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if publicKeyData.Data.PublicKey != "public-key" {
		t.Fatalf("Expected PublicKey to be 'public-key', got %s", publicKeyData.Data.PublicKey)
	}

	if publicKeyData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", publicKeyData.Message)
	}
}
