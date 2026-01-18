package kickcontracts

import (
	"context"

	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickfilters"
)

// Livestream handles operations related to Kick livestreams.
//
// All examples use context.TODO() as a placeholder. Replace with a proper
// context (with timeout/cancel) in production code.
type Livestream interface {
	// SearchLivestreams retrieves livestreams matching the given filters.
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
	//	filters := kickfilters.NewLivestreamsFilter().
	//	    WithCategoryID(42).
	//	    WithLanguage("en").
	//	    WithLimit(20)
	//
	//	livestreams, err := client.Livestream().SearchLivestreams(context.TODO(), accessToken, filters)
	//	if err != nil {
	//		var apiErr *kickerrors.APIError
	//		if errors.As(err, &apiErr) {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	SearchLivestreams(ctx context.Context, accessToken string, filters kickfilters.LivestreamsFilter) (*kickapitypes.LivestreamResponse, error)

	// GetCurrentUserLivestream retrieves the livestream for the user's access token. If the user is not live it will return and empty slice.
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
	//	livestream, err := client.Livestream().GetCurrentUserLivestream(context.TODO(), accessToken)
	//	if err != nil {
	//		var apiErr *kickerrors.APIError
	//		if errors.As(err, &apiErr) {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	GetCurrentUserLivestream(ctx context.Context, accessToken string) (*kickapitypes.LivestreamResponse, error)
}
