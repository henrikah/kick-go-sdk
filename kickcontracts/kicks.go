package kickcontracts

import (
	"context"

	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
)

// Kicks handles displaying kicks.
//
// All examples use context.TODO() as a placeholder. Replace with a proper
// context (with timeout/cancel) in production code.
type Kicks interface {
	// GetKicksLeaderboard gets lifetime, monthly, and weekly leaderboards of users that have sent kicks to the channel the access token is from
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
	//	kicksLeaderboardResponse, err := client.Kicks().GetKicksLeaderboard(context.TODO(), accessToken, nil)
	//	if err != nil {
	//		if apiErr := kickerrors.IsAPIError(err); apiErr != nil {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	GetKicksLeaderboard(ctx context.Context, accessToken string, limit *int) (*kickapitypes.Kicks, error)
}
