package kick_test

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/henrikah/kick-go-sdk/v2"
	"github.com/henrikah/kick-go-sdk/v2/enums/kickscopes"
	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
	"github.com/henrikah/kick-go-sdk/v2/kickoauthtypes"
	"github.com/henrikah/kick-go-sdk/v2/tests/mocks"
)

func Test_ExchangeAuthorizationCodeMissingRedirectURI_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	config := kickoauthtypes.OAuthClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}

	redirectURI := ""
	authorizationCode := "authorization-code"
	codeVerifier := "code-verifier"

	client, _ := kick.NewOAuthClient(config)

	// Act

	tokenData, err := client.ExchangeAuthorizationCode(t.Context(), redirectURI, authorizationCode, codeVerifier)
	// Assert
	if tokenData != nil {
		t.Fatal("Expected tokenData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationErr := kickerrors.IsValidationError(err)

	if validationErr == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationErr.Field != "redirectURI" {
		t.Fatalf("Expected error on field 'redirectURI', got '%s'", validationErr.Field)
	}
}
func Test_ExchangeAuthorizationCodeMissingAuthorizationCode_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	config := kickoauthtypes.OAuthClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}

	redirectURI := "http://localhost"
	authorizationCode := ""
	codeVerifier := "code-verifier"

	client, _ := kick.NewOAuthClient(config)

	// Act

	tokenData, err := client.ExchangeAuthorizationCode(t.Context(), redirectURI, authorizationCode, codeVerifier)
	// Assert
	if tokenData != nil {
		t.Fatal("Expected tokenData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationErr := kickerrors.IsValidationError(err)

	if validationErr == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationErr.Field != "authorizationCode" {
		t.Fatalf("Expected error on field 'authorizationCode', got '%s'", validationErr.Field)
	}
}
func Test_ExchangeAuthorizationCodeMissingCodeVerifier_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	config := kickoauthtypes.OAuthClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}

	redirectURI := "http://localhost"
	authorizationCode := "authorization-code"
	codeVerifier := ""

	client, _ := kick.NewOAuthClient(config)

	// Act

	tokenData, err := client.ExchangeAuthorizationCode(t.Context(), redirectURI, authorizationCode, codeVerifier)
	// Assert
	if tokenData != nil {
		t.Fatal("Expected tokenData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationErr := kickerrors.IsValidationError(err)

	if validationErr == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationErr.Field != "codeVerifier" {
		t.Fatalf("Expected error on field 'codeVerifier', got '%s'", validationErr.Field)
	}
}

func Test_ExchangeAuthorizationCodeWrongCredentials_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"error": "Invalid request"}`

	redirectURI := "http://localhost"
	authorizationCode := "authorization-code"
	codeVerifier := "code-verifier"

	mockClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return mocks.NewMockResponse(http.StatusBadRequest, errorJSON), nil
		},
	}

	config := kickoauthtypes.OAuthClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   mockClient,
	}
	client, _ := kick.NewOAuthClient(config)

	// Act
	tokenData, err := client.ExchangeAuthorizationCode(t.Context(), redirectURI, authorizationCode, codeVerifier)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if tokenData != nil {
		t.Fatal("Expected tokenData to be nil on error")
	}

	apiErr := kickerrors.IsAPIError(err)

	if apiErr == nil {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_ExchangeAuthorizationCode_Success(t *testing.T) {
	// Arrange
	clientID := "test-id"
	clientSecret := "test-secret"
	redirectURI := "http://localhost"
	authorizationCode := "authorization-code"
	codeVerifier := "code-verifier"

	accessToken := "access-token"
	tokenType := "bearer"
	refreshToken := "refresh-token"
	expiresIn := int64(3600)
	scope := kickscopes.Scopes{kickscopes.UserRead, kickscopes.ChannelRead}
	scopeJSON := "user:read channel:read"

	expectedJSON := fmt.Sprintf(
		`{"access_token":"%s", "token_type":"%s", "refresh_token":"%s", "expires_in":%d, "scope":"%s"}`,
		accessToken,
		tokenType,
		refreshToken,
		expiresIn,
		scopeJSON,
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
				"grant_type":    "authorization_code",
				"code":          authorizationCode,
				"redirect_uri":  redirectURI,
				"client_id":     clientID,
				"client_secret": clientSecret,
				"code_verifier": codeVerifier,
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

	config := kickoauthtypes.OAuthClientConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HTTPClient:   httpClient,
	}

	client, _ := kick.NewOAuthClient(config)

	// Act

	tokenData, err := client.ExchangeAuthorizationCode(t.Context(), redirectURI, authorizationCode, codeVerifier)

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

	if tokenData.RefreshToken != refreshToken {
		t.Fatalf("Expected RefreshToken to be %s, got %s", refreshToken, tokenData.RefreshToken)
	}

	if tokenData.ExpiresIn != expiresIn {
		t.Fatalf("Expected ExpiresIn to be %d, got %d", expiresIn, tokenData.ExpiresIn)
	}

	if !reflect.DeepEqual(tokenData.Scope, scope) {
		t.Fatalf("Expected Scope to be %s, got %s", scope, tokenData.Scope)
	}
}
