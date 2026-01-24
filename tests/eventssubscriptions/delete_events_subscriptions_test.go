package kick_test

import (
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk/v2"
	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
	"github.com/henrikah/kick-go-sdk/v2/tests/mocks"
)

func Test_DeleteEventsSubscriptionsMissingAccessToken_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := ""
	subscriptionIDs := []string{"id-1", "id-2"}

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	// Act
	err := client.EventsSubscription().DeleteEventSubscriptions(t.Context(), accessToken, subscriptionIDs)

	// Assert
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

func Test_DeleteEventsSubscriptionsMissingSubscriptionIDs_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := "access-token"
	subscriptionIDs := []string{}

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	// Act
	err := client.EventsSubscription().DeleteEventSubscriptions(t.Context(), accessToken, subscriptionIDs)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationError := kickerrors.IsValidationError(err)

	if validationError == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "subscriptionIDs" {
		t.Fatalf("Expected error on field 'subscriptionIDs', got '%s'", validationError.Field)
	}
}

func Test_DeleteEventsSubscriptionsUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"message": "Invalid request"}`

	accessToken := "access-token"
	subscriptionIDs := []string{"id-1", "id-2"}

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
	err := client.EventsSubscription().DeleteEventSubscriptions(t.Context(), accessToken, subscriptionIDs)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	apiError := kickerrors.IsAPIError(err)

	if apiError == nil {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_DeleteEventsSubscriptions_Success(t *testing.T) {
	// Arrange
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
		HTTPClient: httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act

	err := client.EventsSubscription().DeleteEventSubscriptions(t.Context(), accessToken, subscriptionIDs)

	// Assert
	if err != nil {
		t.Fatal("Expected error to be nil")
	}
}
