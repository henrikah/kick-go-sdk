package kickapitypes

type EventSubscription struct {
	Data    []EventSubscriptionData `json:"data"`
	Message string                  `json:"message"`
}

type EventSubscriptionData struct {
	AppID             string `json:"app_id"`
	BroadcasterUserID int    `json:"broadcaster_user_id"`
	CreatedAt         string `json:"created_at"`
	Event             string `json:"event"`
	ID                string `json:"id"`
	Method            string `json:"method"`
	UpdatedAt         string `json:"updated_at"`
	Version           int    `json:"version"`
}

type CreateEventSubscriptionRequest struct {
	BroadcasterUserID *int          `json:"broadcaster_user_id,omitempty"`
	Events            []EventObject `json:"events"`
	Method            string        `json:"method"`
}

type EventObject struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}

type CreateEventSubscriptionsResponse struct {
	Data    []CreateEventSubscriptionData `json:"data"`
	Message string                        `json:"message"`
}

type CreateEventSubscriptionData struct {
	Error          string `json:"error"`
	Name           string `json:"name"`
	SubscriptionID string `json:"subscription_id"`
	Version        int    `json:"version"`
}
