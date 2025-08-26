package kickerrors

import "fmt"

type WebhookHandlerError struct {
	Type    string
	Message string
}

func (e *WebhookHandlerError) Error() string {
	return fmt.Sprintf("requested handler for type '%s', %s", e.Type, e.Message)
}

func WebhookHandlerNotExists(name string) *WebhookHandlerError {
	return &WebhookHandlerError{
		Type:    name,
		Message: "does not exist",
	}
}
func WebhookHandlerExists(name string) *WebhookHandlerError {
	return &WebhookHandlerError{
		Type:    name,
		Message: "already exist",
	}
}
func SetWebhookHandlerError(name string, message string) *WebhookHandlerError {
	return &WebhookHandlerError{
		Type:    name,
		Message: message,
	}
}
