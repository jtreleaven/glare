package glare

import (
    "time"
)

// WebHook defines the webhook resource available in the Layer API.
type WebHook struct {
    ID              string                  `json:"id"`
    URL             string                  `json:"url"`
    Status          string                  `json:"status"`
    StatusReason    string                  `json:"status_reason"`
    CreatedAt       time.Time               `json:"created_at"`
    Version         string                  `json:"version"`
    TargetURL       string                  `json:"target_url"`
    Events          []string                `json:"events"`
    Secret          string                  `json:"secret"`
    Config          map[string]interface{}  `json:"config"`
}
