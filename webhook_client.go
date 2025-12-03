package kick

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"

	"github.com/henrikah/kick-go-sdk/enums/kickwebhookenum"
	"github.com/henrikah/kick-go-sdk/kickerrors"
	"github.com/henrikah/kick-go-sdk/kickwebhooktypes"
)

// webhook handles registering Kick webhook event handlers and serving incoming webhooks.
type webhook interface {
	// RegisterChatMessageSentHandler registers a handler for ChatMessageSent events.
	//
	// Example:
	//
	//	err := webhookClient.RegisterChatMessageSentHandler(func(w http.ResponseWriter, r *http.Request, headers kickwebhooktypes.KickWebhookHeaders, event kickwebhooktypes.ChatMessageSent) {
	//	    fmt.Println("Chat message received:", event)
	//	})
	//	if err != nil {
	//	    log.Printf("could not register ChatMessageSent handler: %v", err)
	//	    return err
	//	}
	RegisterChatMessageSentHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders, kickwebhooktypes.ChatMessageSent)) error

	// RegisterChannelFollowedHandler registers a handler for ChannelFollowed events.
	//
	RegisterChannelFollowedHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders, kickwebhooktypes.ChannelFollowed)) error

	// RegisterChannelSubscriptionRenewalHandler registers a handler for ChannelSubscriptionRenewal events.
	//
	RegisterChannelSubscriptionRenewalHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders, kickwebhooktypes.ChannelSubscriptionRenewal)) error

	// RegisterChannelSubscriptionGiftsHandler registers a handler for ChannelSubscriptionGifts events.
	//
	RegisterChannelSubscriptionGiftsHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders, kickwebhooktypes.ChannelSubscriptionGifts)) error

	// RegisterChannelSubscriptionNewHandler registers a handler for ChannelSubscriptionNew events.
	//
	RegisterChannelSubscriptionNewHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders, kickwebhooktypes.ChannelSubscriptionNew)) error

	// RegisterLivestreamStatusUpdatedHandler registers a handler for LivestreamStatusUpdated events.
	//
	RegisterLivestreamStatusUpdatedHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders, kickwebhooktypes.LivestreamStatusUpdated)) error

	// RegisterLivestreamMetadataUpdatedHandler registers a handler for LivestreamMetadataUpdated events.
	//
	RegisterLivestreamMetadataUpdatedHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders, kickwebhooktypes.LivestreamMetadataUpdated)) error

	// RegisterModerationBannedHandler registers a handler for ModerationBanned events.
	//
	RegisterModerationBannedHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders, kickwebhooktypes.ModerationBanned)) error

	// RegisterKicksGiftedHandler registers a handler for KicksGifted events.
	//
	RegisterKicksGiftedHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders, kickwebhooktypes.KicksGifted)) error

	// RegisterChannelRewardRedemptionUpdatedHandler registers a handler for Channel Reward Redemption events.
	//
	RegisterChannelRewardRedemptionUpdatedHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders, kickwebhooktypes.ChannelRewardRedemptionUpdated)) error

	// WebhookHandler serves incoming webhook HTTP requests.
	//
	// Example:
	//
	//	http.HandleFunc("/webhook", webhookClient.WebhookHandler)
	WebhookHandler(writer http.ResponseWriter, request *http.Request)

	// WebhookPassthroughHandler verifies the signature of the webhook without unmarshaling the body.
	// Ideal for webhook to message queue transports
	//
	// Example:
	//
	//	http.HandleFunc("/webhook", webhookClient.WebhookPassthroughHandler(func (w http.ResponseWriter, r *http.Request, h kickwebhooktypes.KickWebhookHeaders) {
	// 		// Pass into message queue like Redis, RabbitMQ, Kafka, etc
	//	}))
	WebhookPassthroughHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders)) func(http.ResponseWriter, *http.Request)
}

type webhookClient struct {
	onError   func(error)
	handlers  map[kickwebhookenum.WebhookType]func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders)
	publicKey *rsa.PublicKey
}

// NewWebhookClient creates a new WebhookClient instance using the provided public key.
//
// onError is optional and will be called if a webhook handler returns an error.
//
// Example:
//
//	publicKey := "your-public-key"
//	webhookClient, err := kick.NewWebhookClient(publicKey)
//	if err != nil {
//	    log.Fatalf("could not create WebhookClient: %v", err)
//	    return
//	}
//
//	err = webhookClient.RegisterChatMessageSentHandler(func(
//	    writer http.ResponseWriter,
//	    request *http.Request,
//	    headers kickwebhooktypes.KickWebhookHeaders,
//	    event kickwebhooktypes.ChatMessageSent,
//	) {
//	    writer.WriteHeader(http.StatusOK)
//	})
//	if err != nil {
//	    log.Printf("error registering chat message sent handler: %v", err)
//	}
//
//	http.HandleFunc("/", webhookClient.WebhookHandler)
func NewWebhookClient(publicKey string, onError ...func(error)) (webhook, error) {
	if publicKey == "" {
		return nil, &kickerrors.ValidationError{
			Field:   "publicKey",
			Message: "cannot be empty",
		}
	}

	decodePublicKey, _ := pem.Decode([]byte(publicKey))
	if decodePublicKey == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the public key")
	}
	publicKeyParsed, err := x509.ParsePKIXPublicKey(decodePublicKey.Bytes)
	if err != nil {
		return nil, err
	}
	publicKeyAsserted, ok := publicKeyParsed.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	errorCB := func(err error) {
		fmt.Print(err)
	}
	if len(onError) > 0 {
		errorCB = onError[0]
	}

	return &webhookClient{
		handlers:  make(map[kickwebhookenum.WebhookType]func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders)),
		onError:   errorCB,
		publicKey: publicKeyAsserted,
	}, nil
}

func (c *webhookClient) RegisterChatMessageSentHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders, kickwebhooktypes.ChatMessageSent)) error {
	if _, exists := c.handlers[kickwebhookenum.ChatMessageSent]; exists {
		return kickerrors.WebhookHandlerExists("ChatMessageSent")
	}
	c.handlers[kickwebhookenum.ChatMessageSent] = func(writer http.ResponseWriter, request *http.Request, kickHeaders kickwebhooktypes.KickWebhookHeaders) {
		data, err := decodeJSON[kickwebhooktypes.ChatMessageSent](request)
		if err != nil {
			c.onError(kickerrors.SetInternalWebhookError(kickHeaders.MessageID, err))
			return
		}
		handler(writer, request, kickHeaders, *data)
	}
	return nil
}

func (c *webhookClient) RegisterChannelFollowedHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders, kickwebhooktypes.ChannelFollowed)) error {
	if _, exists := c.handlers[kickwebhookenum.ChannelFollowed]; exists {
		return kickerrors.WebhookHandlerExists("ChannelFollowed")
	}
	c.handlers[kickwebhookenum.ChannelFollowed] = func(writer http.ResponseWriter, request *http.Request, kickHeaders kickwebhooktypes.KickWebhookHeaders) {
		data, err := decodeJSON[kickwebhooktypes.ChannelFollowed](request)
		if err != nil {
			c.onError(kickerrors.SetInternalWebhookError(kickHeaders.MessageID, err))
			return
		}
		handler(writer, request, kickHeaders, *data)
	}
	return nil
}

func (c *webhookClient) RegisterChannelSubscriptionRenewalHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders, kickwebhooktypes.ChannelSubscriptionRenewal)) error {
	if _, exists := c.handlers[kickwebhookenum.ChannelSubscriptionRenewal]; exists {
		return kickerrors.WebhookHandlerExists("ChannelSubscriptionRenewal")
	}
	c.handlers[kickwebhookenum.ChannelSubscriptionRenewal] = func(writer http.ResponseWriter, request *http.Request, kickHeaders kickwebhooktypes.KickWebhookHeaders) {
		data, err := decodeJSON[kickwebhooktypes.ChannelSubscriptionRenewal](request)
		if err != nil {
			c.onError(kickerrors.SetInternalWebhookError(kickHeaders.MessageID, err))
			return
		}
		handler(writer, request, kickHeaders, *data)
	}
	return nil
}

func (c *webhookClient) RegisterChannelSubscriptionGiftsHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders, kickwebhooktypes.ChannelSubscriptionGifts)) error {
	if _, exists := c.handlers[kickwebhookenum.ChannelSubscriptionGifts]; exists {
		return kickerrors.WebhookHandlerExists("ChannelSubscriptionGifts")
	}
	c.handlers[kickwebhookenum.ChannelSubscriptionGifts] = func(writer http.ResponseWriter, request *http.Request, kickHeaders kickwebhooktypes.KickWebhookHeaders) {
		data, err := decodeJSON[kickwebhooktypes.ChannelSubscriptionGifts](request)
		if err != nil {
			c.onError(kickerrors.SetInternalWebhookError(kickHeaders.MessageID, err))
			return
		}
		handler(writer, request, kickHeaders, *data)
	}
	return nil
}

func (c *webhookClient) RegisterChannelSubscriptionNewHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders, kickwebhooktypes.ChannelSubscriptionNew)) error {
	if _, exists := c.handlers[kickwebhookenum.ChannelSubscriptionNew]; exists {
		return kickerrors.WebhookHandlerExists("ChannelSubscriptionNew")
	}

	c.handlers[kickwebhookenum.ChannelSubscriptionNew] = func(writer http.ResponseWriter, request *http.Request, kickHeaders kickwebhooktypes.KickWebhookHeaders) {
		data, err := decodeJSON[kickwebhooktypes.ChannelSubscriptionNew](request)
		if err != nil {
			c.onError(kickerrors.SetInternalWebhookError(kickHeaders.MessageID, err))
			return
		}
		handler(writer, request, kickHeaders, *data)
	}
	return nil
}

func (c *webhookClient) RegisterLivestreamStatusUpdatedHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders, kickwebhooktypes.LivestreamStatusUpdated)) error {
	if _, exists := c.handlers[kickwebhookenum.LivestreamStatusUpdated]; exists {
		return kickerrors.WebhookHandlerExists("LivestreamStatusUpdated")
	}
	c.handlers[kickwebhookenum.LivestreamStatusUpdated] = func(writer http.ResponseWriter, request *http.Request, kickHeaders kickwebhooktypes.KickWebhookHeaders) {
		data, err := decodeJSON[kickwebhooktypes.LivestreamStatusUpdated](request)
		if err != nil {
			c.onError(kickerrors.SetInternalWebhookError(kickHeaders.MessageID, err))
			return
		}
		handler(writer, request, kickHeaders, *data)
	}
	return nil
}

func (c *webhookClient) RegisterLivestreamMetadataUpdatedHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders, kickwebhooktypes.LivestreamMetadataUpdated)) error {
	if _, exists := c.handlers[kickwebhookenum.LivestreamMetadataUpdated]; exists {
		return kickerrors.WebhookHandlerExists("LivestreamMetadataUpdated")
	}
	c.handlers[kickwebhookenum.LivestreamMetadataUpdated] = func(writer http.ResponseWriter, request *http.Request, kickHeaders kickwebhooktypes.KickWebhookHeaders) {
		data, err := decodeJSON[kickwebhooktypes.LivestreamMetadataUpdated](request)
		if err != nil {
			c.onError(kickerrors.SetInternalWebhookError(kickHeaders.MessageID, err))
			return
		}
		handler(writer, request, kickHeaders, *data)
	}
	return nil
}

func (c *webhookClient) RegisterModerationBannedHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders, kickwebhooktypes.ModerationBanned)) error {
	if _, exists := c.handlers[kickwebhookenum.ModerationBanned]; exists {
		return kickerrors.WebhookHandlerExists("ModerationBanned")
	}
	c.handlers[kickwebhookenum.ModerationBanned] = func(writer http.ResponseWriter, request *http.Request, kickHeaders kickwebhooktypes.KickWebhookHeaders) {
		data, err := decodeJSON[kickwebhooktypes.ModerationBanned](request)
		if err != nil {
			c.onError(kickerrors.SetInternalWebhookError(kickHeaders.MessageID, err))
			return
		}
		handler(writer, request, kickHeaders, *data)
	}
	return nil
}

func (c *webhookClient) RegisterKicksGiftedHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders, kickwebhooktypes.KicksGifted)) error {
	if _, exists := c.handlers[kickwebhookenum.KicksGifted]; exists {
		return kickerrors.WebhookHandlerExists("KicksGifted")
	}
	c.handlers[kickwebhookenum.KicksGifted] = func(writer http.ResponseWriter, request *http.Request, kickHeaders kickwebhooktypes.KickWebhookHeaders) {
		data, err := decodeJSON[kickwebhooktypes.KicksGifted](request)
		if err != nil {
			c.onError(kickerrors.SetInternalWebhookError(kickHeaders.MessageID, err))
			return
		}
		handler(writer, request, kickHeaders, *data)
	}
	return nil
}

func (c *webhookClient) RegisterChannelRewardRedemptionUpdatedHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders, kickwebhooktypes.ChannelRewardRedemptionUpdated)) error {
	if _, exists := c.handlers[kickwebhookenum.ChannelRewardRedemptionUpdated]; exists {
		return kickerrors.WebhookHandlerExists("ChannelRewardRedemptionUpdated")
	}
	c.handlers[kickwebhookenum.ChannelRewardRedemptionUpdated] = func(writer http.ResponseWriter, request *http.Request, kickHeaders kickwebhooktypes.KickWebhookHeaders) {
		data, err := decodeJSON[kickwebhooktypes.ChannelRewardRedemptionUpdated](request)
		if err != nil {
			c.onError(kickerrors.SetInternalWebhookError(kickHeaders.MessageID, err))
			return
		}
		handler(writer, request, kickHeaders, *data)
	}
	return nil
}

func (c *webhookClient) WebhookHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	kickHeaders := processKickHeaders(request)

	var err error

	eventType := kickwebhookenum.WebhookType(kickHeaders.Type)

	handler, ok := c.handlers[eventType]
	if !ok || handler == nil {
		c.onError(kickerrors.SetInternalWebhookError(kickHeaders.MessageID, kickerrors.SetWebhookHandlerError(kickHeaders.Type, "not found")))
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		c.onError(kickerrors.SetInternalWebhookError(kickHeaders.MessageID, err))
		writer.WriteHeader(http.StatusBadRequest)
	}
	request.Body = io.NopCloser(bytes.NewReader(body))

	err = c.verifySignature(kickHeaders.MessageID, kickHeaders.MessageTimestamp, body, []byte(kickHeaders.Signature))
	if err != nil {
		c.onError(kickerrors.SetInternalWebhookError(kickHeaders.MessageID, err))
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	handler(writer, request, kickHeaders)
}

func (c *webhookClient) WebhookPassthroughHandler(handler func(http.ResponseWriter, *http.Request, kickwebhooktypes.KickWebhookHeaders)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		kickHeaders := processKickHeaders(request)

		var err error

		body, err := io.ReadAll(request.Body)
		if err != nil {
			c.onError(kickerrors.SetInternalWebhookError(kickHeaders.MessageID, err))
			writer.WriteHeader(http.StatusBadRequest)
		}
		request.Body = io.NopCloser(bytes.NewReader(body))

		err = c.verifySignature(kickHeaders.MessageID, kickHeaders.MessageTimestamp, body, []byte(kickHeaders.Signature))
		if err != nil {
			c.onError(kickerrors.SetInternalWebhookError(kickHeaders.MessageID, err))
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		handler(writer, request, kickHeaders)
	}
}

func processKickHeaders(request *http.Request) kickwebhooktypes.KickWebhookHeaders {
	return kickwebhooktypes.KickWebhookHeaders{
		MessageID:        request.Header.Get("Kick-Event-Message-Id"),
		SubscriptionID:   request.Header.Get("Kick-Event-Subscription-Id"),
		Signature:        request.Header.Get("Kick-Event-Signature"),
		MessageTimestamp: request.Header.Get("Kick-Event-Message-Timestamp"),
		Type:             request.Header.Get("Kick-Event-Type"),
		Version:          request.Header.Get("Kick-Event-Version"),
	}
}

func decodeJSON[T any](request *http.Request) (*T, error) {
	var data T
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&data)
	return &data, err
}

func (c *webhookClient) verifySignature(messageID string, timestamp string, body []byte, requestSignature []byte) error {
	signature := fmt.Appendf(nil, "%s.%s.%s", messageID, timestamp, body)
	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(requestSignature)))

	n, err := base64.StdEncoding.Decode(decoded, requestSignature)
	if err != nil {
		return err
	}

	requestSignature = decoded[:n]
	hashed := sha256.Sum256(signature)

	return rsa.VerifyPKCS1v15(c.publicKey, crypto.SHA256, hashed[:], requestSignature)
}
