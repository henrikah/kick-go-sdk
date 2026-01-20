package kickfilters

import (
	"net/url"
	"strconv"

	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
)

// CategoriesFilter is a fluent builder for constructing filters when searching categories.
type CategoriesFilter interface {
	// WithTags filters by one or more tags.
	WithTags(tags []string) CategoriesFilter

	// WithCategoryIDs filters by category IDs.
	WithCategoryIDs(ids []int64) CategoriesFilter

	// WithNames filters by names.
	WithNames(names []string) CategoriesFilter

	// WithLimit limits the number of results.
	WithLimit(limit int) CategoriesFilter

	// WithCursor applies the cursor received from a previous request for pagination.
	WithCursor(cursor string) CategoriesFilter

	// ToQueryString converts the builder into a query string for the API call.
	//
	// Returns an error if any filter validation fails.
	ToQueryString() (url.Values, error)
}

type categories struct {
	CategoryIDs []int64
	Tags        []string
	Names       []string
	Limit       *int
	Cursor      *string
}

type categoriesFilter struct {
	filter categories
}

// NewCategoriesFilter creates a new builder for categories filters.
//
// Example:
//
//	filters := kickfilters.NewCategoriesFilter().
//	    WithCategoryIDs([]int64{42}).
//	    WithTags([]string{"GoLang"})).
//	    WithLimit(20)
//
//	query, err := filters.ToQueryString()
//	if err != nil {
//	    log.Printf("could not build query string: %v", err)
//	}
func NewCategoriesFilter() CategoriesFilter {
	return &categoriesFilter{
		filter: categories{},
	}
}

func (f *categoriesFilter) WithTags(tags []string) CategoriesFilter {
	f.filter.Tags = tags
	return f
}

func (f *categoriesFilter) WithCategoryIDs(ids []int64) CategoriesFilter {
	f.filter.CategoryIDs = ids
	return f
}

func (f *categoriesFilter) WithNames(names []string) CategoriesFilter {
	f.filter.Names = names
	return f
}

func (f *categoriesFilter) WithLimit(limit int) CategoriesFilter {
	f.filter.Limit = &limit
	return f
}

func (f *categoriesFilter) WithCursor(cursor string) CategoriesFilter {
	f.filter.Cursor = &cursor
	return f
}

func (f *categoriesFilter) ToQueryString() (url.Values, error) {
	values := url.Values{}
	for _, id := range f.filter.CategoryIDs {
		if id < 1 {
			return nil, &kickerrors.ValidationError{
				Field:   "CategoryIDs",
				Message: "must be 1 or higher",
			}
		}
		values.Add("id", strconv.FormatInt(id, 10))
	}
	for _, name := range f.filter.Tags {
		values.Add("tag", name)
	}
	for _, name := range f.filter.Names {
		values.Add("name", name)
	}
	if f.filter.Limit != nil {
		if *f.filter.Limit < 1 {
			return nil, &kickerrors.ValidationError{
				Field:   "Limit",
				Message: "must be 1 or more",
			}
		}
		values.Set("limit", strconv.Itoa(*f.filter.Limit))
	}
	if f.filter.Cursor != nil {
		values.Set("cursor", *f.filter.Cursor)
	}

	return values, nil
}
