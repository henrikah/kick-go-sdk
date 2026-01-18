package kickapitypes

type UserData struct {
	Email          string `json:"email,omitempty"`
	Name           string `json:"name,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
	UserID         int    `json:"user_id,omitempty"`
}

type Users struct {
	Data    []UserData `json:"data,omitempty"`
	Message string     `json:"message,omitempty"`
}
