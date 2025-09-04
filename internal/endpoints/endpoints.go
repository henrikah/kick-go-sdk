// Package endpoints contains all the endpoint constants and the builders for the final end points
package endpoints

import (
	"strconv"

	"github.com/henrikah/kick-go-sdk/internal/helpers"
)

type Endpoint string

const (
	/** Hostnames **/
	idHostname  = "https://id.kick.com"
	apiHostname = "https://api.kick.com"

	// AuthorizationPath is where the user can log in and approve the application's access request.
	userAuthorizationPath = "oauth/authorize"

	/** API Paths **/
	codeExchangePath                     = "oauth/token"
	generateAppAccessTokenPath           = "oauth/token"
	revokeTokenPath                      = "oauth/revoke"
	searchCategoriesPath                 = "public/v1/categories"
	viewCategoryDetailsPath              = "public/v1/categories"
	viewTokenIntrospectPath              = "public/v1/token/introspect"
	viewUsersDetailsPath                 = "public/v1/users"
	viewChannelsDetailsPath              = "public/v1/channels"
	updateChannelDetailsPath             = "public/v1/channels"
	sendChatMessagePath                  = "public/v1/chat"
	banUserPath                          = "public/v1/moderation/bans"
	liftBanPath                          = "public/v1/moderation/bans"
	viewLivestreamsDetailsPath           = "public/v1/livestreams"
	viewCurrentUserLivestreamDetailsPath = "public/v1/livestreams/stats"
	viewWebhookPublicKeyPath             = "public/v1/public-key"
	viewEventsSubscriptionsDetailsPath   = "public/v1/events/subscriptions"
	registerEventsSubscriptionsPath      = "public/v1/events/subscriptions"
	removeEventsSubscriptionsPath        = "public/v1/events/subscriptions"
)

// UserAuthorizationURL is where the user can log in and approve the application's access request.
func UserAuthorizationURL() string {
	return helpers.ConcatURL(idHostname, userAuthorizationPath)
}

// CodeExchangeURL is the URL to exchange the received code into an access token pair
func CodeExchangeURL() string {
	return helpers.ConcatURL(idHostname, codeExchangePath)
}

// GenerateAppAccessTokenURL is the URL to create an access token pair for the app
func GenerateAppAccessTokenURL() string {
	return helpers.ConcatURL(idHostname, generateAppAccessTokenPath)
}

// RevokeTokenURL is the url to revoke access tokens or refresh tokens
func RevokeTokenURL() string {
	return helpers.ConcatURL(idHostname, revokeTokenPath)
}

// SearchCategoriesURL is the url to search for and retrieve categories
func SearchCategoriesURL() string {
	return helpers.ConcatURL(apiHostname, searchCategoriesPath)
}

// ViewCategoryDetailsURL is the url to retrieve details of a specific category
func ViewCategoryDetailsURL(categoryID int) string {
	return helpers.ConcatURL(apiHostname, viewCategoryDetailsPath, strconv.Itoa(categoryID))
}

// ViewTokenIntrospectURL is the url to retrieve details of the current user's tokens
func ViewTokenIntrospectURL() string {
	return helpers.ConcatURL(apiHostname, viewTokenIntrospectPath)
}

// ViewUsersDetailsURL is the url to retrieve details of one or more users
func ViewUsersDetailsURL() string {
	return helpers.ConcatURL(apiHostname, viewUsersDetailsPath)
}

// ViewChannelsDetailsURL is the url to retrieve details of one or more channels
func ViewChannelsDetailsURL() string {
	return helpers.ConcatURL(apiHostname, viewChannelsDetailsPath)
}

// UpdateChannelDetailsURL is the url to partially update a channels details
func UpdateChannelDetailsURL() string {
	return helpers.ConcatURL(apiHostname, updateChannelDetailsPath)
}

// SendChatMessageURL is the url to send a chat message to a channel
func SendChatMessageURL() string {
	return helpers.ConcatURL(apiHostname, sendChatMessagePath)
}

// BanUserURL is the url to ban a user from a channel
func BanUserURL() string {
	return helpers.ConcatURL(apiHostname, banUserPath)
}

// LiftBanURL is the url to remove a ban of a user from a channel
func LiftBanURL() string {
	return helpers.ConcatURL(apiHostname, liftBanPath)
}

// ViewLivestreamsDetailsURL is the url to retrieve details of one or more livestreams
func ViewLivestreamsDetailsURL() string {
	return helpers.ConcatURL(apiHostname, viewLivestreamsDetailsPath)
}

// ViewCurrentUserLivestreamDetailsURL is the url to retrieve details of the current users livestream
func ViewCurrentUserLivestreamDetailsURL() string {
	return helpers.ConcatURL(apiHostname, viewCurrentUserLivestreamDetailsPath)
}

// ViewWebhookPublicKeyURL is the url to retrieve the public key for the webhook signature
func ViewWebhookPublicKeyURL() string {
	return helpers.ConcatURL(apiHostname, viewWebhookPublicKeyPath)
}

// ViewEventsSubscriptionsDetailsURL is the url to retrieve webhook events subscriptions details
func ViewEventsSubscriptionsDetailsURL() string {
	return helpers.ConcatURL(apiHostname, viewEventsSubscriptionsDetailsPath)
}

// RegisterEventsSubscriptionsURL is the url to register webhook events subscriptions
func RegisterEventsSubscriptionsURL() string {
	return helpers.ConcatURL(apiHostname, registerEventsSubscriptionsPath)
}

// RemoveEventsSubscriptionsURL is the url to remove webhook events subscriptions
func RemoveEventsSubscriptionsURL() string {
	return helpers.ConcatURL(apiHostname, removeEventsSubscriptionsPath)
}
