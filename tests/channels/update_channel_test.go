package kick_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk/v2"
	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
	"github.com/henrikah/kick-go-sdk/v2/tests/mocks"
)

func Test_UpdateChannelMissingAccessToken_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := ""

	channelData := kickapitypes.UpdateChannelRequest{
		CategoryID:  1,
		CustomTags:  []string{},
		StreamTitle: "stream-title",
	}

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act
	err := client.Channel().UpdateChannel(t.Context(), accessToken, channelData)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationErr := kickerrors.IsValidationError(err)

	if validationErr == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationErr.Field != "accessToken" {
		t.Fatalf("Expected error on field 'accessToken', got '%s'", validationErr.Field)
	}
}

func Test_UpdateChannelNegativeCategoryID_Error(t *testing.T) {
	// Arrange
	ctx := t.Context()
	httpClient := http.DefaultClient

	accessToken := "access-token"

	channelData := kickapitypes.UpdateChannelRequest{
		CategoryID:  -1,
		CustomTags:  []string{},
		StreamTitle: "stream-title",
	}

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act
	err := client.Channel().UpdateChannel(ctx, accessToken, channelData)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationErr := kickerrors.IsValidationError(err)

	if validationErr == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationErr.Field != "categoryID" {
		t.Fatalf("Expected error on field 'categoryID', got '%s'", validationErr.Field)
	}
}

func Test_UpdateChannelUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"message": "Invalid request"}`

	accessToken := "access-token"

	channelData := kickapitypes.UpdateChannelRequest{
		CategoryID:  1,
		CustomTags:  []string{},
		StreamTitle: "stream-title",
	}

	mockClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return mocks.NewMockResponse(http.StatusUnauthorized, errorJSON), nil
		},
	}

	config := kickapitypes.APIClientConfig{
		HTTPClient: mockClient,
	}
	client, _ := kick.NewAPIClient(config)

	// Act
	err := client.Channel().UpdateChannel(t.Context(), accessToken, channelData)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	apiErr := kickerrors.IsAPIError(err)
	if apiErr == nil {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_UpdateChannelWithMissingFields_Success(t *testing.T) {
	// Arrange
	accessToken := "access-token"

	channelData := kickapitypes.UpdateChannelRequest{
		CustomTags: []string{},
	}

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/channels" {
				t.Fatalf("Unexpected request URL: %s", req.URL.String())
			}

			if req.Method != "PATCH" {
				t.Fatalf("Unexpected request method: %s", req.Method)
			}

			if req.Header.Get("Accept") != "application/json" {
				t.Fatal("Missing Accept header")
			}

			if req.Header.Get("Content-Type") != "application/json" {
				t.Fatal("Missing Content-Type header")
			}

			if req.Header.Get("Authorization") != "Bearer "+accessToken {
				t.Fatal("Missing Authorization header")
			}

			t.Cleanup(func() {
				defer func() {
					if err := req.Body.Close(); err != nil {
						t.Logf("failed to close request body: %v", err)
					}
				}()
			})

			bodyDecoder := json.NewDecoder(req.Body)

			var updateData map[string]any

			err := bodyDecoder.Decode(&updateData)

			if err != nil {
				return mocks.NewMockResponse(http.StatusInternalServerError, ""), nil
			}

			if _, ok := updateData["category_id"]; ok {
				t.Fatal("Expected category_id to be omitted")
			}

			if _, ok := updateData["stream_title"]; ok {
				t.Fatal("Expected stream_title to be omitted")
			}

			return mocks.NewMockResponse(http.StatusOK, ""), nil
		},
	}

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act

	err := client.Channel().UpdateChannel(t.Context(), accessToken, channelData)

	// Assert

	if err != nil {
		t.Fatal("Expected error to be nil")
	}
}

func Test_UpdateChannel_Success(t *testing.T) {
	// Arrange
	accessToken := "access-token"

	channelData := kickapitypes.UpdateChannelRequest{
		CategoryID:  1,
		CustomTags:  []string{},
		StreamTitle: "stream-title",
	}

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/channels" {
				t.Fatalf("Unexpected request URL: %s", req.URL.String())
			}

			if req.Method != "PATCH" {
				t.Fatalf("Unexpected request method: %s", req.Method)
			}

			if req.Header.Get("Accept") != "application/json" {
				t.Fatal("Missing Accept header")
			}

			if req.Header.Get("Content-Type") != "application/json" {
				t.Fatal("Missing Content-Type header")
			}

			if req.Header.Get("Authorization") != "Bearer "+accessToken {
				t.Fatal("Missing Authorization header")
			}

			t.Cleanup(func() {
				defer func() {
					if err := req.Body.Close(); err != nil {
						t.Logf("failed to close request body: %v", err)
					}
				}()
			})

			bodyDecoder := json.NewDecoder(req.Body)

			var updateData kickapitypes.UpdateChannelRequest

			err := bodyDecoder.Decode(&updateData)

			if err != nil {
				return mocks.NewMockResponse(http.StatusInternalServerError, ""), nil
			}
			return mocks.NewMockResponse(http.StatusOK, ""), nil
		},
	}

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act

	err := client.Channel().UpdateChannel(t.Context(), accessToken, channelData)

	// Assert

	if err != nil {
		t.Fatal("Expected error to be nil")
	}
}
