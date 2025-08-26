package kickwebhooktypes

/** Parent structs **/

type ChatMessageSent struct {
	MessageID   string    `json:"message_id"`
	RepliesTo   RepliesTo `json:"replies_to"`
	Broadcaster User      `json:"broadcaster"`
	Sender      User      `json:"sender"`
	Content     string    `json:"content"`
	Emotes      []Emote   `json:"emotes"`
	CreatedAt   string    `json:"created_at"`
}

type ChannelFollowed struct {
	Broadcaster User `json:"broadcaster"`
	Follower    User `json:"follower"`
}

type ChannelSubscriptionRenewal struct {
	Broadcaster User   `json:"broadcaster"`
	Subscriber  User   `json:"subscriber"`
	Duration    int    `json:"duration"`
	CreatedAt   string `json:"created_at"`
	ExpiresAt   string `json:"expires_at"`
}

type ChannelSubscriptionGifts struct {
	Broadcaster User   `json:"broadcaster"`
	Gifter      User   `json:"gifter"`
	Giftees     []User `json:"giftees"`
	CreatedAt   string `json:"created_at"`
	ExpiresAt   string `json:"expires_at"`
}

type ChannelSubscriptionNew struct {
	Broadcaster User   `json:"broadcaster"`
	Subscriber  User   `json:"subscriber"`
	Duration    int    `json:"duration"`
	CreatedAt   string `json:"created_at"`
	ExpiresAt   string `json:"expires_at"`
}

type LivestreamStatusUpdated struct {
	Broadcaster User    `json:"broadcaster"`
	IsLive      bool    `json:"is_live"`
	Title       string  `json:"title"`
	StartedAt   string  `json:"started_at"`
	EndedAt     *string `json:"ended_at"`
}

type LivestreamMetadataUpdated struct {
	Broadcaster User               `json:"broadcaster"`
	Metadata    LivestreamMetadata `json:"metadata"`
}

type ModerationBanned struct {
	Broadcaster User                     `json:"broadcaster"`
	Moderator   User                     `json:"moderator"`
	BannedUser  User                     `json:"banned_user"`
	Metadata    ModerationBannedMetadata `json:"metadata"`
}

/** Body structs **/

type RepliesTo struct {
	MessageID string `json:"message_id"`
	Content   string `json:"content"`
	Sender    User   `json:"sender"`
}

type User struct {
	IsAnonymous     bool      `json:"is_anonymous"`
	UserID          int       `json:"user_id"`
	Username        string    `json:"username"`
	IsVerified      bool      `json:"is_verified"`
	ProfilePictures string    `json:"profile_picture"`
	ChannelSlug     string    `json:"channel_slug"`
	Identity        *Identity `json:"identity"`
}

type Identity struct {
	UsernameColor string  `json:"username_color"`
	Badges        []Badge `json:"badges"`
}

type Badge struct {
	Text  string `json:"text"`
	Type  string `json:"type"`
	Count *int   `json:"count,omitempty"`
}

type Emote struct {
	EmoteID   string     `json:"emote_id"`
	Positions []Position `json:"positions"`
}

type Position struct {
	Start int `json:"s"`
	End   int `json:"e"`
}

type LivestreamMetadata struct {
	Title            string   `json:"title"`
	Language         string   `json:"language"`
	HasMatureContent bool     `json:"has_mature_content"`
	Category         Category `json:"category"`
}

type Category struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
}

type ModerationBannedMetadata struct {
	Reason    string `json:"reason"`
	CreatedAt string `json:"created_at"`
	ExpiresAt string `json:"expires_at"`
}
