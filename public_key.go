package kick

import (
	"context"

	"github.com/henrikah/kick-go-sdk/internal/endpoints"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickcontracts"
)

type publicKeyClient struct {
	client *apiClient
}

func newPublicKeyClient(client *apiClient) kickcontracts.PublicKey {
	return &publicKeyClient{
		client: client,
	}
}
func (c *publicKeyClient) GetWebhookPublicKey(ctx context.Context) (*kickapitypes.PublicKeyResponse, error) {
	var publicKeyResponse kickapitypes.PublicKeyResponse

	if err := c.client.makeGetRequest(ctx, endpoints.ViewWebhookPublicKeyURL(), nil, &publicKeyResponse); err != nil {
		return nil, err
	}

	return &publicKeyResponse, nil
}
