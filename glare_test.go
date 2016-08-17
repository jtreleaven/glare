package glare

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

// TestGetConversationsByUserSuccess should give a user id to the method
// and successfully receive a list of conversations
func TestGetConversationsByUserSuccess(t *testing.T) {
	var mockResult []Conversation
	mockResult = append(mockResult, Conversation{ID: "1", URL: "www.weeee.com", MessagesURL: "layer:///messages/sdfkjasdlfkj", Participants: []string{"A", "B"}})
	mockResult = append(mockResult, Conversation{ID: "2", URL: "localhost", MessagesURL: "layer:///messages/sdkfjlskdjflfkj", Participants: []string{"B", "C"}})

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// mock to list out the articles
	httpmock.RegisterResponder("GET", "https://api.layer.com/apps/123/users/B/conversations",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, mockResult)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	l := New("123", "fjghfjshryfbus", "1.0")
	convos, err := l.GetConversationsByUser("B")
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if !reflect.DeepEqual(convos, mockResult) {
		t.Log("Handled response is different that response given...")
		t.Logf("%+v\n", convos)
		t.Fail()
	}

	t.Log("Success!")
	return
}
