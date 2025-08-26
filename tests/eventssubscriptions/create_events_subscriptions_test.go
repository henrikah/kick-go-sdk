package kick_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk"
	"github.com/henrikah/kick-go-sdk/enums/kickwebhookenum"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickerrors"
	"github.com/henrikah/kick-go-sdk/tests/mocks"
)

func Test_CreateEventsSubscriptionsMissingAccessToken_Error(t *testing.T) {
	// Arrange
	ctx := t.Context()
	httpClient := http.DefaultClient

	accessToken := ""
	events := []kickwebhookenum.WebhookType{kickwebhookenum.ChatMessageSent, kickwebhookenum.ChannelFollowed}

	config := kickapitypes.APIClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	eventsSubscriptionsData, err := client.EventsSubscription.CreateEventSubscriptions(ctx, accessToken, events)

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

func Test_CreateEventsSubscriptionsMissingEvents_Error(t *testing.T) {
	// Arrange
	ctx := t.Context()
	httpClient := http.DefaultClient

	accessToken := "access-token"
	events := []kickwebhookenum.WebhookType{}

	config := kickapitypes.APIClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	eventsSubscriptionsData, err := client.EventsSubscription.CreateEventSubscriptions(ctx, accessToken, events)

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

	if validationError.Field != "events" {
		t.Fatalf("Expected error on field 'events', got '%s'", validationError.Field)
	}
}

func Test_CreateEventsSubscriptionsUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"message": "Invalid request"}`

	ctx := t.Context()

	accessToken := "access-token"
	events := []kickwebhookenum.WebhookType{kickwebhookenum.ChatMessageSent, kickwebhookenum.ChannelFollowed}

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
	eventsSubscriptionsData, err := client.EventsSubscription.CreateEventSubscriptions(ctx, accessToken, events)

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

func Test_CreateEventsSubscriptions_Success(t *testing.T) {
	// Arrange
	ctx := t.Context()
	clientID := "test-id"
	clientSecret := "test-secret"

	accessToken := "access-token"
	events := []kickwebhookenum.WebhookType{kickwebhookenum.ChatMessageSent, kickwebhookenum.ChannelFollowed}

	expectedJSON := `{
		"data": [{
			"error": "error-1",
			"name": "name-1",
			"subscription_id": "subscription-id-1",
			"version": 1
		}],
		"message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/events/subscriptions" {
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
			t.Cleanup(func() {
				defer func() {
					if err := req.Body.Close(); err != nil {
						t.Logf("failed to close request body: %v", err)
					}
				}()
			})

			bodyDecoder := json.NewDecoder(req.Body)

			var updateData map[string]any

			err := bodyDecoder.Decode(&updateData)

			if err != nil {
				return mocks.NewMockResponse(http.StatusInternalServerError, ""), nil
			}

			if _, ok := updateData["broadcaster_user_id"]; ok {
				t.Fatal("Expected broadcaster_user_id to be omitted")
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

	eventsSubscriptionsData, err := client.EventsSubscription.CreateEventSubscriptions(ctx, accessToken, events)

	// Assert
	if eventsSubscriptionsData == nil {
		t.Fatal("Expected eventsSubscriptionsData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if len(eventsSubscriptionsData.Data) != 1 {
		t.Fatalf("Expected 1 slice, got %d", len(eventsSubscriptionsData.Data))
	}

	if eventsSubscriptionsData.Data[0].Name != "name-1" {
		t.Fatalf("Expected Name to be name-1, got %s", eventsSubscriptionsData.Data[0].Name)
	}

	if eventsSubscriptionsData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", eventsSubscriptionsData.Message)
	}
}

func Test_CreateEventsSubscriptionsAsApp_Success(t *testing.T) {
	// Arrange
	ctx := t.Context()
	clientID := "test-id"
	clientSecret := "test-secret"

	accessToken := "access-token"
	broadcasterUserID := 1
	events := []kickwebhookenum.WebhookType{kickwebhookenum.ChatMessageSent, kickwebhookenum.ChannelFollowed}

	expectedJSON := `{
		"data": [{
			"error": "error-1",
			"name": "name-1",
			"subscription_id": "subscription-id-1",
			"version": 1
		}],
		"message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/events/subscriptions" {
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

	eventsSubscriptionsData, err := client.EventsSubscription.CreateEventSubscriptionsAsApp(ctx, accessToken, broadcasterUserID, events)

	// Assert
	if eventsSubscriptionsData == nil {
		t.Fatal("Expected eventsSubscriptionsData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if len(eventsSubscriptionsData.Data) != 1 {
		t.Fatalf("Expected 1 slice, got %d", len(eventsSubscriptionsData.Data))
	}

	if eventsSubscriptionsData.Data[0].Name != "name-1" {
		t.Fatalf("Expected Name to be name-1, got %s", eventsSubscriptionsData.Data[0].Name)
	}

	if eventsSubscriptionsData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", eventsSubscriptionsData.Message)
	}
}
