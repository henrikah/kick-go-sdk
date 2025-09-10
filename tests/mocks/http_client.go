// Package mocks provides mocks for the testing library
package mocks

import (
	"io"
	"net/http"
	"strings"
)

// MockHTTPClient is a test double that implements httpclient.Interface().
// It allows tests to precisely control the HTTP response for a given request.
type MockHTTPClient struct {
	// DoFunc is the function that will be called when Do is invoked.
	// Tests can set this to define the desired behavior for each test case.
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Do calls the configured DoFunc.
func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc == nil {
		// Panic or return a default error if not set, to fail tests explicitly.
		panic("MockHTTPClient's DoFunc is not set")
	}
	return m.DoFunc(req)
}

// NewMockResponse is a helper to create an *http.Response for the mock.
func NewMockResponse(statusCode int, jsonBody string) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(jsonBody)),
	}
}
