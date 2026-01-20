package kick

import (
	"context"

	"github.com/henrikah/kick-go-sdk/v2/internal/endpoints"
	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
	"github.com/henrikah/kick-go-sdk/v2/kickcontracts"
)

type publicKeyClient struct {
	client *apiClient
}

func newPublicKeyService(client *apiClient) kickcontracts.PublicKey {
	return &publicKeyClient{
		client: client,
	}
}
func (c *publicKeyClient) GetWebhookPublicKey(ctx context.Context) (*kickapitypes.PublicKeyResponse, error) {
	var publicKeyResponse kickapitypes.PublicKeyResponse

	if err := c.client.requester.MakeGetRequest(ctx, endpoints.ViewWebhookPublicKeyURL(), nil, &publicKeyResponse); err != nil {
		return nil, err
	}

	return &publicKeyResponse, nil
}
