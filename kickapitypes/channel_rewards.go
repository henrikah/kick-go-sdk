package kickapitypes

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
