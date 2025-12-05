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

	if err := c.client.makeJSONRequest(ctx, http.MethodGet, endpoints.ViewChannelRewardsURL(), nil, &accessToken, &channelRewardsData); err != nil {
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

	if err := kickerrors.ValidateMaxCharacters("Title", channelRewardData.Title, 50); err != nil {
		return nil, err
	}

	if err := kickerrors.ValidateNotNilPointer("Description", channelRewardData.Description); err == nil {
		if err := kickerrors.ValidateMaxCharacters("Description", *channelRewardData.Description, 200); err != nil {
			return nil, err
		}
	}

	var channelRewardResponseData kickapitypes.ChannelReward

	if err := c.client.makeJSONRequest(ctx, http.MethodPost, endpoints.CreateChannelRewardURL(), channelRewardData, &accessToken, &channelRewardResponseData); err != nil {
		return nil, err
	}

	return &channelRewardResponseData, nil
}

func (c *channelRewardClient) DeleteChannelReward(ctx context.Context, accessToken string, rewardID string) error {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return err
	}
	if err := kickerrors.ValidateNotEmpty("rewardID", rewardID); err != nil {
		return err
	}

	err := c.client.makeDeleteRequest(ctx, endpoints.DeleteChannelRewardURL(rewardID), &accessToken, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *channelRewardClient) UpdateChannelReward(ctx context.Context, accessToken string, channelRewardID string, channelRewardData kickapitypes.UpdateChannelReward) (*kickapitypes.ChannelReward, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}

	if err := kickerrors.ValidateNotNilPointer("Cost", channelRewardData.Cost); err == nil {
		if err := kickerrors.ValidateMinValue("Cost", *channelRewardData.Cost, 1); err != nil {
			return nil, err
		}
	}

	if err := kickerrors.ValidateNotNilPointer("Description", channelRewardData.Description); err == nil {
		if err := kickerrors.ValidateMaxCharacters("Description", *channelRewardData.Description, 200); err != nil {
			return nil, err
		}
	}

	if err := kickerrors.ValidateNotNilPointer("Title", channelRewardData.Title); err == nil {
		if err := kickerrors.ValidateMaxCharacters("Title", *channelRewardData.Title, 50); err != nil {
			return nil, err
		}
	}

	var channelRewardResponseData kickapitypes.ChannelReward

	if err := c.client.makeJSONRequest(ctx, http.MethodPatch, endpoints.UpdateChannelRewardURL(channelRewardID), channelRewardData, &accessToken, &channelRewardResponseData); err != nil {
		return nil, err
	}

	return &channelRewardResponseData, nil
}
