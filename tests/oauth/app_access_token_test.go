package kick_test

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/henrikah/kick-go-sdk"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickerrors"
	"github.com/henrikah/kick-go-sdk/tests/mocks"
)

func Test_GetAppAccessTokenWrongCredentials_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"error": "Invalid request"}`

	ctx := t.Context()

	mockClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return mocks.NewMockResponse(http.StatusBadRequest, errorJSON), nil
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
	tokenData, err := client.OAuth.GetAppAccessToken(ctx)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if tokenData != nil {
		t.Fatal("Expected tokenData to be nil on error")
	}

	if !errors.As(err, &apiError) {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_GetAppAccessToken_Success(t *testing.T) {
	// Arrange
	ctx := t.Context()
	clientID := "test-id"
	clientSecret := "test-secret"

	accessToken := "access-token"
	tokenType := "bearer"
	expiresIn := int64(3600)

	expectedJSON := fmt.Sprintf(
		`{"access_token":"%s", "token_type":"%s", "expires_in":%d}`,
		accessToken,
		tokenType,
		expiresIn,
	)

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://id.kick.com/oauth/token" {
				t.Fatalf("Unexpected request URL: %s", req.URL.String())
			}

			if req.Method != "POST" {
				t.Fatalf("Unexpected request method: %s", req.Method)
			}

			if req.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
				t.Fatal("Missing Content-Type header")
			}

			t.Cleanup(func() {
				defer func() {
					if err := req.Body.Close(); err != nil {
						t.Logf("failed to close request body: %v", err)
					}
				}()
			})

			bodyBytes, err := io.ReadAll(req.Body)
			if err != nil {
				t.Fatalf("Failed to read request body: %v", err)
			}
			bodyString := string(bodyBytes)

			parsedValues, err := url.ParseQuery(bodyString)
			if err != nil {
				t.Fatalf("Failed to parse body as form data: %v. Body was: %s", err, bodyString)
			}

			expectedValues := map[string]string{
				"grant_type":    "client_credentials",
				"client_id":     clientID,
				"client_secret": clientSecret,
			}

			for key, expectedValue := range expectedValues {
				actualValue := parsedValues.Get(key)
				if actualValue != expectedValue {
					t.Fatalf("Mismatch for form field '%s'. Expected: '%s', Got: '%s'", key, expectedValue, actualValue)
				}
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

	tokenData, err := client.OAuth.GetAppAccessToken(ctx)

	// Assert
	if tokenData == nil {
		t.Fatal("Expected tokenData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if tokenData.AccessToken != accessToken {
		t.Fatalf("Expected AccessToken to be %s, got %s", accessToken, tokenData.AccessToken)
	}

	if tokenData.TokenType != tokenType {
		t.Fatalf("Expected TokenType to be %s, got %s", tokenType, tokenData.TokenType)
	}

	if tokenData.ExpiresIn != expiresIn {
		t.Fatalf("Expected ExpiresIn to be %d, got %d", expiresIn, tokenData.ExpiresIn)
	}
}
