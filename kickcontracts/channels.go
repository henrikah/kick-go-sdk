package kickcontracts

import (
	"context"

	"github.com/henrikah/kick-go-sdk/kickapitypes"
)

// Channel handles operations related to broadcaster channels.
//
// All examples use context.TODO() as a placeholder. Replace with a proper
// context (with timeout/cancel) in production code.
type Channel interface {
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
