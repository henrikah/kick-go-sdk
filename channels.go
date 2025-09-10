package kick

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/henrikah/kick-go-sdk/internal/endpoints"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickerrors"
)

// channel handles operations related to broadcaster channels.
//
// All examples use context.TODO() as a placeholder. Replace with a proper
// context (with timeout/cancel) in production code.
type channel interface {
	// GetChannelsByBroadcasterUserID retrieves multiple channels by their broadcaster user IDs.
	//
	// Example:
	//
	//	channels, err := client.Channel().GetChannelsByBroadcasterUserID(context.TODO(), accessToken, []int64{123, 456})
	//	if err != nil {
	//	    log.Printf("could not get channels by broadcaster IDs: %v", err)
	//	    return nil, err
	//	}
	GetChannelsByBroadcasterUserID(ctx context.Context, accessToken string, broadcasterUserIDs []int64) (*kickapitypes.Channels, error)

	// GetChannelByBroadcasterUserID retrieves a single channel by its broadcaster user ID.
	//
	// Example:
	//
	//	channel, err := client.Channel().GetChannelByBroadcasterUserID(context.TODO(), accessToken, 123)
	//	if err != nil {
	//	    log.Printf("could not get channel by broadcaster ID: %v", err)
	//	    return nil, err
	//	}
	GetChannelByBroadcasterUserID(ctx context.Context, accessToken string, broadcasterUserID int64) (*kickapitypes.Channels, error)

	// GetCurrentBroadcasterChannel retrieves the channel associated with the currently authenticated broadcaster.
	//
	// Example:
	//
	//	channel, err := client.Channel().GetCurrentBroadcasterChannel(context.TODO(), accessToken)
	//	if err != nil {
	//	    log.Printf("could not get current broadcaster channel: %v", err)
	//	    return nil, err
	//	}
	GetCurrentBroadcasterChannel(ctx context.Context, accessToken string) (*kickapitypes.Channels, error)

	// GetChannelsByBroadcasterSlug retrieves multiple channels by their slugs.
	//
	// Example:
	//
	//	channels, err := client.Channel().GetChannelsByBroadcasterSlug(context.TODO(), accessToken, []string{"slug1", "slug2"})
	//	if err != nil {
	//	    log.Printf("could not get channels by slug: %v", err)
	//	    return nil, err
	//	}
	GetChannelsByBroadcasterSlug(ctx context.Context, accessToken string, slugs []string) (*kickapitypes.Channels, error)

	// UpdateChannel updates a broadcaster's channel with the provided data.
	//
	// Example:
	//
	//	err := client.Channel().UpdateChannel(context.TODO(), accessToken, kickapitypes.UpdateChannelRequest{...})
	//	if err != nil {
	//	    log.Printf("could not update channel: %v", err)
	//	    return err
	//	}
	UpdateChannel(ctx context.Context, accessToken string, updateChannelData kickapitypes.UpdateChannelRequest) error
}

type channelClient struct {
	client *apiClient
}

func newChannelClient(client *apiClient) channel {
	return &channelClient{
		client: client,
	}
}

func (c *channelClient) GetChannelsByBroadcasterUserID(ctx context.Context, accessToken string, broadcasterUserIDs []int64) (*kickapitypes.Channels, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}
	if err := kickerrors.ValidateMaxItems("broadcasterUserIDs", broadcasterUserIDs, 50); err != nil {
		return nil, err
	}

	channelsURL, err := url.Parse(endpoints.ViewChannelsDetailsURL())
	if err != nil {
		return nil, err
	}

	queryParams := channelsURL.Query()
	for _, id := range broadcasterUserIDs {
		queryParams.Add("broadcaster_user_id", strconv.FormatInt(id, 10))
	}
	channelsURL.RawQuery = queryParams.Encode()

	var channelsData kickapitypes.Channels

	if err = c.client.makeJSONRequest(ctx, http.MethodGet, channelsURL.String(), nil, &accessToken, &channelsData); err != nil {
		return nil, err
	}

	return &channelsData, nil
}

func (c *channelClient) GetChannelByBroadcasterUserID(ctx context.Context, accessToken string, broadcasterUserID int64) (*kickapitypes.Channels, error) {
	return c.GetChannelsByBroadcasterUserID(ctx, accessToken, []int64{broadcasterUserID})
}

func (c *channelClient) GetCurrentBroadcasterChannel(ctx context.Context, accessToken string) (*kickapitypes.Channels, error) {
	return c.GetChannelsByBroadcasterUserID(ctx, accessToken, nil)
}

func (c *channelClient) GetChannelsByBroadcasterSlug(ctx context.Context, accessToken string, slugs []string) (*kickapitypes.Channels, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}
	if err := kickerrors.ValidateMaxItems("slugs", slugs, 50); err != nil {
		return nil, err
	}
	if err := kickerrors.ValidateMinItems("slugs", slugs, 1); err != nil {
		return nil, err
	}

	channelsURL, err := url.Parse(endpoints.ViewChannelsDetailsURL())
	if err != nil {
		return nil, err
	}

	queryParams := channelsURL.Query()
	for _, slug := range slugs {
		if err = kickerrors.ValidateMaxCharacters("slugs", slug, 25); err != nil {
			return nil, err
		}
		queryParams.Add("slug", slug)
	}
	channelsURL.RawQuery = queryParams.Encode()

	var channelsData kickapitypes.Channels

	if err = c.client.makeJSONRequest(ctx, http.MethodGet, channelsURL.String(), nil, &accessToken, &channelsData); err != nil {
		return nil, err
	}

	return &channelsData, nil
}

func (c *channelClient) UpdateChannel(ctx context.Context, accessToken string, updateChannelData kickapitypes.UpdateChannelRequest) error {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return err
	}
	if err := kickerrors.ValidateMinValue("categoryID", updateChannelData.CategoryID, 0); err != nil {
		return err
	}

	channelURL, err := url.Parse(endpoints.UpdateChannelDetailsURL())
	if err != nil {
		return err
	}

	if err = c.client.makeJSONRequest(ctx, http.MethodPatch, channelURL.String(), updateChannelData, &accessToken, nil); err != nil {
		return err
	}

	return nil
}
