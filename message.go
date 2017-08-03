package glare

import (
	"time"
)

// Message represents a single message resource from the Layer API
type Message struct {
	ID       string `json:"id,omitempty"`
	URL      string `json:"url"`
	IsUnread bool   `json:"is_unread"`
	Parts    []struct {
		ID       string                 `json:"id,omitempty"`
		MimeType string                 `json:"mime_type"`
		Content  map[string]interface{} `json:"content"`
		Body     string                 `json:"body"`
	} `json:"parts"`
	ReceivedAt      *time.Time        `json:"received_at,omitempty"`
	RecipientStatus map[string]string `json:"recipient_status"`
	Sender          struct {
		Name   string `json:"name,omitempty"`
		UserID string `json:"user_id,omitempty"`
	} `json:"sender"`
	SentAt           *time.Time `json:"sent_at,omitempty"`
	FromConversation struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	} `json:"conversation"`
}
