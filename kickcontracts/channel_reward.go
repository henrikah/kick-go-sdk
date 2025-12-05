package kickcontracts

import (
	"context"

	"github.com/henrikah/kick-go-sdk/kickapitypes"
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
	//	channelRewards, err := client.ChannelReward().GetChannelRewards(context.TODO(), accessToken)
	//	if err != nil {
	//	    log.Printf("could not get channel rewards: %v", err)
	//	    return nil, err
	//	}
	GetChannelRewards(ctx context.Context, accessToken string) (*kickapitypes.ChannelRewards, error)

	// CreateChannelReward creates a channel reward for the channel.
	//
	// Example:
	//
	//	channelRewardData := kickapitypes.CreateChannelReward{
	//		Cost: 100,
	//		Title: "Channel Reward Title"
	//	}
	//
	//	createChannelRewardResult, err := client.ChannelReward().CreateChannelReward(context.TODO(), accessToken, channelRewardData)
	//	if err != nil {
	//	    log.Printf("could not create channel reward: %v", err)
	//	    return nil, err
	//	}
	CreateChannelReward(ctx context.Context, accessToken string, channelRewardData kickapitypes.CreateChannelReward) (*kickapitypes.ChannelReward, error)

	// DeleteChannelReward deletes a channel reward for the channel.
	//
	// Example:
	//
	//	channelRewardID := "some-reward-id"
	//
	//	err := client.ChannelReward().DeleteChannelReward(context.TODO(), accessToken, channelRewardID)
	//	if err != nil {
	//	    log.Printf("could not delete channel reward: %v", err)
	//	    return nil, err
	//	}
	DeleteChannelReward(ctx context.Context, accessToken string, rewardID string) error

	// UpdateChannelReward deletes a channel reward for the channel.
	//
	// Example:
	//
	//	channelRewardID := "some-reward-id"
	//	cost := 100
	//	title := "Channel Reward Title"
	//	channelRewardData := kickapitypes.UpdateChannelReward{
	//		Cost: &cost,
	//		Title: &title,
	//	}
	//	updateChannelRewardData, err := client.ChannelReward().UpdateChannelReward(context.TODO(), accessToken, channelRewardID, channelRewardData)
	//	if err != nil {
	//	    log.Printf("could not update channel reward: %v", err)
	//	    return nil, err
	//	}
	UpdateChannelReward(ctx context.Context, accessToken string, channelRewardID string, channelRewardData kickapitypes.UpdateChannelReward) (*kickapitypes.ChannelReward, error)
}
