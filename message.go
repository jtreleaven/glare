package glare

import (
    "time"
)

// Message represents a single message resource from the Layer API
type Message struct {
    ID                  string                  `json:"id"`
    URL                 string                  `json:"url"`
    IsUnread            bool                    `json:"is_unread"`
    Parts               []map[string]string     `json:"parts"`
    ReceivedAt          time.Time               `json:"received_at"`
    RecipientStatus     map[string]string       `json:"recipient_status"`
    Sender              MessageUser             `json:"sender"`
    SentAt              time.Time               `json:"send_at"`
    FromConversation    MessageConversation     `json:"conversation"`
}

// MessageConversation is a holder for indicating what Conversation a Message
// belongs to without causing a cirular dependency.
type MessageConversation struct {
    ID      string      `json:"id"`
    URL     string      `json:"url"`
}

// MessageUser is a holder for linking a message to the User who sent it
// without causing a circular dependency.
type MessageUser struct {
    Name    string      `json:"name"`
    UserID  string      `json:"user_id"`
}
