# Changelog

## \[2.0.0] - 2026-01-18

### Added

* Added NewOAuthClient.
* Added Get Channel Reward Redemptions
* Added Accept Channel Reward Redemptions
* Added Reject Channel Reward Redemptions

### Changed

* Decoupled OAuth from NewAPIClient.
* Moved the request helper into an internal package enabling usage across both APIClient and OAuthClient
* Implementation of Categories removed and replaced with the new V2 endpoint.
* LivestreamFilterBuilder renamed to LivestreamsFilter

## \[1.8.0] - 2026-01-17

### Notes

* Final release of major version 1.
* Several endpoints used in this version have been deprecated by Kick.
* Endpoint clients are coupled to OAuth in v1; this will be decoupled in v2.
* No further updates or new features will be released for v1. This version is no longer maintained.

### Added

* Added ProfilePicture field to livestream endpoint.

## \[1.7.0] - 2025-12-05

### Added

* Added PinnedTimeSeconds to kicks.gifted webhook payload.
* Added DeleteChatMessage endpoint.
* Added Channel Reward Redemption to webhooks
* Added Channel Reward endpoints to the API (GET, POST, PATCH, DELETE)

### Bug Fix

* Typo in ProfilePicture on User struct as it was accidentally pluralised (ProfilePictures)

## \[1.6.0] - 2025-11-25

### Added

* Added Tags and ViewerCount fields to category.

### Changed

* Changed GetCategories data into its own type as it no longer is the same data structure as GetCategory

## \[1.5.0] - 2025-11-22

### Added

* Added CustomTags field to livestream and channel endpoints

## \[1.4.0] - 2025-10-28

### Added

* Added kick leaderboard endpoint

## \[1.3.0] - 2025-10-21

### Added

* Added kicks.gifted to webhooks event and types

## \[1.2.1] - 2025-09-11

### Changed

* APIClient function interfaces are now public to enable mocking

## \[1.2.0] - 2025-09-11

### Added

* Proper interface for APIClient

### Changed

* APIClient field names are private and now requires a function call to be retrieved

## \[1.1.1] - 2025-09-04

### Fixed

* Get livestream data for current user URL change

## \[1.1.0] - 2025-09-04

### Added

* Webhook passthrough handler with signature verification for webhooks.
* Get livestream data for the current user

## \[1.0.0] - 2025-08-26

* Initial release
