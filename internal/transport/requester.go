// Package transport handles http requests
package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/henrikah/kick-go-sdk/v2/internal/httpclient"
	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
)

type Requester struct {
	httpClient httpclient.ClientInterface
}

func NewRequester(httpClient httpclient.ClientInterface) (*Requester, error) {
	if err := kickerrors.ValidateNotNil("httpClient", httpClient); err != nil {
		return nil, err
	}
	return &Requester{httpClient: httpClient}, nil
}

func (r *Requester) MakeJSONRequest(ctx context.Context, method, urlStr string, requestBody any, accessToken *string, out any) error {
	var bodyReader io.Reader
	if requestBody != nil {
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}
	return r.makeRequestWithBody(ctx, method, urlStr, bodyReader, "application/json", accessToken, out)
}

func (r *Requester) MakeFormRequest(ctx context.Context, method, urlStr string, requestBody io.Reader, accessToken *string, out any) error {
	return r.makeRequestWithBody(ctx, method, urlStr, requestBody, "application/x-www-form-urlencoded", accessToken, out)
}

func (r *Requester) MakeGetRequest(ctx context.Context, urlStr string, accessToken *string, out any) error {
	return r.makeRequestWithBody(ctx, http.MethodGet, urlStr, nil, "", accessToken, out)
}

func (r *Requester) MakeDeleteRequest(ctx context.Context, urlStr string, accessToken *string, out any) error {
	return r.makeRequestWithBody(ctx, http.MethodDelete, urlStr, nil, "", accessToken, out)
}

func (r *Requester) makeRequestWithBody(ctx context.Context, method, urlStr string, body io.Reader, contentType string, accessToken *string, out any) error {
	req, err := http.NewRequestWithContext(ctx, method, urlStr, body)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")

	if accessToken != nil {
		req.Header.Set("Authorization", "Bearer "+*accessToken)
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("failed to close request body: %v", err)
		}
	}()

	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusNoContent {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return kickerrors.SetAPIError(resp.StatusCode, string(bodyBytes), req.URL.String())
	}

	if resp.StatusCode != http.StatusNoContent && out != nil {
		return json.NewDecoder(resp.Body).Decode(out)
	}

	return nil
}
