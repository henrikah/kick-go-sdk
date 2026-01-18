package kickcontracts

import (
	"context"

	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickfilters"
)

// Category handles operations related to Kick categories.
//
// All examples use context.TODO() as a placeholder. Replace with a proper
// context (with timeout/cancel) in production code.
type Category interface {
	// SearchCategories gets all the categories associated with the applied filter.
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
	//	filters := kickfilters.NewCategoriesFilter().
	//	    WithCategoryIDs([]int64{42}).
	//	    WithTags([]string{"GoLang"}).
	//	    WithLimit(20)
	//
	//	categories, err := client.Category().SearchCategories(context.TODO(), accessToken, filters)
	//	if err != nil {
	//		var apiErr *kickerrors.APIError
	//		if errors.As(err, &apiErr) {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	SearchCategories(ctx context.Context, accessToken string, filters kickfilters.CategoriesFilter) (*kickapitypes.GetCategoriesResponse, error)
}
