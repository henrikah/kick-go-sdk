package kickapitypes

type Channels struct {
	Data    []ChannelData `json:"data"`
	Message string        `json:"message"`
}

type ChannelData struct {
	BannerPicture      string        `json:"banner_picture"`
	BroadcasterUserID  int           `json:"broadcaster_user_id"`
	Category           Category      `json:"category"`
	ChannelDescription string        `json:"channel_description"`
	Slug               string        `json:"slug"`
	Stream             ChannelStream `json:"stream"`
	StreamTitle        string        `json:"stream_title"`
}

type ChannelStream struct {
	CustomTags  []string `json:"custom_tags"`
	IsLive      bool     `json:"is_live"`
	IsMature    bool     `json:"is_mature"`
	Key         string   `json:"key"`
	Language    string   `json:"language"`
	StartTime   string   `json:"start_time"`
	Thumbnail   string   `json:"thumbnail"`
	URL         string   `json:"url"`
	ViewerCount int      `json:"viewer_count"`
}

type UpdateChannelRequest struct {
	CategoryID  int64    `json:"category_id,omitempty"`
	CustomTags  []string `json:"custom_tags,omitempty"`
	StreamTitle string   `json:"stream_title,omitempty"`
}
