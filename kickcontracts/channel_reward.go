package kickcontracts

import (
	"context"

	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
	"github.com/henrikah/kick-go-sdk/v2/kickfilters"
)

// ChannelReward handles operations related to a broadcaster channel rewards.
//
// All examples use context.TODO() as a placeholder. Replace with a proper
// context (with timeout/cancel) in production code.
type ChannelReward interface {
	// GetChannelRewards retrieves all active and inactive channel rewards from a channel.
	//
	// Example:
	//
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
	//	    HTTPClient: http.DefaultClient,
	//	})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	channelRewards, err := client.ChannelReward().GetChannelRewards(context.TODO(), accessToken)
	//	if err != nil {
	//		if apiError := kickerrors.IsAPIError(err); apiError != nil {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	GetChannelRewards(ctx context.Context, accessToken string) (*kickapitypes.ChannelRewards, error)

	// CreateChannelReward creates a channel reward for the channel.
	//
	// Example:
	//
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
	//	    HTTPClient: http.DefaultClient,
	//	})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	channelRewardData := kickapitypes.CreateChannelReward{
	//		Cost: 100,
	//		Title: "Channel Reward Title"
	//	}
	//
	//	createChannelRewardResult, err := client.ChannelReward().CreateChannelReward(context.TODO(), accessToken, channelRewardData)
	//	if err != nil {
	//		if apiError := kickerrors.IsAPIError(err); apiError != nil {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	CreateChannelReward(ctx context.Context, accessToken string, channelRewardData kickapitypes.CreateChannelReward) (*kickapitypes.ChannelReward, error)

	// DeleteChannelReward deletes a channel reward for the channel.
	//
	// Example:
	//
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
	//	    HTTPClient: http.DefaultClient,
	//	})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	channelRewardID := "some-reward-id"
	//
	//	err := client.ChannelReward().DeleteChannelReward(context.TODO(), accessToken, channelRewardID)
	//	if err != nil {
	//		if apiError := kickerrors.IsAPIError(err); apiError != nil {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	DeleteChannelReward(ctx context.Context, accessToken string, rewardID string) error

	// UpdateChannelReward deletes a channel reward for the channel.
	//
	// Example:
	//
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
	//	    HTTPClient: http.DefaultClient,
	//	})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	channelRewardID := "some-reward-id"
	//	cost := 100
	//	title := "Channel Reward Title"
	//	channelRewardData := kickapitypes.UpdateChannelReward{
	//		Cost: &cost,
	//		Title: &title,
	//	}
	//
	//	updateChannelRewardData, err := client.ChannelReward().UpdateChannelReward(context.TODO(), accessToken, channelRewardID, channelRewardData)
	//	if err != nil {
	//		if apiError := kickerrors.IsAPIError(err); apiError != nil {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	UpdateChannelReward(ctx context.Context, accessToken string, channelRewardID string, channelRewardData kickapitypes.UpdateChannelReward) (*kickapitypes.ChannelReward, error)

	// GetChannelRewardRedemptions gets redemptions for the channel.
	//
	// Example:
	//
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
	//	    HTTPClient: http.DefaultClient,
	//	})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	filters := kickfilters.NewRewardRedemptionsFilter().
	//	    WithStatus(kickchannelrewardstatus.Pending)
	//
	//	getChannelRewardRedemptionsData, err := client.ChannelReward().GetChannelRewardRedemptions(context.TODO(), accessToken, filters)
	//	if err != nil {
	//		if apiError := kickerrors.IsAPIError(err); apiError != nil {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	GetChannelRewardRedemptions(ctx context.Context, accessToken string, filters kickfilters.RewardRedemptionsFilter) (*kickapitypes.ChannelRewardRedemptions, error)

	// AcceptRewardRedemption accepts reward redemptions.
	//
	// Example:
	//
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
	//	    HTTPClient: http.DefaultClient,
	//	})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	redemptionIDs := []string{"id-1", "id-2"}
	//
	//	acceptChannelRewardRedemptionsData, err := client.ChannelReward().AcceptRewardRedemption(context.TODO(), accessToken, redemptionIDs)
	//	if err != nil {
	//		if apiError := kickerrors.IsAPIError(err); apiError != nil {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	AcceptRewardRedemption(ctx context.Context, accessToken string, redemptionIDs []string) (*kickapitypes.RedemptionDecision, error)

	// RejectRewardRedemption rejects reward redemptions.
	//
	// Example:
	//
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
	//	    HTTPClient: http.DefaultClient,
	//	})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	redemptionIDs := []string{"id-1", "id-2"}
	//
	//	acceptChannelRewardRedemptionsData, err := client.ChannelReward().RejectRewardRedemption(context.TODO(), accessToken, redemptionIDs)
	//	if err != nil {
	//		if apiError := kickerrors.IsAPIError(err); apiError != nil {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	RejectRewardRedemption(ctx context.Context, accessToken string, redemptionIDs []string) (*kickapitypes.RedemptionDecision, error)
}
