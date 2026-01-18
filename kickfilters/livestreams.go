package kickfilters

import (
	"net/url"
	"strconv"

	"github.com/henrikah/kick-go-sdk/enums/kicksortbyenum"
	"github.com/henrikah/kick-go-sdk/kickerrors"
)

// LivestreamsFilter is a fluent builder for constructing filters when searching livestreams.
type LivestreamsFilter interface {
	// WithBroadcasterUserIDs filters by one or more broadcaster user IDs.
	WithBroadcasterUserIDs(ids []int64) LivestreamsFilter

	// WithCategoryID filters by a specific category ID.
	WithCategoryID(id int64) LivestreamsFilter

	// WithLanguage filters by language code (e.g., "en").
	WithLanguage(language string) LivestreamsFilter

	// WithLimit limits the number of results (1–100).
	WithLimit(limit int) LivestreamsFilter

	// WithSortBy sorts results by the given field.
	WithSortBy(sort kicksortbyenum.SortBy) LivestreamsFilter

	// ToQueryString converts the builder into a query string for the API call.
	//
	// Returns an error if any filter validation fails (e.g., invalid IDs or limits).
	ToQueryString() (url.Values, error)
}

type livestreams struct {
	BroadcasterUserIDs []int64
	CategoryID         *int64
	Language           *string
	Limit              *int
	SortBy             *kicksortbyenum.SortBy
}

type livestreamsFilter struct {
	filter livestreams
}

// NewLivestreamsFilter creates a new builder for livestream filters.
//
// Example:
//
//	filters := kickfilters.NewLivestreamsFilter().
//	    WithCategoryID(42).
//	    WithLanguage("en").
//	    WithLimit(20)
//
//	query, err := filters.ToQueryString()
//	if err != nil {
//	    log.Printf("could not build query string: %v", err)
//	}
func NewLivestreamsFilter() LivestreamsFilter {
	return &livestreamsFilter{
		filter: livestreams{},
	}
}

func (f *livestreamsFilter) WithBroadcasterUserIDs(ids []int64) LivestreamsFilter {
	f.filter.BroadcasterUserIDs = ids
	return f
}

func (f *livestreamsFilter) WithCategoryID(id int64) LivestreamsFilter {
	f.filter.CategoryID = &id
	return f
}

func (f *livestreamsFilter) WithLanguage(language string) LivestreamsFilter {
	f.filter.Language = &language
	return f
}

func (f *livestreamsFilter) WithLimit(limit int) LivestreamsFilter {
	f.filter.Limit = &limit
	return f
}

func (f *livestreamsFilter) WithSortBy(sort kicksortbyenum.SortBy) LivestreamsFilter {
	f.filter.SortBy = &sort
	return f
}

func (f *livestreamsFilter) ToQueryString() (url.Values, error) {
	values := url.Values{}
	for _, id := range f.filter.BroadcasterUserIDs {
		if id < 1 {
			return nil, &kickerrors.ValidationError{
				Field:   "BroadcasterUserIDs",
				Message: "must be 1 or higher",
			}
		}
		values.Add("broadcaster_user_id", strconv.FormatInt(id, 10))
	}
	if f.filter.CategoryID != nil {
		if *f.filter.CategoryID < 1 {
			return nil, &kickerrors.ValidationError{
				Field:   "CategoryID",
				Message: "must be 1 or higher",
			}
		}
		values.Set("category_id", strconv.FormatInt(*f.filter.CategoryID, 10))
	}
	if f.filter.Language != nil {
		values.Set("language", *f.filter.Language)
	}
	if f.filter.Limit != nil {
		if *f.filter.Limit < 1 || *f.filter.Limit > 100 {
			return nil, &kickerrors.ValidationError{
				Field:   "Limit",
				Message: "must be between 1 and 100",
			}
		}
		values.Set("limit", strconv.Itoa(*f.filter.Limit))
	}
	if f.filter.SortBy != nil {
		values.Set("sort", string(*f.filter.SortBy))
	}

	return values, nil
}
