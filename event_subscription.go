package kick

import (
	"context"
	"net/http"
	"net/url"

	"github.com/henrikah/kick-go-sdk/enums/kickwebhookenum"
	"github.com/henrikah/kick-go-sdk/internal/endpoints"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickcontracts"
	"github.com/henrikah/kick-go-sdk/kickerrors"
)

type eventsSubscriptionClient struct {
	client *apiClient
}

func newEventsSubscriptionService(client *apiClient) kickcontracts.EventsSubscription {
	return &eventsSubscriptionClient{
		client: client,
	}
}

func (c *eventsSubscriptionClient) GetEventSubscriptions(ctx context.Context, accessToken string) (*kickapitypes.EventSubscription, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}

	var result kickapitypes.EventSubscription

	err := c.client.requester.MakeJSONRequest(ctx, http.MethodGet, endpoints.ViewEventsSubscriptionsDetailsURL(), nil, &accessToken, &result)
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
	err := c.client.requester.MakeJSONRequest(ctx, http.MethodPost, endpoints.ViewEventsSubscriptionsDetailsURL(), createEventSubscriptionsRequest, &accessToken, &createEventSubscriptionsResponse)
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

	err = c.client.requester.MakeDeleteRequest(ctx, deleteEventsSubscriptionsURL.String(), &accessToken, nil)
	if err != nil {
		return err
	}

	return nil
}
