package kickcontracts

import (
	"context"

	"github.com/henrikah/kick-go-sdk/v2/enums/kickwebhookenum"
	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
)

// EventsSubscription handles managing Kick webhook event subscriptions.
//
// All examples use context.TODO() as a placeholder. Replace with a proper
// context (with timeout/cancel) in production code.
type EventsSubscription interface {
	// GetEventSubscriptions retrieves all event subscriptions for the authenticated user.
	//
	// Example:
	//
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
	//	    HTTPClient: http.DefaultClient,
	//	})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	subscriptions, err := client.EventsSubscription().GetEventSubscriptions(context.TODO(), accessToken)
	//	if err != nil {
	//		var apiErr *kickerrors.APIError
	//		if errors.As(err, &apiErr) {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	GetEventSubscriptions(ctx context.Context, accessToken string) (*kickapitypes.EventSubscription, error)

	// CreateEventSubscriptions creates new event subscriptions for the authenticated user.
	//
	// Example:
	//
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
	//	    HTTPClient: http.DefaultClient,
	//	})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	events := []kickwebhookenum.WebhookType{kickwebhookenum.ChatMessageSent}
	//	createEventSubscriptionsResponse, err := client.EventsSubscription().CreateEventSubscriptions(context.TODO(), accessToken, events)
	//	if err != nil {
	//		var apiErr *kickerrors.APIError
	//		if errors.As(err, &apiErr) {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	CreateEventSubscriptions(ctx context.Context, accessToken string, events []kickwebhookenum.WebhookType) (*kickapitypes.CreateEventSubscriptionsResponse, error)

	// CreateEventSubscriptionsAsApp creates new event subscriptions as the app for a given broadcaster.
	//
	// broadcasterUserID specifies which broadcaster the subscriptions are for.
	//
	// Example:
	//
	//
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
	//	    HTTPClient: http.DefaultClient,
	//	})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	createEventSubscriptionsResponse, err := client.EventsSubscription().CreateEventSubscriptionsAsApp(context.TODO(), accessToken, 12345, []kickwebhookenum.WebhookType{kickwebhookenum.ChatMessageSent})
	//	if err != nil {
	//		var apiErr *kickerrors.APIError
	//		if errors.As(err, &apiErr) {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	CreateEventSubscriptionsAsApp(ctx context.Context, accessToken string, broadcasterUserID int, events []kickwebhookenum.WebhookType) (*kickapitypes.CreateEventSubscriptionsResponse, error)

	// DeleteEventSubscriptions deletes event subscriptions by their IDs.
	//
	// Example:
	//
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
	//	    HTTPClient: http.DefaultClient,
	//	})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	err := client.EventsSubscription().DeleteEventSubscriptions(context.TODO(), accessToken, []string{"subscription-id-1", "subscription-id-2"})
	//	if err != nil {
	//		var apiErr *kickerrors.APIError
	//		if errors.As(err, &apiErr) {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	DeleteEventSubscriptions(ctx context.Context, accessToken string, subscriptionIDs []string) error
}
