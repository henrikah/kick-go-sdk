package kick

import (
	"context"
	"net/http"

	"github.com/henrikah/kick-go-sdk/v2/internal/endpoints"
	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
	"github.com/henrikah/kick-go-sdk/v2/kickcontracts"
	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
)

type moderationClient struct {
	client *apiClient
}

func newModerationService(client *apiClient) kickcontracts.Moderation {
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

	if err := c.client.requester.MakeJSONRequest(ctx, http.MethodPost, endpoints.BanUserURL(), moderationRequest, &accessToken, &moderationResponse); err != nil {
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

	if err := c.client.requester.MakeJSONRequest(ctx, http.MethodDelete, endpoints.LiftBanURL(), moderationRequest, &accessToken, &moderationResponse); err != nil {
		return nil, err
	}

	return &moderationResponse, nil
}
