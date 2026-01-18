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

func Test_GetEventsSubscriptionsMissingAccessToken_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := ""

	config := kickapitypes.APIClientConfig{
		HTTPClient:   httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	eventsSubscriptionsData, err := client.EventsSubscription().GetEventSubscriptions(t.Context(), accessToken)

	// Assert
	if eventsSubscriptionsData != nil {
		t.Fatal("Expected eventsSubscriptionsData to be nil")
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

func Test_GetEventsSubscriptionsUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"message": "Invalid request"}`

	accessToken := "access-token"

	mockClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return mocks.NewMockResponse(http.StatusUnauthorized, errorJSON), nil
		},
	}

	config := kickapitypes.APIClientConfig{
		HTTPClient:   mockClient,
	}
	client, _ := kick.NewAPIClient(config)

	var apiError *kickerrors.APIError
	// Act
	eventsSubscriptionsData, err := client.EventsSubscription().GetEventSubscriptions(t.Context(), accessToken)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if eventsSubscriptionsData != nil {
		t.Fatal("Expected eventsSubscriptionsData to be nil on error")
	}

	if !errors.As(err, &apiError) {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_GetEventsSubscriptions_Success(t *testing.T) {
	// Arrange
	accessToken := "access-token"

	expectedJSON := `{
		"data": [{
			"app_id": "app-id-1",
			"broadcaster_user_id": 1,
			"created_at": "created-at-1",
			"event": "event-1",
			"id": "id-1",
			"method": "method-1",
			"updated_at": "updated-at-1",
			"version": 1
		},
		{
			"app_id": "app-id-2",
			"broadcaster_user_id": 2,
			"created_at": "created-at-2",
			"event": "event-2",
			"id": "id-2",
			"method": "method-2",
			"updated_at": "updated-at-2",
			"version": 2
		}],
		"message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/events/subscriptions" {
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
	eventsSubscriptionsData, err := client.EventsSubscription().GetEventSubscriptions(t.Context(), accessToken)

	// Assert
	if eventsSubscriptionsData == nil {
		t.Fatal("Expected eventsSubscriptionsData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if len(eventsSubscriptionsData.Data) != 2 {
		t.Fatalf("Expected Data to be 2 slices, got %d", len(eventsSubscriptionsData.Data))
	}

	if eventsSubscriptionsData.Data[0].AppID != "app-id-1" {
		t.Fatalf("Expected AppID to be true, got %s", eventsSubscriptionsData.Data[0].AppID)
	}

	if eventsSubscriptionsData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", eventsSubscriptionsData.Message)
	}
}
