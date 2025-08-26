package kick

import (
	"context"
	"net/http"

	"github.com/henrikah/kick-go-sdk/internal/endpoints"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickerrors"
)

// chat handles sending messages in a broadcaster's chat channel.
//
// All examples use context.TODO() as a placeholder. Replace with a proper
// context (with timeout/cancel) in production code.
type chat interface {
	// SendChatMessageAsUser sends a chat message as a user in a broadcaster's channel.
	//
	// replyToMessageID is optional.
	//
	// Example:
	//
	//	sendChatResponse, err := client.Chat.SendChatMessageAsUser(context.TODO(), accessToken, 123, nil, "Hello chat!")
	//	if err != nil {
	//	    log.Printf("could not send chat message: %v", err)
	//	    return nil, err
	//	}
	SendChatMessageAsUser(ctx context.Context, accessToken string, broadcasterUserID int, replyToMessageID *string, message string) (*kickapitypes.SendChatResponse, error)

	// SendChatMessageAsBot sends a chat message as a bot in a broadcaster's channel.
	//
	// replyToMessageID is optional.
	//
	// Example:
	//
	//	sendChatResponse, err := client.Chat.SendChatMessageAsBot(context.TODO(), accessToken, nil, "Hello from bot!")
	//	if err != nil {
	//	    log.Printf("could not send bot chat message: %v", err)
	//	    return nil, err
	//	}
	SendChatMessageAsBot(ctx context.Context, accessToken string, replyToMessageID *string, message string) (*kickapitypes.SendChatResponse, error)
}

type chatClient struct {
	client *apiClient
}

func newChatClient(client *apiClient) chat {
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

func (c *chatClient) sendChatMessage(ctx context.Context, accessToken string, chatRequest kickapitypes.SendChatRequest) (*kickapitypes.SendChatResponse, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}
	if err := kickerrors.ValidateChatMessage(chatRequest.Content); err != nil {
		return nil, err
	}

	var chatResponse kickapitypes.SendChatResponse

	if err := c.client.makeJSONRequest(ctx, http.MethodPost, endpoints.SendChatMessageURL(), chatRequest, &accessToken, &chatResponse); err != nil {
		return nil, err
	}

	return &chatResponse, nil
}
