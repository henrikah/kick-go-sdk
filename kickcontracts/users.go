package kickcontracts

import (
	"context"

	"github.com/henrikah/kick-go-sdk/kickapitypes"
)

// User handles operations related to users.
//
// All examples use context.TODO() as a placeholder. Replace with a proper
// context (with timeout/cancel) in production code.
type User interface {
	// TokenIntrospect validates an access token and returns metadata such as
	// validity, expiration, and associated user/client.
	//
	// Example:
	//
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{...})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	introspection, err := client.User().TokenIntrospect(context.TODO(), accessToken)
	//	if err != nil {
	//  	log.Printf("could not introspect token: %v", err)
	//  	return nil, err
	//	}
	TokenIntrospect(ctx context.Context, accessToken string) (*kickapitypes.TokenIntrospect, error)

	// GetUsersByID retrieves multiple users by their IDs.
	//
	// Example:
	//
	//	users, err := client.User().GetUsersByID(context.TODO(), accessToken, []int64{123, 456})
	//	if err != nil {
	//  	log.Printf("could not get users by id: %v", err)
	//  	return nil, err
	//	}
	GetUsersByID(ctx context.Context, accessToken string, userIDs []int64) (*kickapitypes.Users, error)

	// GetUserByID retrieves a single user by ID.
	//
	// Example:
	//
	//	user, err := client.User().GetUserByID(context.TODO(), accessToken, 123)
	//	if err != nil {
	//  	log.Printf("could not get user by id: %v", err)
	//  	return nil, err
	//	}
	GetUserByID(ctx context.Context, accessToken string, userID int64) (*kickapitypes.Users, error)

	// GetCurrentUser retrieves the user associated with the given access token.
	//
	// Example:
	//
	//	currentUser, err := client.User().GetCurrentUser(context.TODO(), accessToken)
	//	if err != nil {
	//  	log.Printf("could not get the current user: %v", err)
	//  	return nil, err
	//	}
	GetCurrentUser(ctx context.Context, accessToken string) (*kickapitypes.Users, error)
}
