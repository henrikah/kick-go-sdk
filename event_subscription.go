package kick

import (
	"context"
	"net/http"
	"net/url"

	"github.com/henrikah/kick-go-sdk/enums/kickwebhookenum"
	"github.com/henrikah/kick-go-sdk/internal/endpoints"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickerrors"
)

// eventsSubscription handles managing Kick webhook event subscriptions.
//
// All examples use context.TODO() as a placeholder. Replace with a proper
// context (with timeout/cancel) in production code.
type eventsSubscription interface {
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

type eventsSubscriptionClient struct {
	client *apiClient
}

func newEventsSubscriptionClient(client *apiClient) eventsSubscription {
	return &eventsSubscriptionClient{
		client: client,
	}
}

func (c *eventsSubscriptionClient) GetEventSubscriptions(ctx context.Context, accessToken string) (*kickapitypes.EventSubscription, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}

	var result kickapitypes.EventSubscription

	err := c.client.makeJSONRequest(ctx, http.MethodGet, endpoints.ViewEventsSubscriptionsDetailsURL(), nil, &accessToken, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *eventsSubscriptionClient) CreateEventSubscriptions(ctx context.Context, accessToken string, events []kickwebhookenum.WebhookType) (*kickapitypes.CreateEventSubscriptionsResponse, error) {
	return c.createEventSubscriptions(ctx, accessToken, nil, events)
}

func (c *eventsSubscriptionClient) CreateEventSubscriptionsAsApp(ctx context.Context, accessToken string, broadcasterUserID int, events []kickwebhookenum.WebhookType) (*kickapitypes.CreateEventSubscriptionsResponse, error) {
	return c.createEventSubscriptions(ctx, accessToken, &broadcasterUserID, events)
}

func (c *eventsSubscriptionClient) createEventSubscriptions(ctx context.Context, accessToken string, broadcasterUserID *int, events []kickwebhookenum.WebhookType) (*kickapitypes.CreateEventSubscriptionsResponse, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}
	if err := kickerrors.ValidateMinItems("events", events, 1); err != nil {
		return nil, err
	}

	createEventSubscriptionsRequest := kickapitypes.CreateEventSubscriptionRequest{
		BroadcasterUserID: broadcasterUserID,
		Events:            make([]kickapitypes.EventObject, len(events)),
		Method:            "webhook",
	}

	for i := range events {
		createEventSubscriptionsRequest.Events[i] = kickapitypes.EventObject{
			Name:    string(events[i]),
			Version: 1,
		}
	}

	var createEventSubscriptionsResponse kickapitypes.CreateEventSubscriptionsResponse
	err := c.client.makeJSONRequest(ctx, http.MethodPost, endpoints.ViewEventsSubscriptionsDetailsURL(), createEventSubscriptionsRequest, &accessToken, &createEventSubscriptionsResponse)
	if err != nil {
		return nil, err
	}

	return &createEventSubscriptionsResponse, nil
}

func (c *eventsSubscriptionClient) DeleteEventSubscriptions(ctx context.Context, accessToken string, subscriptionIDs []string) error {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return err
	}
	if err := kickerrors.ValidateMinItems("subscriptionIDs", subscriptionIDs, 1); err != nil {
		return err
	}

	deleteEventsSubscriptionsURL, err := url.Parse(endpoints.RemoveEventsSubscriptionsURL())
	if err != nil {
		return err
	}

	queryParams := deleteEventsSubscriptionsURL.Query()
	for _, id := range subscriptionIDs {
		queryParams.Add("id", id)
	}
	deleteEventsSubscriptionsURL.RawQuery = queryParams.Encode()

	err = c.client.makeDeleteRequest(ctx, deleteEventsSubscriptionsURL.String(), &accessToken, nil)
	if err != nil {
		return err
	}

	return nil
}
