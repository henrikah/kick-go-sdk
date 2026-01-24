package kick_test

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/henrikah/kick-go-sdk/v2"
	"github.com/henrikah/kick-go-sdk/v2/enums/kickscopes"
	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
	"github.com/henrikah/kick-go-sdk/v2/kickoauthtypes"
	"github.com/henrikah/kick-go-sdk/v2/tests/mocks"
)

func Test_TokenIntrospectMissingAccessToken_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := ""

	config := kickoauthtypes.OAuthClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}
	client, _ := kick.NewOAuthClient(config)

	// Act
	tokenIntrospectData, err := client.TokenIntrospect(t.Context(), accessToken)

	// Assert
	if tokenIntrospectData != nil {
		t.Fatal("Expected tokenIntrospectData to be nil")
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

func Test_TokenIntrospectUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"message": "Invalid request"}`

	accessToken := "access-token"

	mockClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return mocks.NewMockResponse(http.StatusUnauthorized, errorJSON), nil
		},
	}

	config := kickoauthtypes.OAuthClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   mockClient,
	}
	client, _ := kick.NewOAuthClient(config)

	// Act
	tokenIntrospectData, err := client.TokenIntrospect(t.Context(), accessToken)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if tokenIntrospectData != nil {
		t.Fatal("Expected tokenIntrospectData to be nil on error")
	}

	apiError := kickerrors.IsAPIError(err)

	if apiError == nil {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_TokenIntrospect_Success(t *testing.T) {
	// Arrange
	clientID := "test-id"
	tokenType := "user"

	expectedScopes := kickscopes.Scopes{kickscopes.UserRead, kickscopes.ChannelRead}
	expectedJSON := fmt.Sprintf(`{
		"data": {
			"active": true,
			"client_id": "%s",
			"exp": 1,
			"scope": "user:read channel:read",
			"token_type": "%s"
		},
		"message": "test-message"
	}`, clientID, tokenType)

	clientSecret := "test-secret"

	accessToken := "access-token"

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://id.kick.com/oauth/token/introspect" {
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

	config := kickoauthtypes.OAuthClientConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HTTPClient:   httpClient,
	}

	client, _ := kick.NewOAuthClient(config)

	// Act

	tokenIntrospectData, err := client.TokenIntrospect(t.Context(), accessToken)

	// Assert
	if tokenIntrospectData == nil {
		t.Fatal("Expected tokenIntrospectData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if !tokenIntrospectData.Data.Active {
		t.Fatalf("Expected Active to be true, got %t", tokenIntrospectData.Data.Active)
	}

	if tokenIntrospectData.Data.ClientID != clientID {
		t.Fatalf("Expected ClientID to be %s, got %s", clientID, tokenIntrospectData.Data.ClientID)
	}

	if !reflect.DeepEqual(tokenIntrospectData.Data.Scopes, expectedScopes) {
		t.Fatalf("Expected Scope to be %s, got %s", expectedScopes, tokenIntrospectData.Data.Scopes)
	}

	if tokenIntrospectData.Data.TokenType != tokenType {
		t.Fatalf("Expected TokenType to be %s, got %s", tokenType, tokenIntrospectData.Data.TokenType)
	}
	if tokenIntrospectData.Data.Expires != 1 {
		t.Fatalf("Expected Expires to be %d, got %d", 1, tokenIntrospectData.Data.Expires)
	}
}
