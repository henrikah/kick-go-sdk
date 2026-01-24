package kickerrors

import (
	"errors"
	"fmt"
)

type APIError struct {
	StatusCode int
	Message    string
	URL        string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("kick API error (%d): %s", e.StatusCode, e.Message)
}

func SetAPIError(statusCode int, message string, url string) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Message:    message,
		URL:        url,
	}
}

func IsAPIError(err error) *APIError {
	var apiError *APIError
	if errors.As(err, &apiError) {
		return apiError
	}
	return nil
}
