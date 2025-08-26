# Kick Go SDK
[![Go Reference](https://pkg.go.dev/badge/github.com/henrikah/kick-go-sdk.svg)](https://pkg.go.dev/github.com/henrikah/kick-go-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/henrikah/kick-go-sdk)](https://goreportcard.com/report/github.com/henrikah/kick-go-sdk)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/henrikah/kick-go-sdk)](https://github.com/henrikah/kick-go-sdk/releases)
[![Build](https://github.com/henrikah/kick-go-sdk/actions/workflows/ci.yml/badge.svg)](https://github.com/henrikah/kick-go-sdk/actions/workflows/ci.yml)

A Go SDK for interacting with the [Kick API](https://docs.kick.com/) and handling webhooks.

This SDK provides clients for accessing Kick's API and for processing webhook events. Designed to be easy to use.

---

## Features

* **API Client** – Access Kick API endpoints such as users, livestreams, events, and moderation.
* **OAuth** – Support for PKCE OAuth2 flow and app access tokens.
* **Webhook Client** – Handle incoming Kick webhook events securely using the public key.
* **Typed Event Data** – Handlers receive typed structs for each event type.
* **Combined Workflow** – Retrieve the webhook public key directly from the API to set up your webhook client automatically.

---

## Installation

```bash
go get github.com/henrikah/kick-go-sdk
```

---

## Quickstart: API Client

```go
apiClient, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
    ClientID:     "your-client-id",
    ClientSecret: "your-client-secret",
    HTTPClient:   http.DefaultClient,
})
if err != nil {
    log.Fatalf("could not create APIClient: %v", err)
}

currentUser, err := apiClient.User.GetCurrentUser(context.TODO(), "user-access-token")
if err != nil {
    log.Printf("could not get current user: %v", err)
}
log.Println("Logged in as:", currentUser.Data[0].Username)
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
// ClientID and ClientSecret can be dummy text if you only want the get the Public Key
apiClient, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
    ClientID:     "your-client-id",
    ClientSecret: "your-client-secret",
    HTTPClient:   http.DefaultClient,
})
if err != nil {
    log.Fatalf("could not create APIClient: %v", err)
}

// Fetch the webhook public key directly from the API
publicKeyResp, err := apiClient.PublicKey.GetWebhookPublicKey(context.TODO())
if err != nil {
    log.Fatalf("could not get the public key for the webhook: %v", err)
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

## License

MIT License

---

## Disclaimer

The Kick API documentation is sometimes ambiguous. This SDK makes educational guesses about certain endpoints and behaviors. There may be limitations or quirks in the API that are not handled by this SDK.
