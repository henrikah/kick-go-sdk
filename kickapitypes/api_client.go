package kickapitypes

import (
	"github.com/henrikah/kick-go-sdk/v2/internal/httpclient"
)

type APIClientConfig struct {
	HTTPClient httpclient.ClientInterface
}
