package kick_test

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/henrikah/kick-go-sdk"
	"github.com/henrikah/kick-go-sdk/enums/kickscopes"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickerrors"
	"github.com/henrikah/kick-go-sdk/tests/mocks"
)

func Test_TokenIntrospectMissingAccessToken_Error(t *testing.T) {
	// Arrange
	ctx := t.Context()
	httpClient := http.DefaultClient

	accessToken := ""

	config := kickapitypes.APIClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	tokenIntrospectData, err := client.User().TokenIntrospect(ctx, accessToken)

	// Assert
	if tokenIntrospectData != nil {
		t.Fatal("Expected tokenIntrospectData to be nil")
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

func Test_TokenIntrospectUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"message": "Invalid request"}`

	ctx := t.Context()

	accessToken := "access-token"

	mockClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return mocks.NewMockResponse(http.StatusUnauthorized, errorJSON), nil
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
	tokenIntrospectData, err := client.User().TokenIntrospect(ctx, accessToken)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if tokenIntrospectData != nil {
		t.Fatal("Expected tokenIntrospectData to be nil on error")
	}

	if !errors.As(err, &apiError) {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_TokenIntrospect_Success(t *testing.T) {
	// Arrange
	ctx := t.Context()
	clientID := "test-id"
	clientSecret := "test-secret"

	accessToken := "access-token"

	expectedScopes := kickscopes.Scopes{kickscopes.UserRead, kickscopes.ChannelRead}
	expectedJSON := `{
		"data": {
			"active": true,
			"client_id": "client-id",
			"exp": 1,
			"scope": "user:read channel:read",
			"token_type": "access_token"
		},
		"message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/token/introspect" {
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
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HTTPClient:   httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act

	tokenIntrospectData, err := client.User().TokenIntrospect(ctx, accessToken)

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

	if !reflect.DeepEqual(tokenIntrospectData.Data.Scopes, expectedScopes) {
		t.Fatalf("Expected Scope to be %s, got %s", expectedScopes, tokenIntrospectData.Data.Scopes)
	}
}
