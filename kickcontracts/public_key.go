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
	//	client, err := kick.NewAPIClient(kickapitypes.APIClientConfig{...})
	//	if err != nil {
	//	    log.Fatal(err)
	//	}
	//
	//	publicKeyResp, err := client.PublicKey().GetWebhookPublicKey(context.TODO())
	//	if err != nil {
	//  	log.Printf("could not get webhook public key: %v", err)
	//  	return nil, err
	//	}
	GetWebhookPublicKey(ctx context.Context) (*kickapitypes.PublicKeyResponse, error)
}
