package unit_tests_3

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
	"fmt"

)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestDoStuffWithRoundTripper(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		// equals(t, req.URL.String(), "http://example.com/some/path")
		return &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
 			// Must be set to non-nil value or it panics
			Header:     make(http.Header),
		}
	})

	api := API{client, "http://example.com"}
	body, err := api.DoStuff()
	fmt.Println(body, err)
	// ok(t, err)
	// equals(t, []byte("OK"), body)

}