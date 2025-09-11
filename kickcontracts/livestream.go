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
