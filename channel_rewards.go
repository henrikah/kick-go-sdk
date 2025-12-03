package kick

import (
	"context"
	"net/http"

	"github.com/henrikah/kick-go-sdk/internal/endpoints"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickcontracts"
	"github.com/henrikah/kick-go-sdk/kickerrors"
)

type channelRewardClient struct {
	client *apiClient
}

func newChannelRewardClient(client *apiClient) kickcontracts.ChannelReward {
	return &channelRewardClient{
		client: client,
	}
}

func (c *channelRewardClient) GetChannelRewards(ctx context.Context, accessToken string) (*kickapitypes.ChannelRewards, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}

	var channelRewardsData kickapitypes.ChannelRewards

	if err := c.client.makeJSONRequest(ctx, http.MethodGet, endpoints.ViewChannelRewards(), nil, &accessToken, &channelRewardsData); err != nil {
		return nil, err
	}

	return &channelRewardsData, nil
}

func (c *channelRewardClient) CreateChannelReward(ctx context.Context, accessToken string, channelRewardData kickapitypes.CreateChannelReward) (*kickapitypes.ChannelReward, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}

	if err := kickerrors.ValidateMinValue("Cost", channelRewardData.Cost, 1); err != nil {
		return nil, err
	}

	var channelRewardResponseData kickapitypes.ChannelReward

	if err := c.client.makeJSONRequest(ctx, http.MethodPost, endpoints.CreateChannelReward(), channelRewardData, &accessToken, &channelRewardResponseData); err != nil {
		return nil, err
	}

	return &channelRewardResponseData, nil
}
