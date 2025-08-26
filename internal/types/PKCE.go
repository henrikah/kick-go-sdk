// Package types contains internal data types used across the SDK.
package types

type PKCE struct {
	CodeVerifier  string
	CodeChallenge string
	Method        string
}
