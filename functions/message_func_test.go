package functions

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"mysite.com/myfunctions/common/writers"
)

func TestOutputMessage(t *testing.T) {
	message := "Hello this is a test message"
	encodedMessage := url.QueryEscape(message)
	r, err := http.NewRequest("GET", fmt.Sprintf("/?message=%s", encodedMessage), nil)
	if err != nil {
		t.Fatal(err)
	}
	// execute the function
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(OutputMessage)
	handler.ServeHTTP(w, r)
	resp := w.Result()
	// check we got a valid response code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", resp.StatusCode, http.StatusOK)
	}
	// check there was no error getting the body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	// check the response body
	stringBody := string(body)
	jw := writers.NewMessageWriter(message)
	messageString, _ := jw.JSONString()
	if stringBody != messageString {
		t.Errorf("wrong response body: got %v want %v", body, "Hello, World!\n")
	}
}
