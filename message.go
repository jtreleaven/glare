package glare

import (
    "time"
)

// Message represents a single message resource from the Layer API
type Message struct {
    ID                  string                  `json:"id"`
    URL                 string                  `json:"url"`
    IsUnread            bool                    `json:"is_unread"`
    Parts               []MessagePart           `json:"parts"`
    ReceivedAt          time.Time               `json:"received_at"`
    RecipientStatus     map[string]string       `json:"recipient_status"`
    Sender              Actor                   `json:"sender"`
    SentAt              time.Time               `json:"send_at"`
    FromConversation    MessageConversation     `json:"conversation"`
}

// MessageConversation is a holder for indicating what Conversation a Message
// belongs to without causing a cirular dependency.
type MessageConversation struct {
    ID      string      `json:"id"`
    URL     string      `json:"url"`
}

// Actor is a holder for linking a message to the User who sent it
// without causing a circular dependency.
type Actor struct {
    Name    string      `json:"name,omitempty"`
    UserID  string      `json:"user_id,omitempty"`
}

// MessagePart represents a single part of a message body.
type MessagePart struct {
    ID          string                  `json:"id"`
    MimeType    string                  `json:"mime_type"`
    Content     map[string]interface{}  `json:"content"`
    Body        string                  `json:"body"`
}
