package kickerrors

import (
	"fmt"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid value for '%s': %s", e.Field, e.Message)
}

func ValidateAccessToken(accessToken string) error {
	if accessToken == "" {
		return &ValidationError{
			Field:   "accessToken",
			Message: "cannot be empty",
		}
	}
	return nil
}

func ValidatePageNumber(pageNumber int) error {
	if pageNumber < 1 {
		return &ValidationError{
			Field:   "pageNumber",
			Message: "cannot be less than one",
		}
	}
	return nil
}

func ValidateCategoryID(categoryID int) error {
	if categoryID < 1 {
		return &ValidationError{
			Field:   "categoryID",
			Message: "cannot be less than one",
		}
	}
	return nil
}

func ValidateBroadcasterUserID(broadcasterUserID int) error {
	if broadcasterUserID < 1 {
		return &ValidationError{
			Field:   "broadcasterUserID",
			Message: "cannot be less than one",
		}
	}
	return nil
}

func ValidateUserID(userID int) error {
	if userID < 1 {
		return &ValidationError{
			Field:   "userID",
			Message: "cannot be less than one",
		}
	}
	return nil
}

func ValidateChatMessage(message string) error {
	if message == "" {
		return &ValidationError{
			Field:   "message",
			Message: "cannot be empty",
		}
	}
	return nil
}

func ValidateNotEmpty(field string, value string) error {
	if value == "" {
		return &ValidationError{
			Field:   field,
			Message: "cannot be empty",
		}
	}
	return nil
}

func ValidateNotNil(field string, param any) error {
	if param == nil {
		return &ValidationError{
			Field:   field,
			Message: "cannot be nil",
		}
	}
	return nil
}

func ValidateMaxItems[T any](field string, values []T, max int) error {
	if len(values) > max {
		return &ValidationError{
			Field:   field,
			Message: fmt.Sprintf("cannot have more than %d items", max),
		}
	}
	return nil
}

func ValidateMinItems[T any](field string, values []T, min int) error {
	if len(values) < min {
		return &ValidationError{
			Field:   field,
			Message: fmt.Sprintf("cannot have less than %d items", min),
		}
	}
	return nil
}

func ValidateMinValue[T int | int64](field string, values T, min int) error {
	if values < T(min) {
		return &ValidationError{
			Field:   field,
			Message: fmt.Sprintf("cannot be less than %d", min),
		}
	}
	return nil
}

func ValidateMaxCharacters(field, value string, max int) error {
	if len(value) > max {
		return &ValidationError{
			Field:   field,
			Message: fmt.Sprintf("cannot be longer than %d characters", max),
		}
	}
	return nil
}

func ValidateBetween[T int | int64](field string, value T, min, max T) error {
	if value < min || value > max {
		return &ValidationError{
			Field:   field,
			Message: fmt.Sprintf("must be between %d and %d", min, max),
		}
	}
	return nil
}
