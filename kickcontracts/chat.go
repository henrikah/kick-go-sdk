package kickcontracts

import (
	"context"

	"github.com/henrikah/kick-go-sdk/kickapitypes"
)

// Chat handles sending messages in a broadcaster's chat channel.
//
// All examples use context.TODO() as a placeholder. Replace with a proper
// context (with timeout/cancel) in production code.
type Chat interface {
	// SendChatMessageAsUser sends a chat message as a user in a broadcaster's channel.
	//
	// replyToMessageID is optional.
	//
	// Example:
	//
	//	sendChatResponse, err := client.Chat().SendChatMessageAsUser(context.TODO(), accessToken, 123, nil, "Hello chat!")
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
	//	sendChatResponse, err := client.Chat().SendChatMessageAsBot(context.TODO(), accessToken, nil, "Hello from bot!")
	//	if err != nil {
	//	    log.Printf("could not send bot chat message: %v", err)
	//	    return nil, err
	//	}
	SendChatMessageAsBot(ctx context.Context, accessToken string, replyToMessageID *string, message string) (*kickapitypes.SendChatResponse, error)
}
