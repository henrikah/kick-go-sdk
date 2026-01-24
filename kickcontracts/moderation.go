package kickcontracts

import (
	"context"

	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
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
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
	//	    HTTPClient: http.DefaultClient,
	//	})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	moderationResponse, err := client.Moderation().TimeOutUser(context.TODO(), accessToken, broadcasterID, userID, 600, nil)
	//	if err != nil {
	//		if apiError := kickerrors.IsAPIError(err); apiError != nil {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	TimeOutUser(ctx context.Context, accessToken string, broadcasterUserID, userID, durationInSeconds int, reason *string) (*kickapitypes.ModerationResponse, error)

	// BanUser permanently bans a user from a broadcaster's channel.
	//
	// reason is optional.
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
	//	moderationResponse, err := client.Moderation().BanUser(context.TODO(), accessToken, broadcasterID, userID, nil)
	//	if err != nil {
	//		if apiError := kickerrors.IsAPIError(err); apiError != nil {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	BanUser(ctx context.Context, accessToken string, broadcasterUserID, userID int, reason *string) (*kickapitypes.ModerationResponse, error)

	// UnbanUser removes a ban from a user.
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
	//	moderationResponse, err := client.Moderation().UnbanUser(context.TODO(), accessToken, broadcasterID, userID)
	//	if err != nil {
	//		if apiError := kickerrors.IsAPIError(err); apiError != nil {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	UnbanUser(ctx context.Context, accessToken string, broadcasterUserID, userID int) (*kickapitypes.ModerationResponse, error)
}
