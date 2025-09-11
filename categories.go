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

type categoryClient struct {
	client *apiClient
}

func newCategoryClient(client *apiClient) kickcontracts.Category {
	return &categoryClient{
		client: client,
	}
}
func (c *categoryClient) SearchCategories(ctx context.Context, accessToken string, searchQuery string, pageNumber int) (*kickapitypes.GetCategoriesResponse, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}
	if err := kickerrors.ValidatePageNumber(pageNumber); err != nil {
		return nil, err
	}

	categoriesURL, err := url.Parse(endpoints.SearchCategoriesURL())
	if err != nil {
		return nil, err
	}

	queryParams := categoriesURL.Query()
	queryParams.Set("q", searchQuery)
	queryParams.Set("page", strconv.Itoa(pageNumber))

	categoriesURL.RawQuery = queryParams.Encode()

	var categoriesData kickapitypes.GetCategoriesResponse

	if err = c.client.makeJSONRequest(ctx, http.MethodGet, categoriesURL.String(), nil, &accessToken, &categoriesData); err != nil {
		return nil, err
	}

	return &categoriesData, nil
}
func (c *categoryClient) GetCategoryByCategoryID(ctx context.Context, accessToken string, categoryID int) (*kickapitypes.GetCategoryResponse, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}
	if err := kickerrors.ValidateCategoryID(categoryID); err != nil {
		return nil, err
	}

	categoriesURL, err := url.Parse(endpoints.ViewCategoryDetailsURL(categoryID))
	if err != nil {
		return nil, err
	}

	var categoryData kickapitypes.GetCategoryResponse

	if err := c.client.makeJSONRequest(ctx, http.MethodGet, categoriesURL.String(), nil, &accessToken, &categoryData); err != nil {
		return nil, err
	}

	return &categoryData, nil
}
