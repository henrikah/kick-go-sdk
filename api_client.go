package kick

import (
	"github.com/henrikah/kick-go-sdk/internal/transport"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickcontracts"
)

type apiClient struct {
	category           kickcontracts.Category
	channel            kickcontracts.Channel
	channelReward      kickcontracts.ChannelReward
	chat               kickcontracts.Chat
	eventsSubscription kickcontracts.EventsSubscription
	kicks              kickcontracts.Kicks
	livestream         kickcontracts.Livestream
	moderation         kickcontracts.Moderation
	publicKey          kickcontracts.PublicKey
	user               kickcontracts.User
	requester          *transport.Requester
}

// NewAPIClient creates a new APIClient instance with the provided configuration.
//
// Example:
//
//	apiClient, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
//		HTTPClient:   http.DefaultClient,
//	})
//	if err != nil {
//		log.Fatalf("could not create APIClient: %v", err)
//	}
func NewAPIClient(clientConfig kickapitypes.APIClientConfig) (*apiClient, error) {
	requester, err := transport.NewRequester(clientConfig.HTTPClient)
	if err != nil {
		return nil, err
	}

	client := &apiClient{
		requester: requester,
	}

	client.category = newCategoryService(client)
	client.channel = newChannelService(client)
	client.channelReward = newChannelRewardService(client)
	client.chat = newChatService(client)
	client.eventsSubscription = newEventsSubscriptionService(client)
	client.kicks = newKicksService(client)
	client.livestream = newLivestreamService(client)
	client.moderation = newModerationService(client)
	client.publicKey = newPublicKeyService(client)
	client.user = newUserService(client)

	return client, nil
}

func (c *apiClient) Category() kickcontracts.Category {
	return c.category
}

func (c *apiClient) Channel() kickcontracts.Channel {
	return c.channel
}

func (c *apiClient) ChannelReward() kickcontracts.ChannelReward {
	return c.channelReward
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

func (c *apiClient) PublicKey() kickcontracts.PublicKey {
	return c.publicKey
}

func (c *apiClient) User() kickcontracts.User {
	return c.user
}
