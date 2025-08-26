package kick_test

import (
	"errors"
	"net/http"
	"strconv"
	"testing"

	"github.com/henrikah/kick-go-sdk"
	"github.com/henrikah/kick-go-sdk/kickapitypes"
	"github.com/henrikah/kick-go-sdk/kickerrors"
	"github.com/henrikah/kick-go-sdk/tests/mocks"
)

func Test_GetCategoriesMissingAccessToken_Error(t *testing.T) {
	// Arrange
	ctx := t.Context()
	httpClient := http.DefaultClient

	accessToken := ""
	searchQuery := "search-query"
	page := 1

	config := kickapitypes.APIClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	categoriesData, err := client.Category.SearchCategories(ctx, accessToken, searchQuery, page)

	// Assert
	if categoriesData != nil {
		t.Fatal("Expected categoriesData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &validationError) {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "accessToken" {
		t.Fatalf("Expected error on field 'accessToken', got '%s'", validationError.Field)
	}
}

func Test_GetCategoriesInvalidPage_Error(t *testing.T) {
	// Arrange
	ctx := t.Context()
	httpClient := http.DefaultClient

	accessToken := "access-token"
	searchQuery := "search-query"
	page := 0

	config := kickapitypes.APIClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	var validationError *kickerrors.ValidationError

	// Act
	categoriesData, err := client.Category.SearchCategories(ctx, accessToken, searchQuery, page)

	// Assert
	if categoriesData != nil {
		t.Fatal("Expected categoriesData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.As(err, &validationError) {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationError.Field != "pageNumber" {
		t.Fatalf("Expected error on field 'pageNumber', got '%s'", validationError.Field)
	}
}

func Test_GetCategoriesUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"message": "Invalid request"}`

	ctx := t.Context()

	accessToken := "access-token"
	searchQuery := "search-query"
	page := 1

	mockClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return mocks.NewMockResponse(http.StatusUnauthorized, errorJSON), nil
		},
	}

	config := kickapitypes.APIClientConfig{
		ClientID:     "test-id",
		ClientSecret: "test-secret",
		HTTPClient:   mockClient,
	}
	client, _ := kick.NewAPIClient(config)

	var apiError *kickerrors.APIError
	// Act
	categoriesData, err := client.Category.SearchCategories(ctx, accessToken, searchQuery, page)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if categoriesData != nil {
		t.Fatal("Expected categoriesData to be nil on error")
	}

	if !errors.As(err, &apiError) {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_GetCategories_Success(t *testing.T) {
	// Arrange
	ctx := t.Context()
	clientID := "test-id"
	clientSecret := "test-secret"

	accessToken := "access-token"
	searchQuery := "search-query"
	page := 1

	expectedJSON := `{
		"data": [{
			"id": 1,
			"name": "test-category",
			"thumbnail": "https://test-thumbnail"
		}], "message": "test-message"
	}`

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != "https://api.kick.com/public/v1/categories?page="+strconv.Itoa(page)+"&q="+searchQuery {
				t.Fatalf("Unexpected request URL: %s", req.URL.String())
			}

			if req.Method != "GET" {
				t.Fatalf("Unexpected request method: %s", req.Method)
			}

			if req.Header.Get("Accept") != "application/json" {
				t.Fatal("Missing Accept header")
			}

			if req.Header.Get("Authorization") != "Bearer "+accessToken {
				t.Fatal("Missing Authorization header")
			}

			return mocks.NewMockResponse(http.StatusOK, expectedJSON), nil
		},
	}

	config := kickapitypes.APIClientConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HTTPClient:   httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act

	categoriesData, err := client.Category.SearchCategories(ctx, accessToken, searchQuery, page)

	// Assert
	if categoriesData == nil {
		t.Fatal("Expected categoriesData to not be nil")
	}

	if err != nil {
		t.Fatal("Expected error to be nil")
	}

	if len(categoriesData.Data) != 1 {
		t.Fatalf("Expected Data to be 1 slice, got %d", len(categoriesData.Data))
	}

	if categoriesData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", categoriesData.Message)
	}
}
