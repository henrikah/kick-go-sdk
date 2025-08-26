package kickerrors

import "fmt"

type InternalWebhookError struct {
	MessageID string
	Err       error
}

func (e *InternalWebhookError) Error() string {
	return fmt.Sprintf("error on message: '%s', %s", e.MessageID, e.Err)
}

func SetInternalWebhookError(messageID string, err error) *InternalWebhookError {
	return &InternalWebhookError{
		MessageID: messageID,
		Err:       err,
	}
}
