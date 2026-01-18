package kickcontracts

type APIClient interface {
	Category() Category
	Channel() Channel
	Chat() Chat
	EventsSubscription() EventsSubscription
	Livestream() Livestream
	Moderation() Moderation
	PublicKey() PublicKey
	User() User
}
