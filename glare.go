package glare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://api.layer.com"

var client = &http.Client{}

// EditRequest represents the body of a PUT request to the Layer API
type EditRequest struct {
	Operation string      `json:"operation"`
	Property  string      `json:"property"`
	Value     interface{} `json:"value"`
}

// Layer is the primary struct that acts as the receiver for the API methods
type Layer struct {
	ID      string
	Token   string
	Version string
}

// ExtractUUID returns the 36 character uuid value at the end of a layer id.
func ExtractUUID(id string) string {
	if len(id) < 36 {
		return id
	}
	return id[len(id)-36:]
}

// New is a convenience method for easily creating
func New(id string, token string, version string) Layer {
	return Layer{ID: id, Token: token, Version: version}
}

// -----------------------------------------------------------------------------
// ------------------------- Conversation Methods ------------------------------
// -----------------------------------------------------------------------------

// GetConversationsByUser is the method for retrieving all conversations
// from the perspective of a user.
func (l Layer) GetConversationsByUser(userID string) ([]Conversation, error) {
	var conversations []Conversation
	url := fmt.Sprintf("%s/apps/%s/users/%s/conversations", baseURL, l.ID, userID)
	res, err := makeLayerGetRequest(url, l.ID, l.Token, false)
	if err != nil {
		return conversations, err
	} else if res.StatusCode < 200 || res.StatusCode > 299 {
		return conversations, fmt.Errorf("Status Code: %d, Status: %s\n%+v\n", res.StatusCode, res.Status, res)
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
	res, err := makeLayerGetRequest(url, l.ID, l.Token, false)
	if err != nil {
		return conversation, err
	} else if res.StatusCode < 200 || res.StatusCode > 299 {
		return conversation, fmt.Errorf("Status Code: %d, Status: %s\n%+v\n", res.StatusCode, res.Status, res)
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
	url := fmt.Sprintf("%s/apps/%s/conversations/%s", baseURL, l.ID, conversationID)
	res, err := makeLayerGetRequest(url, l.ID, l.Token, false)
	if err != nil {
		return conversation, err
	} else if res.StatusCode < 200 || res.StatusCode > 299 {
		return conversation, fmt.Errorf("Status Code: %d, Status: %s\n%+v\n", res.StatusCode, res.Status, res)
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
	res, err := makeLayerPostRequest(url, l.Token, l.Version, false, false, pending)
	if err != nil {
		return conversation, err
	}

	if err = json.NewDecoder(res.Body).Decode(&conversation); err != nil {
		return conversation, err
	}
	return conversation, nil
}

// EditConversation will make a request to Layer with an EditRequest body to
// modify the properties on the given conversation.
func (l Layer) EditConversation(c Conversation, changes []EditRequest) (Conversation, error) {
	var conversation Conversation
	url := fmt.Sprintf("%s/apps/%s/conversations/%s", baseURL, l.ID, c.ID)
	res, err := makeLayerPostRequest(url, l.Token, l.Version, true, false, changes)
	if err != nil {
		return conversation, err
	} else if res.StatusCode != 200 && res.StatusCode != 201 {
		return conversation, err
	}
	if err = json.NewDecoder(res.Body).Decode(&conversation); err != nil {
		return conversation, err
	}
	return conversation, nil
}

// DeleteConversation will delete an existing conversation and applies
// globally to all members of the conversation and across devices
func (l Layer) DeleteConversation(remove Conversation) error {
	url := fmt.Sprintf("%s/apps/%s/conversations/%s", baseURL, l.ID, remove.ID)
	res, err := makeLayerDeleteRequest(url, l.Token, l.Version, false)
	if err != nil {
		return err
	} else if res.StatusCode != 204 {
		return err
	}
	return nil
}

// -----------------------------------------------------------------------------
// ---------------------------- Message Methods --------------------------------
// -----------------------------------------------------------------------------

// SendMessage will take the given Message object and Post that data to the
// Layer API for the given conversation.
func (l Layer) SendMessage(m Message, c Conversation) (Message, error) {
	var message Message
	url := fmt.Sprintf("%s/apps/%s/conversations/%s/messages", baseURL, l.ID, c.ID)
	res, err := makeLayerPostRequest(url, l.Token, l.Version, false, false, m)
	if err != nil {
		return message, err
	} else if res.StatusCode != 201 {
		return message, fmt.Errorf("Unable to send message to layer, got response code %d", res.StatusCode)
	}
	if err = json.NewDecoder(res.Body).Decode(&message); err != nil {
		return message, err
	}

	return message, nil
}

// RetrieveMessages will return a slice of messages from the given conversation
// which pertains to the System perspective.
func (l Layer) RetrieveMessages(c Conversation) ([]Message, error) {
	var messages []Message
	url := fmt.Sprintf("%s/apps/%s/conversations/%s/messages", baseURL, l.ID, c.ID)
	res, err := makeLayerGetRequest(url, l.Token, l.Version, false)
	if err != nil {
		return messages, err
	} else if res.StatusCode != 200 {
		return messages, err
	}

	if err = json.NewDecoder(res.Body).Decode(&messages); err != nil {
		return messages, err
	}
	return messages, nil
}

// RetrieveMessagesByUser will return a slice of message objects that are
// associated to the given userID and conversation
func (l Layer) RetrieveMessagesByUser(userID string, c Conversation) ([]Message, error) {
	var messages []Message
	url := fmt.Sprintf("%s/apps/%s/users/%s/conversations/%s/messages", baseURL, l.ID, userID, c.ID)
	res, err := makeLayerGetRequest(url, l.Token, l.Version, false)
	if err != nil {
		return messages, err
	} else if res.StatusCode != 200 {
		return messages, err
	}

	if err = json.NewDecoder(res.Body).Decode(&messages); err != nil {
		return messages, err
	}
	return messages, nil
}

// DeleteMessage will delete the given message from the given conversation.
func (l Layer) DeleteMessage(m Message, c Conversation) error {
	url := fmt.Sprintf("%s/apps/%s/conversations/%s/messages/%s", baseURL, l.ID, c.ID, m.ID)
	_, err := makeLayerDeleteRequest(url, l.Token, l.Version, false)
	if err != nil {
		return err
	}
	return nil
}

// -----------------------------------------------------------------------------
// --------------------------- Identity Methods --------------------------------
// -----------------------------------------------------------------------------

// RegisterIdentity will create a new known user within Layer
func (l Layer) RegisterIdentity(id string, i Identity) error {
	url := fmt.Sprintf("%s/apps/%s/users/%s/identity", baseURL, l.ID, id)
	_, err := makeLayerPostRequest(url, l.Token, l.Version, false, false, i)
	return err
}

// UpdateIdentity will change the Identity match the given id with the
// new value passed into EditRequest.
func (l Layer) UpdateIdentity(id string, changes EditRequest) (Identity, error) {
	var identity Identity
	url := fmt.Sprintf("%s/apps/%s/users/%s/identity", baseURL, l.ID, id)
	res, err := makeLayerPostRequest(url, l.Token, l.Version, true, false, changes)
	if err != nil {
		return identity, err
	}

	if err = json.NewDecoder(res.Body).Decode(&identity); err != nil {
		return identity, err
	}
	return identity, nil
}

// RetrieveIdentity will fetch the identity matching the given id from the Layer API
func (l Layer) RetrieveIdentity(id string) (Identity, error) {
	var identity Identity
	url := fmt.Sprintf("%s/apps/%s/users/%s/identity", baseURL, l.ID, id)
	res, err := makeLayerGetRequest(url, l.Token, l.Version, false)
	if err != nil {
		return identity, err
	}

	if err = json.NewDecoder(res.Body).Decode(&identity); err != nil {
		return identity, err
	}

	return identity, nil
}

// DeleteIdentity will remove an Identity from Layer matching the given ID value
func (l Layer) DeleteIdentity(id string) error {
	url := fmt.Sprintf("%s/apps/%s/users/%s/identity", baseURL, l.ID, id)
	_, err := makeLayerDeleteRequest(url, l.Token, l.Version, false)
	return err
}

// -----------------------------------------------------------------------------
// ---------------------------- WebHook Methods --------------------------------
// -----------------------------------------------------------------------------

// RegisterWebHook will make a post request with the new webhook and return the
// newly created Layer API webhook object.
func (l Layer) RegisterWebHook(created WebHook) (WebHook, error) {
	var webhook WebHook
	url := fmt.Sprintf("%s/apps/%s/webhooks", baseURL, l.ID)
	res, err := makeLayerPostRequest(url, l.Token, l.Version, false, true, created)
	if err != nil {
		return webhook, err
	}

	if err = json.NewDecoder(res.Body).Decode(&webhook); err != nil {
		return webhook, err
	}
	return webhook, nil
}

// ListWebHooks will retrieve all existing WebHooks for your Layer Account.
func (l Layer) ListWebHooks() ([]WebHook, error) {
	var webhooks []WebHook
	url := fmt.Sprintf("%s/apps/%s/webhooks", baseURL, l.ID)
	res, err := makeLayerGetRequest(url, l.Token, l.Version, true)
	if err != nil {
		return webhooks, err
	}

	if err = json.NewDecoder(res.Body).Decode(&webhooks); err != nil {
		return webhooks, err
	}
	return webhooks, nil
}

// GetWebHook will retrieve an existing WebHook from your Layer Account matching
// the given ID.
func (l Layer) GetWebHook(id string) (WebHook, error) {
	var webhook WebHook
	url := fmt.Sprintf("%s/apps/%s/webhooks/%s", baseURL, l.ID, id)
	res, err := makeLayerGetRequest(url, l.Token, l.Version, true)
	if err != nil {
		return webhook, err
	}

	if err = json.NewDecoder(res.Body).Decode(&webhook); err != nil {
		return webhook, err
	}
	return webhook, nil
}

// ActivateWebHook will make a request to Layer to activate the given WebHook
func (l Layer) ActivateWebHook(w WebHook) (WebHook, error) {
	var webhook WebHook
	url := fmt.Sprintf("%s/apps/%s/webhooks/%s/activate", baseURL, l.ID, w.ID)
	res, err := makeLayerPostRequest(url, l.Token, l.Version, false, true, w)
	if err != nil {
		return webhook, err
	}

	if err = json.NewDecoder(res.Body).Decode(&webhook); err != nil {
		return webhook, err
	}

	return webhook, nil
}

// DeactivateWebHook will do the opposite of the activate function and deactivate
// the given webhook to no longer be sent data
func (l Layer) DeactivateWebHook(w WebHook) (WebHook, error) {
	var webhook WebHook
	url := fmt.Sprintf("%s/apps/%s/webhooks/%s/deactivate", baseURL, l.ID, w.ID)
	res, err := makeLayerPostRequest(url, l.Token, l.Version, false, true, w)
	if err != nil {
		return webhook, err
	}

	if err = json.NewDecoder(res.Body).Decode(&webhook); err != nil {
		return webhook, err
	}

	return webhook, nil
}

// DeleteWebHook will remove the given WebHook instance from your Layer Account
func (l Layer) DeleteWebHook(w WebHook) error {
	url := fmt.Sprintf("%s/apps/%s/webhooks/%s", baseURL, l.ID, w.ID)
	_, err := makeLayerDeleteRequest(url, l.Token, l.Version, true)
	return err
}

// -----------------------------------------------------------------------------
// --------------------------- PRIVATE FUNCTIONS -------------------------------
// -----------------------------------------------------------------------------

func makeLayerGetRequest(url string, token string, version string, isWebhook bool) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &http.Response{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if isWebhook {
		req.Header.Add("Accept", fmt.Sprintf("application/vnd.layer.webhooks+json; version=%s", version))
	} else {
		req.Header.Add("Accept", fmt.Sprintf("application/vnd.layer+json; version=%s", version))
	}
	return client.Do(req)
}

func makeLayerPostRequest(url string, token string, version string, isPatch bool, isWebhook bool, body interface{}) (*http.Response, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return &http.Response{}, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(buf))
	if err != nil {
		return &http.Response{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if isWebhook {
		req.Header.Add("Accept", fmt.Sprintf("application/vnd.layer.webhooks+json; version=%s", version))
	} else {
		req.Header.Add("Accept", fmt.Sprintf("application/vnd.layer+json; version=%s", version))
	}
	if isPatch {
		req.Header.Add("X-HTTP-Method-Override", "PATCH")
		req.Header.Add("Content-Type", "application/vnd.layer-patch+json")
	} else {
		req.Header.Add("Content-Type", "application/json")
	}

	return client.Do(req)
}

func makeLayerDeleteRequest(url string, token string, version string, isWebhook bool) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return &http.Response{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if isWebhook {
		req.Header.Add("Accept", fmt.Sprintf("application/vnd.layer.webhooks+json; version=%s", version))
	} else {
		req.Header.Add("Accept", fmt.Sprintf("application/vnd.layer+json; version=%s", version))
	}

	return client.Do(req)
}
