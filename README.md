# Kick Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/henrikah/kick-go-sdk.svg)](https://pkg.go.dev/github.com/henrikah/kick-go-sdk)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/henrikah/kick-go-sdk)](https://github.com/henrikah/kick-go-sdk/releases)
[![Build](https://github.com/henrikah/kick-go-sdk/actions/workflows/main.yaml/badge.svg)](https://github.com/henrikah/kick-go-sdk/actions/workflows/main.yaml)

A Go SDK for interacting with the [Kick API](https://docs.kick.com/) and handling webhooks.

This SDK provides clients for accessing Kick's API and for processing webhook events. Designed to be easy to use.

---

## Features

- **API Client** – Access Kick API endpoints such as users, livestreams, events, and moderation.
- **OAuth Client** – Support for PKCE OAuth2 flow and app access tokens.
- **Webhook Client** – Handle incoming Kick webhook events securely using the public key.
- **Typed Event Data** – Handlers receive typed structs for each event type.
- **Combined Workflow** – Retrieve the webhook public key directly from the API to set up your webhook client automatically.

---

## Installation

```bash
go get github.com/henrikah/kick-go-sdk
```

---

## Quickstart: API Client

```go
oAuthClient, err := kick.NewOAuthClient(kickapitypes.APIClientConfig{
    ClientID:     "your-client-id",
    ClientSecret: "your-client-secret",
    HTTPClient:   http.DefaultClient,
})
if err != nil {
    log.Fatalf("could not create OAuthClient: %v", err)
}

apiClient, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
    HTTPClient:   http.DefaultClient,
})
if err != nil {
    log.Fatalf("could not create APIClient: %v", err)
}

accessToken, err := oAuthClient.GetAppAccessToken(context.TODO())
if err != nil {
	var apiErr *kickerrors.APIError
	if errors.As(err, &apiErr) {
		log.Fatalf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	} else {
		log.Fatalf("internal error: %v", err)
	}
}

categorySearchData, err := apiClient.Category().SearchCategories(context.TODO(), accessToken, kickfilters.NewCategoriesFilter().WithNames([]string{"Software Development"}))
if err != nil {
	var apiErr *kickerrors.APIError
	if errors.As(err, &apiErr) {
		log.Fatalf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	} else {
		log.Fatalf("internal error: %v", err)
	}
}
log.Println("Found category:", categorySearchData.Data[0].Name)
```

---

## Quickstart: Webhook Client

```go
webhookClient, err := kick.NewWebhookClient("your-public-key")
if err != nil {
    log.Fatalf("could not create WebhookClient: %v", err)
}

err = webhookClient.RegisterChatMessageSentHandler(func(
    writer http.ResponseWriter,
    request *http.Request,
    headers kickwebhooktypes.KickWebhookHeaders,
    data kickwebhooktypes.ChatMessageSent,
) {
    writer.WriteHeader(http.StatusOK)
})
if err != nil {
    log.Printf("error registering chat message sent handler: %v", err)
}

http.HandleFunc("/webhook", webhookClient.WebhookHandler)
```

---

## Quickstart: Combined API + Webhook Client

You can automatically retrieve the public key from the API and use it to set up the webhook client without manually copying the key.

```go
apiClient, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
    HTTPClient:   http.DefaultClient,
})
if err != nil {
    log.Fatalf("could not create APIClient: %v", err)
}

// Fetch the webhook public key directly from the API
publicKeyResp, err := apiClient.PublicKey().GetWebhookPublicKey(context.TODO())
if err != nil {
	var apiErr *kickerrors.APIError
	if errors.As(err, &apiErr) {
		log.Fatalf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	} else {
		log.Fatalf("internal error: %v", err)
	}
}

webhookClient, err := kick.NewWebhookClient(publicKeyResp.Data.PublicKey)
if err != nil {
    log.Fatalf("could not create WebhookClient: %v", err)
}

// Register handlers
err = webhookClient.RegisterChatMessageSentHandler(func(writer http.ResponseWriter, request *http.Request, headers kickwebhooktypes.KickWebhookHeaders, data kickwebhooktypes.ChatMessageSent) {
    writer.WriteHeader(http.StatusOK)
})
if err != nil {
    log.Printf("error registering chat message sent handler: %v", err)
}

http.HandleFunc("/webhook", webhookClient.WebhookHandler)
```

---

## Disclaimer

The Kick API documentation is sometimes ambiguous. This SDK makes educational guesses about certain endpoints and behaviors. There may be limitations or quirks in the API that are not handled by this SDK.
