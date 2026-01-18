package kickapitypes

import (
	"github.com/henrikah/kick-go-sdk/internal/httpclient"
)

type APIClientConfig struct {
	HTTPClient httpclient.ClientInterface
}
