package kick_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/henrikah/kick-go-sdk/v2"
	"github.com/henrikah/kick-go-sdk/v2/kickapitypes"
	"github.com/henrikah/kick-go-sdk/v2/kickerrors"
	"github.com/henrikah/kick-go-sdk/v2/kickfilters"
	"github.com/henrikah/kick-go-sdk/v2/tests/mocks"
)

func Test_GetCategoriesMissingAccessToken_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := ""

	filter := kickfilters.NewCategoriesFilter().WithCategoryIDs([]int64{42})

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	// Act
	categoriesData, err := client.Category().SearchCategories(t.Context(), accessToken, filter)

	// Assert
	if categoriesData != nil {
		t.Fatal("Expected categoriesData to be nil")
	}

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

func Test_GetCategoriesInvalidCategoryID_Error(t *testing.T) {
	// Arrange
	httpClient := http.DefaultClient

	accessToken := "access-token"
	filter := kickfilters.NewCategoriesFilter().WithCategoryIDs([]int64{-42})

	config := kickapitypes.APIClientConfig{
		HTTPClient: httpClient,
	}
	client, _ := kick.NewAPIClient(config)

	// Act
	categoriesData, err := client.Category().SearchCategories(t.Context(), accessToken, filter)

	// Assert
	if categoriesData != nil {
		t.Fatal("Expected categoriesData to be nil")
	}

	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	validationErr := kickerrors.IsValidationError(err)
	if validationErr == nil {
		t.Fatalf("Expected validation error, got %T", err)
	}

	if validationErr.Field != "CategoryIDs" {
		t.Fatalf("Expected error on field 'CategoryIDs', got '%s'", validationErr.Field)
	}
}

func Test_GetCategoriesUnAuthorized_Error(t *testing.T) {
	// Arrange
	errorJSON := `{"message": "Invalid request"}`

	accessToken := "access-token"
	filter := kickfilters.NewCategoriesFilter().WithCategoryIDs([]int64{42})

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
	categoriesData, err := client.Category().SearchCategories(t.Context(), accessToken, filter)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if categoriesData != nil {
		t.Fatal("Expected categoriesData to be nil on error")
	}

	apiErr := kickerrors.IsAPIError(err)
	if apiErr == nil {
		t.Fatalf("Expected API error, got %T", err)
	}
}

func Test_GetCategories_Success(t *testing.T) {
	// Arrange

	accessToken := "access-token"
	var id int64 = 42
	tag := "GoLang"
	name := "Test Category"
	thumbnail := "https://test-thumbnail"
	filter := kickfilters.NewCategoriesFilter().WithCategoryIDs([]int64{id})

	expectedJSON := fmt.Sprintf(`{
		"data": [{
			"id": %d,
			"tags": ["%s"],
			"name": "%s",
			"thumbnail": "%s"
		}],
		"message": "test-message",
		"pagination": {
			"next_cursor": ""
		}
	}`, id, tag, name, thumbnail)

	httpClient := &mocks.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if req.URL.String() != fmt.Sprintf("https://api.kick.com/public/v2/categories?id=%d", id) {
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
		HTTPClient: httpClient,
	}

	client, _ := kick.NewAPIClient(config)

	// Act

	categoriesData, err := client.Category().SearchCategories(t.Context(), accessToken, filter)

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

	if categoriesData.Data[0].ID != int(id) {
		t.Fatalf("Expected ID to be %d, got %d", id, categoriesData.Data[0].ID)
	}

	if categoriesData.Data[0].Name != name {
		t.Fatalf("Expected Name to be %s, got %s", name, categoriesData.Data[0].Name)
	}

	if categoriesData.Data[0].Thumbnail != thumbnail {
		t.Fatalf("Expected Thumbnail to be %s, got %s", thumbnail, categoriesData.Data[0].Thumbnail)
	}

	if categoriesData.Data[0].Tags[0] != tag {
		t.Fatalf("Expected Tags to be %s, got %s", tag, categoriesData.Data[0].Tags[0])
	}

	if categoriesData.Message != "test-message" {
		t.Fatalf("Expected Message to be %s, got %s", "test-message", categoriesData.Message)
	}
}
