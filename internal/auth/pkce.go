// Package auth provides helper functions for authentication-related operations,
// including PKCE (Proof Key for Code Exchange) generation for OAuth flows.
package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"

	"github.com/henrikah/kick-go-sdk/internal/types"
)

// pkceMethodS256 is the only PKCE method currently supported by Kick.
const pkceMethodS256 = "S256"

// GeneratePKCE creates a PKCE (Proof Key for Code Exchange) pair using the S256 method.
// It returns a code verifier and its corresponding code challenge, suitable for use
// in OAuth authorization flows.
func GeneratePKCE() (*types.PKCE, error) {
	verifier, err := codeVerifier()
	if err != nil {
		return nil, err
	}

	challenge := codeChallengeS256(verifier)

	return &types.PKCE{
		CodeVerifier:  verifier,
		CodeChallenge: challenge,
		Method:        pkceMethodS256,
	}, nil
}

func codeVerifier() (string, error) {
	randomBytes := make([]byte, 64)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(randomBytes), nil
}

func codeChallengeS256(verifier string) string {
	sum := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(sum[:])
}
