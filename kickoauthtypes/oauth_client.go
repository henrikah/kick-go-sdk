package kickoauthtypes

import (
	"github.com/henrikah/kick-go-sdk/v2/internal/httpclient"
)

type OAuthClientConfig struct {
	ClientID     string
	ClientSecret string
	HTTPClient   httpclient.ClientInterface
}
