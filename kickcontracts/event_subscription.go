package kickcontracts

import (
	"context"

	"github.com/henrikah/kick-go-sdk/enums/kickwebhookenum"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
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
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{...})
	//	if err != nil {
	//	    log.Printf("could not create APIClient: %v", err)
	//	    return nil, err
	//	}
	//
	//	subscriptions, err := client.EventsSubscription().GetEventSubscriptions(context.TODO(), accessToken)
	//	if err != nil {
	//	    log.Printf("could not get event subscriptions: %v", err)
	//	    return nil, err
	//	}
	GetEventSubscriptions(ctx context.Context, accessToken string) (*kickapitypes.EventSubscription, error)

	// CreateEventSubscriptions creates new event subscriptions for the authenticated user.
	//
	// Example:
	//
	//	events := []kickwebhookenum.WebhookType{kickwebhookenum.ChatMessageSent}
	//	createEventSubscriptionsResponse, err := client.EventsSubscription().CreateEventSubscriptions(context.TODO(), accessToken, events)
	//	if err != nil {
	//	    log.Printf("could not create event subscriptions: %v", err)
	//	    return nil, err
	//	}
	CreateEventSubscriptions(ctx context.Context, accessToken string, events []kickwebhookenum.WebhookType) (*kickapitypes.CreateEventSubscriptionsResponse, error)

	// CreateEventSubscriptionsAsApp creates new event subscriptions as the app for a given broadcaster.
	//
	// broadcasterUserID specifies which broadcaster the subscriptions are for.
	//
	// Example:
	//
	//	createEventSubscriptionsResponse, err := client.EventsSubscription().CreateEventSubscriptionsAsApp(context.TODO(), accessToken, 12345, []kickwebhookenum.WebhookType{kickwebhookenum.ChatMessageSent})
	//	if err != nil {
	//	    log.Printf("could not create app event subscriptions: %v", err)
	//	    return nil, err
	//	}
	CreateEventSubscriptionsAsApp(ctx context.Context, accessToken string, broadcasterUserID int, events []kickwebhookenum.WebhookType) (*kickapitypes.CreateEventSubscriptionsResponse, error)

	// DeleteEventSubscriptions deletes event subscriptions by their IDs.
	//
	// Example:
	//
	//	err := client.EventsSubscription().DeleteEventSubscriptions(context.TODO(), accessToken, []string{"subscription-id-1", "subscription-id-2"})
	//	if err != nil {
	//	    log.Printf("could not delete event subscriptions: %v", err)
	//	    return err
	//	}
	DeleteEventSubscriptions(ctx context.Context, accessToken string, subscriptionIDs []string) error
}
