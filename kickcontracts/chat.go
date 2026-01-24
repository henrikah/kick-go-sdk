package kickcontracts

import (
	"context"

	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
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
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
	//	    HTTPClient: http.DefaultClient,
	//	})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	sendChatResponse, err := client.Chat().SendChatMessageAsUser(context.TODO(), accessToken, 123, nil, "Hello chat!")
	//	if err != nil {
	//		if apiError := kickerrors.IsAPIError(err); apiError != nil {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	SendChatMessageAsUser(ctx context.Context, accessToken string, broadcasterUserID int, replyToMessageID *string, message string) (*kickapitypes.SendChatResponse, error)

	// SendChatMessageAsBot sends a chat message as a bot in a broadcaster's channel.
	//
	// replyToMessageID is optional.
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
	//	sendChatResponse, err := client.Chat().SendChatMessageAsBot(context.TODO(), accessToken, nil, "Hello from bot!")
	//	if err != nil {
	//		if apiError := kickerrors.IsAPIError(err); apiError != nil {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	SendChatMessageAsBot(ctx context.Context, accessToken string, replyToMessageID *string, message string) (*kickapitypes.SendChatResponse, error)

	// DeleteChatMessage deletes the message with the corresponding message ID from the broadcasters channel.
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

	//	err := client.Chat().DeleteChatMessage(context.TODO(), accessToken, "message-id")
	//	if err != nil {
	//		if apiError := kickerrors.IsAPIError(err); apiError != nil {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	DeleteChatMessage(ctx context.Context, accessToken string, messageID string) error
}
