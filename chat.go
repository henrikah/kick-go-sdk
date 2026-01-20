package kick

import (
	"context"
	"net/http"

	"github.com/henrikah/kick-go-sdk/v2/internal/endpoints"
	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
	"github.com/henrikah/kick-go-sdk/v2/kickcontracts"
	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
)

type chatClient struct {
	client *apiClient
}

func newChatService(client *apiClient) kickcontracts.Chat {
	return &chatClient{
		client: client,
	}
}

func (c *chatClient) SendChatMessageAsUser(ctx context.Context, accessToken string, broadcasterUserID int, replyToMessageID *string, message string) (*kickapitypes.SendChatResponse, error) {
	if err := kickerrors.ValidateBroadcasterUserID(broadcasterUserID); err != nil {
		return nil, err
	}

	chatRequest := kickapitypes.SendChatRequest{
		BroadcasterUserID: &broadcasterUserID,
		Content:           message,
		ReplyToMessageID:  replyToMessageID,
		Type:              "user",
	}

	return c.sendChatMessage(ctx, accessToken, chatRequest)
}

func (c *chatClient) SendChatMessageAsBot(ctx context.Context, accessToken string, replyToMessageID *string, message string) (*kickapitypes.SendChatResponse, error) {
	chatRequest := kickapitypes.SendChatRequest{
		Content:          message,
		ReplyToMessageID: replyToMessageID,
		Type:             "bot",
	}

	return c.sendChatMessage(ctx, accessToken, chatRequest)
}

func (c *chatClient) DeleteChatMessage(ctx context.Context, accessToken string, messageID string) error {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return err
	}
	if err := kickerrors.ValidateNotEmpty("messageID", messageID); err != nil {
		return err
	}

	err := c.client.requester.MakeDeleteRequest(ctx, endpoints.DeleteChatMessageURL(messageID), &accessToken, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *chatClient) sendChatMessage(ctx context.Context, accessToken string, chatRequest kickapitypes.SendChatRequest) (*kickapitypes.SendChatResponse, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}
	if err := kickerrors.ValidateChatMessage(chatRequest.Content); err != nil {
		return nil, err
	}

	var chatResponse kickapitypes.SendChatResponse

	if err := c.client.requester.MakeJSONRequest(ctx, http.MethodPost, endpoints.SendChatMessageURL(), chatRequest, &accessToken, &chatResponse); err != nil {
		return nil, err
	}

	return &chatResponse, nil
}
