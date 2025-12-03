package kick_test

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/henrikah/kick-go-sdk"
	"github.com/henrikah/kick-go-sdk/enums/kickchannelrewardstatus"
	"github.com/henrikah/kick-go-sdk/enums/kickwebhookenum"
	"github.com/henrikah/kick-go-sdk/kickwebhooktypes"
)

func generateKeyPair(t *testing.T) (*rsa.PrivateKey, string) {
	t.Helper()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	publicKey := &privateKey.PublicKey

	pubBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		t.Fatal(err)
	}

	pubPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	})

	return privateKey, string(pubPEM)
}

func signPayload(t *testing.T, priv *rsa.PrivateKey, messageID string, timestamp string, body []byte) string {
	t.Helper()

	hashed := sha256.Sum256(fmt.Appendf(nil, "%s.%s.%s", messageID, timestamp, body))

	sig, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed[:])
	if err != nil {
		t.Fatal(err)
	}

	return base64.StdEncoding.EncodeToString(sig)
}

func makeRequest(t *testing.T, payload any, messageID, timestamp, signature, webhookType string) *http.Request {
	t.Helper()
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	req.Header.Set("Kick-Event-Message-Id", messageID)
	req.Header.Set("Kick-Event-Message-Timestamp", timestamp)
	req.Header.Set("Kick-Event-Signature", signature)
	req.Header.Set("Kick-Event-Type", webhookType)
	return req
}

func Test_WebookInstatiation_Success(t *testing.T) {
	// Arrange
	_, publicKey := generateKeyPair(t)

	webhooksClient, err := kick.NewWebhookClient(publicKey)

	// Assert
	if webhooksClient == nil {
		t.Fatal("Expected webhooksClient to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}
}

func Test_WebookInstatiationWithErrorCallback_Success(t *testing.T) {
	// Arrange
	_, publicKey := generateKeyPair(t)

	callback := func(err error) { fmt.Println(err) }
	// Act

	webhooksClient, err := kick.NewWebhookClient(publicKey, callback)

	// Assert
	if webhooksClient == nil {
		t.Fatal("Expected webhooksClient to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}
}
func TestWebhookHandlers(t *testing.T) {
	shouldFail := true

	errorFunc := func(err error) {
		if !shouldFail {
			t.Fatalf("Expected error occurred: %v", err)
		}
	}
	privKey, pubPEM := generateKeyPair(t)
	client, err := kick.NewWebhookClient(pubPEM, errorFunc)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name     string
		webhook  kickwebhookenum.WebhookType
		payload  any
		register func() error
	}{
		{
			name:    "ChatMessageSent",
			webhook: kickwebhookenum.ChatMessageSent,
			payload: kickwebhooktypes.ChatMessageSent{Content: "hi"},
			register: func() error {
				return client.RegisterChatMessageSentHandler(func(w http.ResponseWriter, r *http.Request, h kickwebhooktypes.KickWebhookHeaders, data kickwebhooktypes.ChatMessageSent) {
					if data.Content != "hi" {
						t.Error("wrong data")
					}
				})
			},
		},
		{
			name:    "ChannelFollowed",
			webhook: kickwebhookenum.ChannelFollowed,
			payload: kickwebhooktypes.ChannelFollowed{Follower: kickwebhooktypes.User{UserID: 123}},
			register: func() error {
				return client.RegisterChannelFollowedHandler(func(w http.ResponseWriter, r *http.Request, h kickwebhooktypes.KickWebhookHeaders, data kickwebhooktypes.ChannelFollowed) {
					if data.Follower.UserID != 123 {
						t.Error("wrong data")
					}
				})
			},
		},
		{
			name:    "ChannelSubscriptionRenewal",
			webhook: kickwebhookenum.ChannelSubscriptionRenewal,
			payload: kickwebhooktypes.ChannelSubscriptionRenewal{Subscriber: kickwebhooktypes.User{UserID: 123}},
			register: func() error {
				return client.RegisterChannelSubscriptionRenewalHandler(func(w http.ResponseWriter, r *http.Request, h kickwebhooktypes.KickWebhookHeaders, data kickwebhooktypes.ChannelSubscriptionRenewal) {
					if data.Subscriber.UserID != 123 {
						t.Error("wrong data")
					}
				})
			},
		},
		{
			name:    "ChannelSubscriptionGifts",
			webhook: kickwebhookenum.ChannelSubscriptionGifts,
			payload: kickwebhooktypes.ChannelSubscriptionGifts{Gifter: kickwebhooktypes.User{UserID: 123}},
			register: func() error {
				return client.RegisterChannelSubscriptionGiftsHandler(func(w http.ResponseWriter, r *http.Request, h kickwebhooktypes.KickWebhookHeaders, data kickwebhooktypes.ChannelSubscriptionGifts) {
					if data.Gifter.UserID != 123 {
						t.Error("wrong data")
					}
				})
			},
		},
		{
			name:    "ChannelSubscriptionNew",
			webhook: kickwebhookenum.ChannelSubscriptionNew,
			payload: kickwebhooktypes.ChannelSubscriptionNew{Subscriber: kickwebhooktypes.User{UserID: 123}},
			register: func() error {
				return client.RegisterChannelSubscriptionNewHandler(func(w http.ResponseWriter, r *http.Request, h kickwebhooktypes.KickWebhookHeaders, data kickwebhooktypes.ChannelSubscriptionNew) {
					if data.Subscriber.UserID != 123 {
						t.Error("wrong data")
					}
				})
			},
		},
		{
			name:    "LivestreamStatusUpdated",
			webhook: kickwebhookenum.LivestreamStatusUpdated,
			payload: kickwebhooktypes.LivestreamStatusUpdated{Broadcaster: kickwebhooktypes.User{UserID: 123}},
			register: func() error {
				return client.RegisterLivestreamStatusUpdatedHandler(func(w http.ResponseWriter, r *http.Request, h kickwebhooktypes.KickWebhookHeaders, data kickwebhooktypes.LivestreamStatusUpdated) {
					if data.Broadcaster.UserID != 123 {
						t.Error("wrong data")
					}
				})
			},
		},
		{
			name:    "LivestreamMetadataUpdated",
			webhook: kickwebhookenum.LivestreamMetadataUpdated,
			payload: kickwebhooktypes.LivestreamMetadataUpdated{Broadcaster: kickwebhooktypes.User{UserID: 123}},
			register: func() error {
				return client.RegisterLivestreamMetadataUpdatedHandler(func(w http.ResponseWriter, r *http.Request, h kickwebhooktypes.KickWebhookHeaders, data kickwebhooktypes.LivestreamMetadataUpdated) {
					if data.Broadcaster.UserID != 123 {
						t.Error("wrong data")
					}
				})
			},
		},
		{
			name:    "ModerationBanned",
			webhook: kickwebhookenum.ModerationBanned,
			payload: kickwebhooktypes.ModerationBanned{BannedUser: kickwebhooktypes.User{UserID: 123}},
			register: func() error {
				return client.RegisterModerationBannedHandler(func(w http.ResponseWriter, r *http.Request, h kickwebhooktypes.KickWebhookHeaders, data kickwebhooktypes.ModerationBanned) {
					if data.BannedUser.UserID != 123 {
						t.Error("wrong data")
					}
				})
			},
		},
		{
			name:    "KicksGifted",
			webhook: kickwebhookenum.KicksGifted,
			payload: kickwebhooktypes.KicksGifted{Sender: kickwebhooktypes.User{UserID: 123}},
			register: func() error {
				return client.RegisterKicksGiftedHandler(func(w http.ResponseWriter, r *http.Request, h kickwebhooktypes.KickWebhookHeaders, data kickwebhooktypes.KicksGifted) {
					if data.Sender.UserID != 123 {
						t.Error("wrong data")
					}
				})
			},
		},
		{
			name:    "ChannelRewardRedemptionUpdated",
			webhook: kickwebhookenum.ChannelRewardRedemptionUpdated,
			payload: kickwebhooktypes.ChannelRewardRedemptionUpdated{Redeemer: kickwebhooktypes.Redeemer{UserID: 123}, Status: kickchannelrewardstatus.Accepted},
			register: func() error {
				return client.RegisterChannelRewardRedemptionUpdatedHandler(func(w http.ResponseWriter, r *http.Request, h kickwebhooktypes.KickWebhookHeaders, data kickwebhooktypes.ChannelRewardRedemptionUpdated) {
					if data.Redeemer.UserID != 123 {
						t.Error("wrong data")
					}
					if data.Status != "accepted" {
						t.Error("incorrect enum usage")
					}
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.register(); err != nil {
				t.Fatal(err)
			}

			if err := tt.register(); err == nil {
				t.Fatal("expected error for duplicate handler registration")
			}

			// Valid request
			shouldFail = false
			messageID := "msg1"
			timestamp := "ts1"
			payloadBytes, _ := json.Marshal(tt.payload)
			sig := signPayload(t, privKey, messageID, timestamp, payloadBytes)
			req := makeRequest(t, tt.payload, messageID, timestamp, sig, string(tt.webhook))
			rr := httptest.NewRecorder()

			client.WebhookHandler(rr, req)
			// Invalid request signature
			shouldFail = true
			reqBad := makeRequest(t, tt.payload, messageID, timestamp, "invalidsig", string(tt.webhook))
			rr = httptest.NewRecorder()
			client.WebhookHandler(rr, reqBad)
		})
	}
}
