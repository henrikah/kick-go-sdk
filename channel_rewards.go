package kick

import (
	"context"
	"net/http"
	"net/url"

	"github.com/henrikah/kick-go-sdk/v2/internal/endpoints"
	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
	"github.com/henrikah/kick-go-sdk/v2/kickcontracts"
	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
	"github.com/henrikah/kick-go-sdk/v2/kickfilters"
)

type channelRewardClient struct {
	client *apiClient
}

func newChannelRewardService(client *apiClient) kickcontracts.ChannelReward {
	return &channelRewardClient{
		client: client,
	}
}

func (c *channelRewardClient) GetChannelRewards(ctx context.Context, accessToken string) (*kickapitypes.ChannelRewards, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}

	var channelRewardsData kickapitypes.ChannelRewards

	if err := c.client.requester.MakeJSONRequest(ctx, http.MethodGet, endpoints.ViewChannelRewardsURL(), nil, &accessToken, &channelRewardsData); err != nil {
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

	if err := c.client.requester.MakeJSONRequest(ctx, http.MethodPost, endpoints.CreateChannelRewardURL(), channelRewardData, &accessToken, &channelRewardResponseData); err != nil {
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

	err := c.client.requester.MakeDeleteRequest(ctx, endpoints.DeleteChannelRewardURL(rewardID), &accessToken, nil)
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

	if err := c.client.requester.MakeJSONRequest(ctx, http.MethodPatch, endpoints.UpdateChannelRewardURL(channelRewardID), channelRewardData, &accessToken, &channelRewardResponseData); err != nil {
		return nil, err
	}

	return &channelRewardResponseData, nil
}

func (c *channelRewardClient) GetChannelRewardRedemptions(ctx context.Context, accessToken string, filters kickfilters.RewardRedemptionsFilter) (*kickapitypes.ChannelRewardRedemptions, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}

	getChannelRewardRedemptionsURL, err := url.Parse(endpoints.ViewChannelRewardRedemptionsURL())
	if err != nil {
		return nil, err
	}

	if filters != nil {
		queryParams, filterErr := filters.ToQueryString()
		if filterErr != nil {
			return nil, filterErr
		}

		getChannelRewardRedemptionsURL.RawQuery = queryParams.Encode()
	}

	var channelRewardRedemptionsData kickapitypes.ChannelRewardRedemptions

	if err := c.client.requester.MakeGetRequest(ctx, getChannelRewardRedemptionsURL.String(), &accessToken, &channelRewardRedemptionsData); err != nil {
		return nil, err
	}

	return &channelRewardRedemptionsData, nil
}

func (c *channelRewardClient) AcceptRewardRedemption(ctx context.Context, accessToken string, redemptionIDs []string) (*kickapitypes.RedemptionDecision, error) {
	acceptRewardRedemptionURL, err := url.Parse(endpoints.AcceptChannelRewardRedemptionsURL())
	if err != nil {
		return nil, err
	}
	return c.channelRewardRedemptionDecision(ctx, accessToken, acceptRewardRedemptionURL.String(), redemptionIDs)
}
func (c *channelRewardClient) RejectRewardRedemption(ctx context.Context, accessToken string, redemptionIDs []string) (*kickapitypes.RedemptionDecision, error) {
	rejectRewardRedemptionURL, err := url.Parse(endpoints.RejectChannelRewardRedemptionsURL())
	if err != nil {
		return nil, err
	}
	return c.channelRewardRedemptionDecision(ctx, accessToken, rejectRewardRedemptionURL.String(), redemptionIDs)
}

func (c *channelRewardClient) channelRewardRedemptionDecision(ctx context.Context, accessToken string, url string, redemptionIDs []string) (*kickapitypes.RedemptionDecision, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}

	if err := kickerrors.ValidateMinItems("redemptionIDs", redemptionIDs, 1); err != nil {
		return nil, err
	}

	if err := kickerrors.ValidateMaxItems("redemptionIDs", redemptionIDs, 25); err != nil {
		return nil, err
	}

	redemptionData := kickapitypes.RedemptionsIDs{
		IDs: redemptionIDs,
	}
	var redemptionDecisionResponseData kickapitypes.RedemptionDecision

	if err := c.client.requester.MakeJSONRequest(ctx, http.MethodPost, url, redemptionData, &accessToken, &redemptionDecisionResponseData); err != nil {
		return nil, err
	}

	return &redemptionDecisionResponseData, nil
}
