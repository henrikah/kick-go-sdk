package kick

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/henrikah/kick-go-sdk/v2/internal/endpoints"
	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
	"github.com/henrikah/kick-go-sdk/v2/kickcontracts"
	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
)

type kicksClient struct {
	client *apiClient
}

func newKicksService(client *apiClient) kickcontracts.Kicks {
	return &kicksClient{
		client: client,
	}
}

func (c *kicksClient) GetKicksLeaderboard(ctx context.Context, accessToken string, limit *int) (*kickapitypes.Kicks, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}

	if err := kickerrors.ValidateNotNil("limit", limit); err != nil {
		if err = kickerrors.ValidateBetween("limit", *limit, 1, 100); err != nil {
			return nil, err
		}
	}

	kicksLeaderboardURL, err := url.Parse(endpoints.ViewKicksLeaderboardURL())
	if err != nil {
		return nil, err
	}

	if limit != nil {
		queryParams := kicksLeaderboardURL.Query()
		queryParams.Add("top", strconv.Itoa(*limit))
		kicksLeaderboardURL.RawQuery = queryParams.Encode()
	}
	var kicksLeaderboardData kickapitypes.Kicks

	if err = c.client.requester.MakeJSONRequest(ctx, http.MethodGet, kicksLeaderboardURL.String(), nil, &accessToken, &kicksLeaderboardData); err != nil {
		return nil, err
	}

	return &kicksLeaderboardData, nil
}
