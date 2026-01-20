package kickfilters

import (
	"net/url"

	"github.com/henrikah/kick-go-sdk/v2/enums/kickchannelrewardstatus"
)

// RewardRedemptionsFilter is a fluent builder for constructing filters when searching reward redemptions.
type RewardRedemptionsFilter interface {
	// WithRewardID gets the redemption by Reward ID.
	WithRewardID(rewardID string) RewardRedemptionsFilter

	// WithStatus filters by a specific reward status.
	WithStatus(status kickchannelrewardstatus.ChannelRewardStatus) RewardRedemptionsFilter

	// WithRewardIDs selectes multiple rewards by their ID.
	WithRewardIDs(rewardIDs []string) RewardRedemptionsFilter

	// WithCursor applies the cursor received from a previous request for pagination.
	WithCursor(cursor string) RewardRedemptionsFilter

	// ToQueryString converts the builder into a query string for the API call.
	//
	// Returns an error if any filter validation fails (e.g., invalid IDs or limits).
	ToQueryString() (url.Values, error)
}

type rewardRedemptions struct {
	RewardID  *string
	Status    *kickchannelrewardstatus.ChannelRewardStatus
	RewardIDs []string
	Cursor    *string
}

type rewardRedemptionsFilter struct {
	filter rewardRedemptions
}

// NewRewardRedemptionsFilter creates a new builder for livestream filters.
//
// Example:
//
//	filters := kickfilters.NewRewardRedemptionsFilter().
//	    WithStatus(kickchannelrewardstatus.Pending)
//
//	query, err := filters.ToQueryString()
//	if err != nil {
//	    log.Printf("could not build query string: %v", err)
//	}
func NewRewardRedemptionsFilter() RewardRedemptionsFilter {
	return &rewardRedemptionsFilter{
		filter: rewardRedemptions{},
	}
}

func (f *rewardRedemptionsFilter) WithRewardID(rewardID string) RewardRedemptionsFilter {
	f.filter.RewardID = &rewardID
	return f
}

func (f *rewardRedemptionsFilter) WithStatus(status kickchannelrewardstatus.ChannelRewardStatus) RewardRedemptionsFilter {
	f.filter.Status = &status
	return f
}

func (f *rewardRedemptionsFilter) WithRewardIDs(rewardIDs []string) RewardRedemptionsFilter {
	f.filter.RewardIDs = rewardIDs
	return f
}
func (f *rewardRedemptionsFilter) WithCursor(cursor string) RewardRedemptionsFilter {
	f.filter.Cursor = &cursor
	return f
}

func (f *rewardRedemptionsFilter) ToQueryString() (url.Values, error) {
	values := url.Values{}
	if f.filter.RewardID != nil {
		values.Set("reward_id", *f.filter.RewardID)
	}

	if f.filter.Status != nil {
		values.Set("status", string(*f.filter.Status))
	}

	for _, id := range f.filter.RewardIDs {
		values.Add("id", id)
	}

	if f.filter.Cursor != nil {
		values.Set("cursor", *f.filter.Cursor)
	}

	return values, nil
}
