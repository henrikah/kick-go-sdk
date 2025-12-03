// Package kickwebhookenum contains enums for the webhooks.
package kickwebhookenum

type WebhookType string

const (
	ChatMessageSent                WebhookType = "chat.message.sent"
	ChannelFollowed                WebhookType = "channel.followed"
	ChannelSubscriptionRenewal     WebhookType = "channel.subscription.renewal"
	ChannelSubscriptionGifts       WebhookType = "channel.subscription.gifts"
	ChannelSubscriptionNew         WebhookType = "channel.subscription.new"
	LivestreamStatusUpdated        WebhookType = "livestream.status.updated"
	LivestreamMetadataUpdated      WebhookType = "livestream.metadata.updated"
	ModerationBanned               WebhookType = "moderation.banned"
	KicksGifted                    WebhookType = "kicks.gifted"
	ChannelRewardRedemptionUpdated WebhookType = "channel.reward.redemption.updated"
)
