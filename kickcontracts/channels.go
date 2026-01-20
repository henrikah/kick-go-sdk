package kickcontracts

import (
	"context"

	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
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
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
	//	    HTTPClient: http.DefaultClient,
	//	})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	channels, err := client.Channel().GetChannelsByBroadcasterUserID(context.TODO(), accessToken, []int64{123, 456})
	//	if err != nil {
	//		var apiErr *kickerrors.APIError
	//		if errors.As(err, &apiErr) {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	GetChannelsByBroadcasterUserID(ctx context.Context, accessToken string, broadcasterUserIDs []int64) (*kickapitypes.Channels, error)

	// GetChannelByBroadcasterUserID retrieves a single channel by its broadcaster user ID.
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
	//	channel, err := client.Channel().GetChannelByBroadcasterUserID(context.TODO(), accessToken, 123)
	//	if err != nil {
	//		var apiErr *kickerrors.APIError
	//		if errors.As(err, &apiErr) {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	GetChannelByBroadcasterUserID(ctx context.Context, accessToken string, broadcasterUserID int64) (*kickapitypes.Channels, error)

	// GetCurrentBroadcasterChannel retrieves the channel associated with the currently authenticated broadcaster.
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
	//	channel, err := client.Channel().GetCurrentBroadcasterChannel(context.TODO(), accessToken)
	//	if err != nil {
	//		var apiErr *kickerrors.APIError
	//		if errors.As(err, &apiErr) {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	GetCurrentBroadcasterChannel(ctx context.Context, accessToken string) (*kickapitypes.Channels, error)

	// GetChannelsByBroadcasterSlug retrieves multiple channels by their slugs.
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
	//	channels, err := client.Channel().GetChannelsByBroadcasterSlug(context.TODO(), accessToken, []string{"slug1", "slug2"})
	//	if err != nil {
	//		var apiErr *kickerrors.APIError
	//		if errors.As(err, &apiErr) {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	GetChannelsByBroadcasterSlug(ctx context.Context, accessToken string, slugs []string) (*kickapitypes.Channels, error)

	// UpdateChannel updates a broadcaster's channel with the provided data.
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
	//	err := client.Channel().UpdateChannel(context.TODO(), accessToken, kickapitypes.UpdateChannelRequest{...})
	//	if err != nil {
	//		var apiErr *kickerrors.APIError
	//		if errors.As(err, &apiErr) {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	UpdateChannel(ctx context.Context, accessToken string, updateChannelData kickapitypes.UpdateChannelRequest) error
}
