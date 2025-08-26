package kickfilters

import (
	"net/url"
	"strconv"

	"github.com/henrikah/kick-go-sdk/enums/kicksortbyenum"
	"github.com/henrikah/kick-go-sdk/kickerrors"
)

// LivestreamFilterBuilder is a fluent builder for constructing filters when searching livestreams.
//
// Use the builder to safely configure filters and pass it directly into
// client methods such as livestream.SearchLivestreams.
type LivestreamFilterBuilder interface {
	// WithBroadcasterUserIDs filters by one or more broadcaster user IDs.
	WithBroadcasterUserIDs(ids []int64) LivestreamFilterBuilder

	// WithCategoryID filters by a specific category ID.
	WithCategoryID(id int64) LivestreamFilterBuilder

	// WithLanguage filters by language code (e.g., "en").
	WithLanguage(language string) LivestreamFilterBuilder

	// WithLimit limits the number of results (1–100).
	WithLimit(limit int) LivestreamFilterBuilder

	// WithSortBy sorts results by the given field.
	WithSortBy(sort kicksortbyenum.SortBy) LivestreamFilterBuilder

	// ToQueryString converts the builder into a query string for the API call.
	//
	// Returns an error if any filter validation fails (e.g., invalid IDs or limits).
	ToQueryString() (url.Values, error)
}

type livestreamFilter struct {
	BroadcasterUserIDs []int64
	CategoryID         *int64
	Language           *string
	Limit              *int
	SortBy             *kicksortbyenum.SortBy
}

type livestreamFilterBuilder struct {
	filter livestreamFilter
}

// NewLivestreamFilterBuilder creates a new builder for livestream filters.
//
// Example:
//
//	filters := kickfilters.NewLivestreamFilterBuilder().
//	    WithCategoryID(42).
//	    WithLanguage("en").
//	    WithLimit(20)
//
//	query, err := filters.ToQueryString()
//	if err != nil {
//	    log.Printf("could not build query string: %v", err)
//	    return nil, err
//	}
func NewLivestreamFilterBuilder() LivestreamFilterBuilder {
	return &livestreamFilterBuilder{
		filter: livestreamFilter{},
	}
}

func (f *livestreamFilterBuilder) WithBroadcasterUserIDs(ids []int64) LivestreamFilterBuilder {
	f.filter.BroadcasterUserIDs = ids
	return f
}

func (f *livestreamFilterBuilder) WithCategoryID(id int64) LivestreamFilterBuilder {
	f.filter.CategoryID = &id
	return f
}

func (f *livestreamFilterBuilder) WithLanguage(language string) LivestreamFilterBuilder {
	f.filter.Language = &language
	return f
}

func (f *livestreamFilterBuilder) WithLimit(limit int) LivestreamFilterBuilder {
	f.filter.Limit = &limit
	return f
}

func (f *livestreamFilterBuilder) WithSortBy(sort kicksortbyenum.SortBy) LivestreamFilterBuilder {
	f.filter.SortBy = &sort
	return f
}

func (f *livestreamFilterBuilder) ToQueryString() (url.Values, error) {
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
