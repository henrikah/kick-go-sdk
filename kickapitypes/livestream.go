package kickapitypes

type LivestreamResponse struct {
	Data    []LivestreamResponseData `json:"data"`
	Message string                   `json:"message"`
}

type LivestreamResponseData struct {
	BroadcasterUserID int      `json:"broadcaster_user_id"`
	Category          Category `json:"category"`
	ChannelID         int      `json:"channel_id"`
	CustomTags        []string `json:"custom_tags"`
	HasMatureContent  bool     `json:"has_mature_content"`
	Language          string   `json:"language"`
	Slug              string   `json:"slug"`
	StartedAt         string   `json:"started_at"`
	StreamTitle       string   `json:"stream_title"`
	Thumbnail         string   `json:"thumbnail"`
	ViewerCount       int      `json:"viewer_count"`
}
