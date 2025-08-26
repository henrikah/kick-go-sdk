// Package httpclient provides an abstraction over http.Client for testing.
// ClientInterface can be implemented by any type that performs HTTP requests.
// The compile-time check ensures *http.Client satisfies this interface.
package httpclient

import "net/http"

type ClientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

var _ ClientInterface = (*http.Client)(nil)
