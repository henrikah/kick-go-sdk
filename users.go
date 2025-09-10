package kick

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/henrikah/kick-go-sdk/internal/endpoints"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickerrors"
)

// user handles operations related to users.
//
// All examples use context.TODO() as a placeholder. Replace with a proper
// context (with timeout/cancel) in production code.
type user interface {
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

type userClient struct {
	client *apiClient
}

func newUserClient(client *apiClient) user {
	return &userClient{
		client: client,
	}
}
func (c *userClient) TokenIntrospect(ctx context.Context, accessToken string) (*kickapitypes.TokenIntrospect, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}

	var tokenIntrospectData kickapitypes.TokenIntrospect

	if err := c.client.makeJSONRequest(ctx, http.MethodPost, endpoints.ViewTokenIntrospectURL(), nil, &accessToken, &tokenIntrospectData); err != nil {
		return nil, err
	}

	return &tokenIntrospectData, nil
}
func (c *userClient) GetUsersByID(ctx context.Context, accessToken string, userIDs []int64) (*kickapitypes.Users, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}

	usersURL, err := url.Parse(endpoints.ViewUsersDetailsURL())
	if err != nil {
		return nil, err
	}

	queryParams := usersURL.Query()
	for _, id := range userIDs {
		queryParams.Add("id", strconv.FormatInt(id, 10))
	}
	usersURL.RawQuery = queryParams.Encode()

	var usersData kickapitypes.Users

	if err = c.client.makeJSONRequest(ctx, http.MethodGet, usersURL.String(), nil, &accessToken, &usersData); err != nil {
		return nil, err
	}

	return &usersData, nil
}

func (c *userClient) GetUserByID(ctx context.Context, accessToken string, userID int64) (*kickapitypes.Users, error) {
	return c.GetUsersByID(ctx, accessToken, []int64{userID})
}

func (c *userClient) GetCurrentUser(ctx context.Context, accessToken string) (*kickapitypes.Users, error) {
	return c.GetUsersByID(ctx, accessToken, nil)
}
