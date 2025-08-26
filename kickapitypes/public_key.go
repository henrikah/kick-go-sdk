package kickapitypes

type PublicKeyResponse struct {
	Data    PublicKeyData `json:"data"`
	Message string        `json:"message"`
}

type PublicKeyData struct {
	PublicKey string `json:"public_key"`
}
