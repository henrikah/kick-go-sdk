package kickcontracts

import (
	"context"

	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
)

// User handles operations related to users.
//
// All examples use context.TODO() as a placeholder. Replace with a proper
// context (with timeout/cancel) in production code.
type User interface {

	// GetUsersByID retrieves multiple users by their IDs.
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
	//	users, err := client.User().GetUsersByID(context.TODO(), accessToken, []int64{123, 456})
	//	if err != nil {
	//		if apiError := kickerrors.IsAPIError(err); apiError != nil {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	GetUsersByID(ctx context.Context, accessToken string, userIDs []int64) (*kickapitypes.Users, error)

	// GetUserByID retrieves a single user by ID.
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
	//	user, err := client.User().GetUserByID(context.TODO(), accessToken, 123)
	//	if err != nil {
	//		if apiError := kickerrors.IsAPIError(err); apiError != nil {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	GetUserByID(ctx context.Context, accessToken string, userID int64) (*kickapitypes.Users, error)

	// GetCurrentUser retrieves the user associated with the given access token.
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
	//	currentUser, err := client.User().GetCurrentUser(context.TODO(), accessToken)
	//	if err != nil {
	//		if apiError := kickerrors.IsAPIError(err); apiError != nil {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	GetCurrentUser(ctx context.Context, accessToken string) (*kickapitypes.Users, error)
}
