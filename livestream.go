package kick

import (
	"context"
	"net/http"
	"net/url"

	"github.com/henrikah/kick-go-sdk/internal/endpoints"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickerrors"
	"github.com/henrikah/kick-go-sdk/kickfilters"
)

// livestream handles operations related to Kick livestreams.
//
// All examples use context.TODO() as a placeholder. Replace with a proper
// context (with timeout/cancel) in production code.
type livestream interface {
	// SearchLivestreams retrieves livestreams matching the given filters.
	//
	// Example:
	//
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{...})
	//	if err != nil {
	//	    log.Printf("could not create APIClient: %v", err)
	//	    return nil, err
	//	}
	//
	//	filters := kickfilters.NewLivestreamFilterBuilder().
	//	    WithCategoryID(42).
	//	    WithLanguage("en").
	//	    WithLimit(20)
	//
	//	livestreams, err := client.Livestream().SearchLivestreams(context.TODO(), accessToken, filters)
	//	if err != nil {
	//	    log.Printf("could not search livestreams: %v", err)
	//	    return nil, err
	//	}
	SearchLivestreams(ctx context.Context, accessToken string, filters kickfilters.LivestreamFilterBuilder) (*kickapitypes.LivestreamResponse, error)

	// GetCurrentUserLivestream retrieves the livestream for the user's access token. If the user is not live it will return and empty slice.
	//
	// Example:
	//
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{...})
	//	if err != nil {
	//	    log.Printf("could not create APIClient: %v", err)
	//	    return nil, err
	//	}
	//
	//	livestream, err := client.Livestream().GetCurrentUserLivestream(context.TODO(), accessToken)
	//	if err != nil {
	//	    log.Printf("could not current users livestream: %v", err)
	//	    return nil, err
	//	}
	GetCurrentUserLivestream(ctx context.Context, accessToken string) (*kickapitypes.LivestreamResponse, error)
}

type livestreamClient struct {
	client *apiClient
}

func newLivestreamClient(client *apiClient) livestream {
	return &livestreamClient{
		client: client,
	}
}

func (c *livestreamClient) SearchLivestreams(ctx context.Context, accessToken string, filters kickfilters.LivestreamFilterBuilder) (*kickapitypes.LivestreamResponse, error) {
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

	if err = c.client.makeJSONRequest(ctx, http.MethodGet, livestreamURL.String(), nil, &accessToken, &livestreamResponse); err != nil {
		return nil, err
	}

	return &livestreamResponse, nil
}
func (c *livestreamClient) GetCurrentUserLivestream(ctx context.Context, accessToken string) (*kickapitypes.LivestreamResponse, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}

	var livestreamResponse kickapitypes.LivestreamResponse

	if err := c.client.makeJSONRequest(ctx, http.MethodGet, endpoints.ViewCurrentUserLivestreamDetailsURL(), nil, &accessToken, &livestreamResponse); err != nil {
		return nil, err
	}

	return &livestreamResponse, nil
}
