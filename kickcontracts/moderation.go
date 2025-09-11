package kickcontracts

import (
	"context"

	"github.com/henrikah/kick-go-sdk/kickapitypes"
)

// Moderation handles user moderation actions such as timeouts and bans.
//
// All examples use context.TODO() as a placeholder. Replace with a proper
// context (with timeout/cancel) in production code.
type Moderation interface {
	// TimeOutUser temporarily restricts a user's ability to chat.
	//
	// reason is optional.
	//
	// Example:
	//
	//	moderationResponse, err := client.Moderation().TimeOutUser(context.TODO(), accessToken, broadcasterID, userID, 600, nil)
	//	if err != nil {
	//	    log.Printf("could not timeout user: %v", err)
	//	    return nil, err
	//	}
	TimeOutUser(ctx context.Context, accessToken string, broadcasterUserID, userID, durationInSeconds int, reason *string) (*kickapitypes.ModerationResponse, error)

	// BanUser permanently bans a user from a broadcaster's channel.
	//
	// reason is optional.
	//
	// Example:
	//
	//	moderationResponse, err := client.Moderation().BanUser(context.TODO(), accessToken, broadcasterID, userID, nil)
	//	if err != nil {
	//	    log.Printf("could not ban user: %v", err)
	//	    return nil, err
	//	}
	BanUser(ctx context.Context, accessToken string, broadcasterUserID, userID int, reason *string) (*kickapitypes.ModerationResponse, error)

	// UnbanUser removes a ban from a user.
	//
	// Example:
	//
	//	moderationResponse, err := client.Moderation().UnbanUser(context.TODO(), accessToken, broadcasterID, userID)
	//	if err != nil {
	//	    log.Printf("could not unban user: %v", err)
	//	    return nil, err
	//	}
	UnbanUser(ctx context.Context, accessToken string, broadcasterUserID, userID int) (*kickapitypes.ModerationResponse, error)
}
