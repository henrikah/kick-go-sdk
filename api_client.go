package kick

import (
	"github.com/henrikah/kick-go-sdk/internal/httpclient"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickerrors"
)

type APIClient interface {
	Category() category
	Channel() channel
	Chat() chat
	EventsSubscription() eventsSubscription
	Livestream() livestream
	Moderation() moderation
	OAuth() oAuth
	PublicKey() publicKey
	User() user
}

type apiClient struct {
	category           category
	channel            channel
	chat               chat
	eventsSubscription eventsSubscription
	livestream         livestream
	moderation         moderation
	oAuth              oAuth
	publicKey          publicKey
	user               user
	clientID           string
	clientSecret       string
	httpClient         httpclient.ClientInterface
}

// NewAPIClient creates a new APIClient instance with the provided configuration.
//
// Example:
//
//	apiClient, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
//		ClientID:     "your-client-id",
//		ClientSecret: "your-client-secret",
//		HTTPClient:   http.DefaultClient,
//	})
//	if err != nil {
//		log.Fatalf("could not create APIClient: %v", err)
//	}
func NewAPIClient(clientConfig kickapitypes.APIClientConfig) (*apiClient, error) {
	if err := kickerrors.ValidateNotEmpty("ClientID", clientConfig.ClientID); err != nil {
		return nil, err
	}
	if err := kickerrors.ValidateNotEmpty("ClientSecret", clientConfig.ClientSecret); err != nil {
		return nil, err
	}

	if err := kickerrors.ValidateNotNil("HTTPClient", clientConfig.HTTPClient); err != nil {
		return nil, err
	}

	client := &apiClient{
		clientID:     clientConfig.ClientID,
		clientSecret: clientConfig.ClientSecret,
		httpClient:   clientConfig.HTTPClient,
	}

	client.category = newCategoryClient(client)
	client.channel = newChannelClient(client)
	client.chat = newChatClient(client)
	client.eventsSubscription = newEventsSubscriptionClient(client)
	client.livestream = newLivestreamClient(client)
	client.moderation = newModerationClient(client)
	client.oAuth = newOAuthClient(client)
	client.publicKey = newPublicKeyClient(client)
	client.user = newUserClient(client)

	return client, nil
}

func (c *apiClient) Category() category {
	return c.category
}
func (c *apiClient) Channel() channel {
	return c.channel
}
func (c *apiClient) Chat() chat {
	return c.chat
}
func (c *apiClient) EventsSubscription() eventsSubscription {
	return c.eventsSubscription
}
func (c *apiClient) Livestream() livestream {
	return c.livestream
}
func (c *apiClient) Moderation() moderation {
	return c.moderation
}
func (c *apiClient) OAuth() oAuth {
	return c.oAuth
}
func (c *apiClient) PublicKey() publicKey {
	return c.publicKey
}
func (c *apiClient) User() user {
	return c.user
}
