package kick

import (
	"github.com/henrikah/kick-go-sdk/internal/httpclient"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickcontracts"
	"github.com/henrikah/kick-go-sdk/kickerrors"
)

type apiClient struct {
	category           kickcontracts.Category
	channel            kickcontracts.Channel
	chat               kickcontracts.Chat
	eventsSubscription kickcontracts.EventsSubscription
	kicks              kickcontracts.Kicks
	livestream         kickcontracts.Livestream
	moderation         kickcontracts.Moderation
	oAuth              kickcontracts.OAuth
	publicKey          kickcontracts.PublicKey
	user               kickcontracts.User
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
	client.kicks = newKicksClient(client)
	client.livestream = newLivestreamClient(client)
	client.moderation = newModerationClient(client)
	client.oAuth = newOAuthClient(client)
	client.publicKey = newPublicKeyClient(client)
	client.user = newUserClient(client)

	return client, nil
}

func (c *apiClient) Category() kickcontracts.Category {
	return c.category
}
func (c *apiClient) Channel() kickcontracts.Channel {
	return c.channel
}
func (c *apiClient) Chat() kickcontracts.Chat {
	return c.chat
}
func (c *apiClient) EventsSubscription() kickcontracts.EventsSubscription {
	return c.eventsSubscription
}
func (c *apiClient) Kicks() kickcontracts.Kicks {
	return c.kicks
}
func (c *apiClient) Livestream() kickcontracts.Livestream {
	return c.livestream
}
func (c *apiClient) Moderation() kickcontracts.Moderation {
	return c.moderation
}
func (c *apiClient) OAuth() kickcontracts.OAuth {
	return c.oAuth
}
func (c *apiClient) PublicKey() kickcontracts.PublicKey {
	return c.publicKey
}
func (c *apiClient) User() kickcontracts.User {
	return c.user
}
