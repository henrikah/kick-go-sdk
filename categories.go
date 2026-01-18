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

type categoryClient struct {
	client *apiClient
}

func newCategoryService(client *apiClient) kickcontracts.Category {
	return &categoryClient{
		client: client,
	}
}
func (c *categoryClient) SearchCategories(ctx context.Context, accessToken string, filters kickfilters.CategoriesFilter) (*kickapitypes.GetCategoriesResponse, error) {
	if err := kickerrors.ValidateAccessToken(accessToken); err != nil {
		return nil, err
	}

	categoriesURL, err := url.Parse(endpoints.SearchCategoriesURL())
	if err != nil {
		return nil, err
	}

	if filters != nil {
		queryParams, filterErr := filters.ToQueryString()
		if filterErr != nil {
			return nil, filterErr
		}

		categoriesURL.RawQuery = queryParams.Encode()
	}

	var categoriesData kickapitypes.GetCategoriesResponse

	if err = c.client.requester.MakeJSONRequest(ctx, http.MethodGet, categoriesURL.String(), nil, &accessToken, &categoriesData); err != nil {
		return nil, err
	}

	return &categoriesData, nil
}
