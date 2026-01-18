package kickcontracts

import (
	"context"

	"github.com/henrikah/kick-go-sdk/kickapitypes"
)

// PublicKey handles operations related to kick's public keys.
//
// All examples use context.TODO() as a placeholder. Replace with a proper
// context (with timeout/cancel) in production code.
type PublicKey interface {
	// GetWebhookPublicKey retrieves the public key used to verify incoming webhooks.
	//
	// Example:
	//
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{
	//	    HTTPClient: http.DefaultClient,
	//	})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	publicKeyResp, err := client.PublicKey().GetWebhookPublicKey(context.TODO())
	//	if err != nil {
	//		var apiErr *kickerrors.APIError
	//		if errors.As(err, &apiErr) {
	//			log.Printf("API error: %d %s", apiErr.StatusCode, apiErr.Message)
	//		} else {
	//			log.Printf("internal error: %v", err)
	//		}
	//	}
	GetWebhookPublicKey(ctx context.Context) (*kickapitypes.PublicKeyResponse, error)
}
