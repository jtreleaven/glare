package glare

import (
    "encoding/json"
    "net/http"
    "bytes"
    "fmt"
)

const baseURL = "https://api.layer.com"
var client = &http.Client{}
// Layer is the primary struct that acts as the receiver for the API methods
type Layer struct {
    ID      string
    Token   string
    Version string
}

// New is a convenience method for easily creating
func New(id string, token string, version string) Layer {
    return Layer{ID: id, Token: token, Version: version}
}

// GetConversationsByUser is the method for retrieving all conversations
// from the perspective of a user.
func (l Layer) GetConversationsByUser(userID string) ([]Conversation, error) {
    var conversations []Conversation
    url := fmt.Sprintf("%s/apps/%s/users/%s/conversations", baseURL, l.ID, userID)
    res, err := makeLayerGetRequest(url, l.ID, l.Token)
    if err != nil {
        return conversations, err
    } else if res.StatusCode != 200 {
        return conversations, fmt.Errorf("Status Code: %d, Status: %s", res.StatusCode, res.Status)
    }

    if err = json.NewDecoder(res.Body).Decode(&conversations); err != nil {
        return conversations, err
    }

    return conversations, nil
}

// GetConversationByUser is the method for retrieving a conversation
// from the perspective of a user.
func (l Layer) GetConversationByUser(userID string, conversationID string) (Conversation, error) {
    var conversation Conversation
    url := fmt.Sprintf("%s/apps/%s/users/%s/conversations/%s", baseURL, l.ID, userID, conversationID)
    res, err := makeLayerGetRequest(url, l.ID, l.Token)
    if err != nil {
        return conversation, err
    } else if res.StatusCode != 200 {
        return conversation, fmt.Errorf("Status Code: %d, Status: %s", res.StatusCode, res.Status)
    }

    if err = json.NewDecoder(res.Body).Decode(&conversation); err != nil {
        return conversation, err
    }

    return conversation, nil
}

// GetConversationByID is the method for retrieving a conversation from the
// perspective of the system with only the conversation UUID
func (l Layer) GetConversationByID(conversationID string) (Conversation, error) {
    var conversation Conversation
    url := fmt.Sprintf("%s/apps/%s/users/%s/conversations/%s", baseURL, l.ID, userID, conversationID)
    res, err := makeLayerGetRequest(url, l.ID, l.Token)
    if err != nil {
        return conversation, err
    } else if res.StatusCode != 200 {
        return conversation, fmt.Errorf("Status Code: %d, Status: %s", res.StatusCode, res.Status)
    }

    if err = json.NewDecoder(res.Body).Decode(&conversation); err != nil {
        return conversation, err
    }

    return conversation, nil
}

// CreateConversation will make a request to Layer for a new Conversation to
// be created using the given conversation object.
func (l Layer) CreateConversation(pending Conversation) (Conversation, error) {
    var conversation Conversation
    url := fmt.Sprintf("%s/apps/%s/conversations", baseURL, l.ID)
    res, err := makeLayerPostRequest(url, l.Token, l.Version, false, pending)
    if err != nil {
        return conversation, err
    }
}


// -------------------- PRIVATE FUNCTIONS -------------------------

func makeLayerGetRequest(url string, token string, version string) (*http.Response, error) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return &http.Response{}, err
    }
    req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
    req.Header.Add("Accept", fmt.Sprintf("application/vnd.layer.webhooks+json; version=%s", version))
    return client.Do(req)
}

func makeLayerPostRequest(url string, token string, version string, isPatch bool, body interface{}) (*http.Response, error) {
    buf, err := json.Marshal(body)
    if err != nil {
        return &http.Response{}, err
    }
    req, err := http.NewRequest("POST", url, bytes.NewReader(buf))
    if err != nil {
        return &http.Response{}, err
    }
    req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
    req.Header.Add("Accept", fmt.Sprintf("application/vnd.layer.webhooks+json; version=%s", version))
    if isPatch {
        req.Header.Add("X-HTTP-Method-Override", "PATCH")
        req.Header.Add("Content-Type", "application/vnd.layer-patch+json")
    } else {
        req.Header.Add("Content-Type", "application/json")
    }

    return client.Do(req)
}
