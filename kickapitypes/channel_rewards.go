package kickapitypes

import "github.com/henrikah/kick-go-sdk/v2/enums/kickchannelrewardstatus"

type ChannelRewards struct {
	Data    []ChannelRewardData `json:"data"`
	Message string              `json:"message"`
}

type ChannelReward struct {
	Data    ChannelRewardData `json:"data"`
	Message string            `json:"message"`
}

type ChannelRewardData struct {
	ID                                string `json:"id"`
	BackgroundColor                   string `json:"background_color"`
	Cost                              int    `json:"cost"`
	Description                       string `json:"description"`
	IsEnabled                         bool   `json:"is_enabled"`
	IsPaused                          bool   `json:"is_paused"`
	IsUserInputRequired               bool   `json:"is_user_input_required"`
	ShouldRedemptionsSkipRequestQueue bool   `json:"should_redemptions_skip_request_queue"`
	Title                             string `json:"title"`
}

type CreateChannelReward struct {
	BackgroundColor                   *string `json:"background_color,omitempty"`
	Cost                              int     `json:"cost"`
	Description                       *string `json:"description,omitempty"`
	IsEnabled                         *bool   `json:"is_enabled,omitempty"`
	IsUserInputRequired               *bool   `json:"is_user_input_required,omitempty"`
	ShouldRedemptionsSkipRequestQueue *bool   `json:"should_redemptions_skip_request_queue,omitempty"`
	Title                             string  `json:"title"`
}

type UpdateChannelReward struct {
	BackgroundColor                   *string `json:"background_color,omitempty"`
	Cost                              *int    `json:"cost,omitempty"`
	Description                       *string `json:"description,omitempty"`
	IsEnabled                         *bool   `json:"is_enabled,omitempty"`
	IsPaused                          *bool   `json:"is_paused,omitempty"`
	IsUserInputRequired               *bool   `json:"is_user_input_required,omitempty"`
	ShouldRedemptionsSkipRequestQueue *bool   `json:"should_redemptions_skip_request_queue,omitempty"`
	Title                             *string `json:"title,omitempty"`
}

type ChannelRewardRedemptions struct {
	Data       []ChannelRewardRedemptionData `json:"data"`
	Message    string                        `json:"message"`
	Pagination Pagination                    `json:"pagination"`
}

type ChannelRewardRedemptionData struct {
	Redemptions []RedemptionData `json:"redemptions"`
	Reward      RewardData       `json:"reward"`
}

type Redeemer struct {
	BroadcasterUserID uint64 `json:"broadcaster_user_id"`
}

type RedemptionData struct {
	ID         string                                      `json:"id"`
	RedeemedAt string                                      `json:"redeemed_at"`
	Redeemer   Redeemer                                    `json:"redeemer"`
	Status     kickchannelrewardstatus.ChannelRewardStatus `json:"status"`
	UserInput  string                                      `json:"user_input"`
}

type RewardData struct {
	CanManage   *bool   `json:"can_manage,omitempty"`
	Cost        *uint64 `json:"cost,omitempty"`
	Description *string `json:"description,omitempty"`
	ID          string  `json:"id"`
	IsDeleted   *bool   `json:"is_deleted,omitempty"`
	Title       string  `json:"title"`
}

type RedemptionsIDs struct {
	IDs []string `json:"ids"`
}

type RedemptionDecision struct {
	Data    []RedemptionDecisionData `json:"data"`
	Message string                   `json:"message"`
}

type RedemptionDecisionData struct {
	ID     string `json:"id"`
	Reason string `json:"reason"`
}
