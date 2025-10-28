package kickcontracts

import (
	"context"

	"github.com/henrikah/kick-go-sdk/kickapitypes"
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
	//	kicksLeaderboardResponse, err := client.Kicks().GetKicksLeaderboard(context.TODO(), accessToken, nil)
	//	if err != nil {
	//	    log.Printf("could not get kicks leaderboards: %v", err)
	//	    return nil, err
	//	}
	GetKicksLeaderboard(ctx context.Context, accessToken string, limit *int) (*kickapitypes.Kicks, error)
}
