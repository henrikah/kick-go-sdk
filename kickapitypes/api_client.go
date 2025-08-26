package kickapitypes

import (
	"github.com/henrikah/kick-go-sdk/internal/httpclient"
)

type APIClientConfig struct {
	ClientID     string
	ClientSecret string
	HTTPClient   httpclient.ClientInterface
}
