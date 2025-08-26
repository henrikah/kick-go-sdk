package kick

import (
	"github.com/henrikah/kick-go-sdk/internal/httpclient"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickerrors"
)

type apiClient struct {
	Category           category
	Channel            channel
	Chat               chat
	EventsSubscription eventsSubscription
	Livestream         livestream
	Moderation         moderation
	OAuth              oAuth
	PublicKey          publicKey
	User               user
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

	client.Category = newCategoryClient(client)
	client.Channel = newChannelClient(client)
	client.Chat = newChatClient(client)
	client.EventsSubscription = newEventsSubscriptionClient(client)
	client.Livestream = newLivestreamClient(client)
	client.Moderation = newModerationClient(client)
	client.OAuth = newOAuthClient(client)
	client.PublicKey = newPublicKeyClient(client)
	client.User = newUserClient(client)

	return client, nil
}
