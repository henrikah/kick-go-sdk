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

type kicksClient struct {
	client *apiClient
}

func newKicksClient(client *apiClient) kickcontracts.Kicks {
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

	if err = c.client.makeJSONRequest(ctx, http.MethodGet, kicksLeaderboardURL.String(), nil, &accessToken, &kicksLeaderboardData); err != nil {
		return nil, err
	}

	return &kicksLeaderboardData, nil
}
