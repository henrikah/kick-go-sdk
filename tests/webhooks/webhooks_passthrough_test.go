package kick_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/henrikah/kick-go-sdk"
	"github.com/henrikah/kick-go-sdk/kickwebhooktypes"
)

func Test_WebookPassthroughInstatiation_Success(t *testing.T) {
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

func Test_WebookPassthroughInstatiationWithErrorCallback_Success(t *testing.T) {
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

// fake body struct
type testPayload struct {
	Message string `json:"message"`
}

func Test_WebhookPassthroughHandler_Success(t *testing.T) {
	privKey, pubPEM := generateKeyPair(t)
	client, err := kick.NewWebhookClient(pubPEM)
	if err != nil {
		t.Fatal(err)
	}

	payload := testPayload{Message: "test-message"}
	body, _ := json.Marshal(payload)
	messageID := "test-message-id"
	timestamp := "test-timestamp"

	sig := signPayload(t, privKey, messageID, timestamp, body)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	req.Header.Set("Kick-Event-Message-Id", messageID)
	req.Header.Set("Kick-Event-Message-Timestamp", timestamp)
	req.Header.Set("Kick-Event-Signature", sig)

	w := httptest.NewRecorder()

	handler := client.WebhookPassthroughHandler(func(w http.ResponseWriter, r *http.Request, h kickwebhooktypes.KickWebhookHeaders) {
		var got testPayload
		err := json.NewDecoder(r.Body).Decode(&got)
		if err != nil {
			t.Fatalf("decoder failed: %v", err)
		}
		if payload.Message != got.Message {
			t.Fatal("message in and out are not equal")
		}
		w.WriteHeader(http.StatusOK)
	})

	handler(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Fatalf("expected 200 status code, got: %d", w.Result().StatusCode)
	}
}
func Test_WebhookPassthroughHandlerIncorrectSignature_Unauthorized(t *testing.T) {
	privKey, pubPEM := generateKeyPair(t)
	client, err := kick.NewWebhookClient(pubPEM)
	if err != nil {
		t.Fatal(err)
	}

	payload := testPayload{Message: "test-message"}
	body, _ := json.Marshal(payload)
	messageID := "test-message-id"
	timestamp := "test-timestamp"

	sig := signPayload(t, privKey, messageID+"error", timestamp, body)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
	req.Header.Set("Kick-Event-Message-Id", messageID)
	req.Header.Set("Kick-Event-Message-Timestamp", timestamp)
	req.Header.Set("Kick-Event-Signature", sig)

	w := httptest.NewRecorder()

	handler := client.WebhookPassthroughHandler(func(w http.ResponseWriter, r *http.Request, h kickwebhooktypes.KickWebhookHeaders) {
		var got testPayload
		err := json.NewDecoder(r.Body).Decode(&got)
		if err != nil {
			t.Fatalf("decoder failed: %v", err)
		}
		if payload.Message != got.Message {
			t.Fatal("message in and out are not equal")
		}
		w.WriteHeader(http.StatusOK)
	})

	handler(w, req)

	if w.Result().StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected 401 status code, got: %d", w.Result().StatusCode)
	}
}
