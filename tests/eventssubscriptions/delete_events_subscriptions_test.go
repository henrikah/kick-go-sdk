package kick_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickerrors"
	"github.com/henrikah/kick-go-sdk/tests/mocks"
)

func Test_DeleteEventsSubscriptionsMissingAccessToken_Error(t *testing.T) {
	// Arrange
	ctx := t.Context()
	httpClient := http.DefaultClient

	accessToken := ""
	subscriptionIDs := []string{"id-1", "id-2"}

	config := kickapitypes.APIClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	err := client.EventsSubscription.DeleteEventSubscriptions(ctx, accessToken, subscriptionIDs)

	// Assert
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

func Test_DeleteEventsSubscriptionsMissingSubscriptionIDs_Error(t *testing.T) {
	// Arrange
	ctx := t.Context()
	httpClient := http.DefaultClient

	accessToken := "access-token"
	subscriptionIDs := []string{}

	config := kickapitypes.APIClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	err := client.EventsSubscription.DeleteEventSubscriptions(ctx, accessToken, subscriptionIDs)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &validationError) {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "subscriptionIDs" {
		t.Fatalf("Expected error on field 'subscriptionIDs', got '%s'", validationError.Field)
	}
}

func Test_DeleteEventsSubscriptionsUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"message": "Invalid request"}`

	ctx := t.Context()

	accessToken := "access-token"
	subscriptionIDs := []string{"id-1", "id-2"}

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
	err := client.EventsSubscription.DeleteEventSubscriptions(ctx, accessToken, subscriptionIDs)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &apiError) {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_DeleteEventsSubscriptions_Success(t *testing.T) {
	// Arrange
	ctx := t.Context()
	clientID := "test-id"
	clientSecret := "test-secret"

	accessToken := "access-token"
	subscriptionIDs := []string{"id-1", "id-2"}

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/events/subscriptions?id=id-1&id=id-2" {
				t.Fatalf("Unexpected request URL: %s", req.URL.String())
			}

			if req.Method != "DELETE" {
				t.Fatalf("Unexpected request method: %s", req.Method)
			}

			if req.Header.Get("Accept") != "application/json" {
				t.Fatal("Missing Accept header")
			}

			if req.Header.Get("Authorization") != "Bearer "+accessToken {
				t.Fatal("Missing Authorization header")
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

	err := client.EventsSubscription.DeleteEventSubscriptions(ctx, accessToken, subscriptionIDs)

	// Assert
	if err != nil {
		t.Fatal("Expected error to be nil")
	}
}
