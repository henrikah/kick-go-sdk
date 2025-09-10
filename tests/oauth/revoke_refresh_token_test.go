package kick_test

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/henrikah/kick-go-sdk"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickerrors"
	"github.com/henrikah/kick-go-sdk/tests/mocks"
)

func Test_RevokeRefreshTokenMissingToken_Error(t *testing.T) {
	// Arrange
	ctx := t.Context()
	httpClient := http.DefaultClient
	refreshToken := ""

	config := kickapitypes.APIClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError
	// Act
	err := client.OAuth().RevokeRefreshToken(ctx, refreshToken)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &validationError) {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "refreshToken" {
		t.Fatalf("Expected error on field 'refreshToken', got '%s'", validationError.Field)
	}
}

func Test_RevokeRefreshToken_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"error": "Invalid request"}`

	ctx := t.Context()
	refreshToken := "refresh-token"

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
	err := client.OAuth().RevokeRefreshToken(ctx, refreshToken)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &apiError) {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_RevokeRefreshToken_Success(t *testing.T) {
	// Arrange
	ctx := t.Context()
	clientID := "test-id"
	clientSecret := "test-secret"

	refreshToken := "refresh-token"

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://id.kick.com/oauth/revoke" {
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
				"token":           refreshToken,
				"token_hint_type": "refresh_token",
			}

			for key, expectedValue := range expectedValues {
				actualValue := parsedValues.Get(key)
				if actualValue != expectedValue {
					t.Fatalf("Mismatch for form field '%s'. Expected: '%s', Got: '%s'", key, expectedValue, actualValue)
				}
			}
			return mocks.NewMockResponse(http.StatusOK, ""), nil
		},
	}

	config := kickapitypes.APIClientConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HTTPClient:   httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act

	err := client.OAuth().RevokeRefreshToken(ctx, refreshToken)

	// Assert
	if err != nil {
		t.Fatal("Expected error to be nil")
	}
}
