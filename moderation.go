package kick

import (
	"context"
	"net/http"

	"github.com/henrikah/kick-go-sdk/internal/endpoints"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickerrors"
)

// moderation handles user moderation actions such as timeouts and bans.
//
// All examples use context.TODO() as a placeholder. Replace with a proper
// context (with timeout/cancel) in production code.
type moderation interface {
	// TimeOutUser temporarily restricts a user's ability to chat.
	//
	// reason is optional.
	//
	// Example:
	//
	//	moderationResponse, err := client.Moderation.TimeOutUser(context.TODO(), accessToken, broadcasterID, userID, 600, nil)
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
	//	moderationResponse, err := client.Moderation.BanUser(context.TODO(), accessToken, broadcasterID, userID, nil)
	//	if err != nil {
	//	    log.Printf("could not ban user: %v", err)
	//	    return nil, err
	//	}
	BanUser(ctx context.Context, accessToken string, broadcasterUserID, userID int, reason *string) (*kickapitypes.ModerationResponse, error)

	// UnbanUser removes a ban from a user.
	//
	// Example:
	//
	//	moderationResponse, err := client.Moderation.UnbanUser(context.TODO(), accessToken, broadcasterID, userID)
	//	if err != nil {
	//	    log.Printf("could not unban user: %v", err)
	//	    return nil, err
	//	}
	UnbanUser(ctx context.Context, accessToken string, broadcasterUserID, userID int) (*kickapitypes.ModerationResponse, error)
}

type moderationClient struct {
	client *apiClient
}

func newModerationClient(client *apiClient) moderation {
	return &moderationClient{
		client: client,
	}
}
func (c *moderationClient) TimeOutUser(ctx context.Context, accessToken string, broadcasterUserID int, userID int, durationInSeconds int, reason *string) (*kickapitypes.ModerationResponse, error) {

	timeoutRequest := kickapitypes.ModerationRequest{
		BroadcasterUserID: broadcasterUserID,
		Duration:          &durationInSeconds,
		Reason:            reason,
		UserID:            userID,
	}

	return c.moderateUser(ctx, accessToken, timeoutRequest)
}

func (c *moderationClient) BanUser(ctx context.Context, accessToken string, broadcasterUserID int, userID int, reason *string) (*kickapitypes.ModerationResponse, error) {

	banRequest := kickapitypes.ModerationRequest{
		BroadcasterUserID: broadcasterUserID,
		Duration:          nil,
		Reason:            reason,
		UserID:            userID,
	}

	return c.moderateUser(ctx, accessToken, banRequest)
}

func (c *moderationClient) moderateUser(ctx context.Context, accessToken string, moderationRequest kickapitypes.ModerationRequest) (*kickapitypes.ModerationResponse, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}
	if err := kickerrors.ValidateBroadcasterUserID(moderationRequest.BroadcasterUserID); err != nil {
		return nil, err
	}
	if err := kickerrors.ValidateUserID(moderationRequest.UserID); err != nil {
		return nil, err
	}
	if moderationRequest.Duration != nil {
		if err := kickerrors.ValidateBetween("duration", *moderationRequest.Duration, 1, 10080); err != nil {
			return nil, err
		}
	}

	var moderationResponse kickapitypes.ModerationResponse

	if err := c.client.makeJSONRequest(ctx, http.MethodPost, endpoints.BanUserURL(), moderationRequest, &accessToken, &moderationResponse); err != nil {
		return nil, err
	}

	return &moderationResponse, nil
}

func (c *moderationClient) UnbanUser(ctx context.Context, accessToken string, broadcasterUserID int, userID int) (*kickapitypes.ModerationResponse, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}
	if err := kickerrors.ValidateBroadcasterUserID(broadcasterUserID); err != nil {
		return nil, err
	}
	if err := kickerrors.ValidateUserID(userID); err != nil {
		return nil, err
	}

	moderationRequest := kickapitypes.UnBanRequest{
		BroadcasterUserID: broadcasterUserID,
		UserID:            userID,
	}

	var moderationResponse kickapitypes.ModerationResponse

	if err := c.client.makeJSONRequest(ctx, http.MethodDelete, endpoints.LiftBanURL(), moderationRequest, &accessToken, &moderationResponse); err != nil {
		return nil, err
	}

	return &moderationResponse, nil
}
