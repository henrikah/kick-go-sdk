package kickcontracts

import (
	"context"

	"github.com/henrikah/kick-go-sdk/kickapitypes"
)

// Category handles operations related to Kick categories.
//
// All examples use context.TODO() as a placeholder. Replace with a proper
// context (with timeout/cancel) in production code.
type Category interface {
	// SearchCategories gets all the categories associated with the searchQuery.
	//
	// Example:
	//
	//	categories, err := client.Category().SearchCategories(context.TODO(), accessToken, searchQuery, pageNumber)
	//	if err != nil {
	//	    log.Printf("could not search categories: %v", err)
	//	    return nil, err
	//	}
	SearchCategories(ctx context.Context, accessToken string, searchQuery string, pageNumber int) (*kickapitypes.GetCategoriesResponse, error)

	// GetCategoryByCategoryID gets detailed data for the category associated with the categoryID.
	//
	// Example:
	//
	//	category, err := client.Category().GetCategoryByCategoryID(context.TODO(), accessToken, categoryID)
	//	if err != nil {
	//	    log.Printf("could not get category by ID: %v", err)
	//	    return nil, err
	//	}
	GetCategoryByCategoryID(ctx context.Context, accessToken string, categoryID int) (*kickapitypes.GetCategoryResponse, error)
}
