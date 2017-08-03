package glare

import (
	"time"
)

// Conversation represents a single conversation resource from the Layer API
type Conversation struct {
	ID                 string                 `json:"id,omitempty"`
	URL                string                 `json:"url"`
	MessagesURL        string                 `json:"messages_url"`
	CreatedAt          *time.Time             `json:"created_at,omitempty"`
	Participants       []string               `json:"participants"`
	MetaData           map[string]interface{} `json:"metadata"`
	Distinct           bool                   `json:"distinct"`
	LastMessage        Message                `json:"last_message"`
	UnreadMessageCount int                    `json:"unread_message_count"`
}
