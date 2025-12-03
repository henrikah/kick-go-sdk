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
	//	    log.Printf("could not get channel rewards: %v", err)
	//	    return nil, err
	//	}
	CreateChannelReward(ctx context.Context, accessToken string, channelRewardData kickapitypes.CreateChannelReward) (*kickapitypes.ChannelReward, error)
}
