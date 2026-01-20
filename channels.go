package kick

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/henrikah/kick-go-sdk/v2/internal/endpoints"
	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
	"github.com/henrikah/kick-go-sdk/v2/kickcontracts"
	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
)

type channelClient struct {
	client *apiClient
}

func newChannelService(client *apiClient) kickcontracts.Channel {
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

	if err = c.client.requester.MakeJSONRequest(ctx, http.MethodGet, channelsURL.String(), nil, &accessToken, &channelsData); err != nil {
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

	if err = c.client.requester.MakeJSONRequest(ctx, http.MethodGet, channelsURL.String(), nil, &accessToken, &channelsData); err != nil {
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

	if err = c.client.requester.MakeJSONRequest(ctx, http.MethodPatch, channelURL.String(), updateChannelData, &accessToken, nil); err != nil {
		return err
	}

	return nil
}
