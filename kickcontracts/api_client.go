package kickcontracts

type APIClient interface {
	Category() Category
	Channel() Channel
	Chat() Chat
	EventsSubscription() EventsSubscription
	Livestream() Livestream
	Moderation() Moderation
	OAuth() OAuth
	PublicKey() PublicKey
	User() User
}
