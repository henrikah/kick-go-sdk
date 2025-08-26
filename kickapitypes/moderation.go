package kickapitypes

type ModerationResponse struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type ModerationRequest struct {
	BroadcasterUserID int     `json:"broadcaster_user_id"`
	Duration          *int    `json:"duration,omitempty"`
	Reason            *string `json:"reason,omitempty"`
	UserID            int     `json:"user_id"`
}

type UnBanRequest struct {
	BroadcasterUserID int `json:"broadcaster_user_id"`
	UserID            int `json:"user_id"`
}
