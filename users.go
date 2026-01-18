package kick

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/henrikah/kick-go-sdk/internal/endpoints"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickcontracts"
	"github.com/henrikah/kick-go-sdk/kickerrors"
)

type userClient struct {
	client *apiClient
}

func newUserService(client *apiClient) kickcontracts.User {
	return &userClient{
		client: client,
	}
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

	if err = c.client.requester.MakeJSONRequest(ctx, http.MethodGet, usersURL.String(), nil, &accessToken, &usersData); err != nil {
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
