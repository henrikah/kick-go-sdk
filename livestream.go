package kick

import (
	"context"
	"net/http"
	"net/url"

	"github.com/henrikah/kick-go-sdk/internal/endpoints"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickcontracts"
	"github.com/henrikah/kick-go-sdk/kickerrors"
	"github.com/henrikah/kick-go-sdk/kickfilters"
)

type livestreamClient struct {
	client *apiClient
}

func newLivestreamService(client *apiClient) kickcontracts.Livestream {
	return &livestreamClient{
		client: client,
	}
}

func (c *livestreamClient) SearchLivestreams(ctx context.Context, accessToken string, filters kickfilters.LivestreamsFilter) (*kickapitypes.LivestreamResponse, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}

	livestreamURL, err := url.Parse(endpoints.ViewLivestreamsDetailsURL())
	if err != nil {
		return nil, err
	}

	if filters != nil {
		queryParams, filterErr := filters.ToQueryString()
		if filterErr != nil {
			return nil, filterErr
		}

		livestreamURL.RawQuery = queryParams.Encode()
	}

	var livestreamResponse kickapitypes.LivestreamResponse

	if err = c.client.requester.MakeJSONRequest(ctx, http.MethodGet, livestreamURL.String(), nil, &accessToken, &livestreamResponse); err != nil {
		return nil, err
	}

	return &livestreamResponse, nil
}
func (c *livestreamClient) GetCurrentUserLivestream(ctx context.Context, accessToken string) (*kickapitypes.LivestreamResponse, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}

	var livestreamResponse kickapitypes.LivestreamResponse

	if err := c.client.requester.MakeJSONRequest(ctx, http.MethodGet, endpoints.ViewCurrentUserLivestreamDetailsURL(), nil, &accessToken, &livestreamResponse); err != nil {
		return nil, err
	}

	return &livestreamResponse, nil
}
