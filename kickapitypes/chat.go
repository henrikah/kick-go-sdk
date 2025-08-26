package kickapitypes

type SendChatResponse struct {
	Data    SendChatResponseData `json:"data"`
	Message string               `json:"message"`
}

type SendChatResponseData struct {
	IsSent    bool   `json:"is_sent"`
	MessageID string `json:"message_id"`
}

type SendChatRequest struct {
	BroadcasterUserID *int    `json:"broadcaster_user_id,omitempty"`
	Content           string  `json:"content"`
	ReplyToMessageID  *string `json:"reply_to_message_id,omitempty"`
	Type              string  `json:"type"`
}
