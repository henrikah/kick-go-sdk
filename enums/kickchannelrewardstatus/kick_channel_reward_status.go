// Package kickchannelrewardstatus contains enums for the channel reward status.
package kickchannelrewardstatus

type ChannelRewardStatus string

const (
	Accepted ChannelRewardStatus = "accepted"
	Pending  ChannelRewardStatus = "pending"
	Rejected ChannelRewardStatus = "rejected"
)
