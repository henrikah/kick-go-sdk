package kickapitypes

import "github.com/henrikah/kick-go-sdk/enums/kickscopes"

type TokenIntrospectData struct {
	Active    bool              `json:"active,omitempty"`
	ClientID  string            `json:"client_id,omitempty"`
	Expires   int               `json:"exp,omitempty"`
	Scopes    kickscopes.Scopes `json:"scope,omitempty"`
	TokenType string            `json:"token_type,omitempty"`
}

type TokenIntrospect struct {
	Data    TokenIntrospectData `json:"data,omitempty"`
	Message string              `json:"message,omitempty"`
}

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
