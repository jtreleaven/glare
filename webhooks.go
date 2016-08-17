package glare

import (
	"time"
)

// WebHook defines the webhook resource available in the Layer API.
type WebHook struct {
	ID           string                 `json:"id"`
	URL          string                 `json:"url"`
	Status       string                 `json:"status"`
	StatusReason string                 `json:"status_reason"`
	CreatedAt    time.Time              `json:"created_at"`
	Version      string                 `json:"version"`
	TargetURL    string                 `json:"target_url"`
	Events       []string               `json:"events"`
	Secret       string                 `json:"secret"`
	Config       map[string]interface{} `json:"config"`
}

// A WebHookMessagePayload represents the request body sent from a Layer webhook.
type WebHookMessagePayload struct {
	Event struct {
		ID        string    `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		Type      string    `json:"type"`
		Actor     struct {
			Name   string `json:"name,omitempty"`
			UserID string `json:"user_id,omitempty"`
		} `json:"actor"`
	} `json:"event"`
	Message Message                `json:"message"`
	Config  map[string]interface{} `json:"config"`
}

// A WebHookConversationPayload represents the request body sent from a Layer webhook.
type WebHookConversationPayload struct {
	Event struct {
		ID        string    `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		Type      string    `json:"type"`
		Actor     struct {
			Name   string `json:"name,omitempty"`
			UserID string `json:"user_id,omitempty"`
		} `json:"actor"`
	} `json:"event"`
	Conversation Conversation           `json:"conversation"`
	Config       map[string]interface{} `json:"config"`
}

// WebHookEvent contains information about the event that caused the webhook to fire.
type WebHookEvent struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Type      string    `json:"type"`
	Actor     struct {
		Name   string `json:"name,omitempty"`
		UserID string `json:"user_id,omitempty"`
	} `json:"actor"`
}
