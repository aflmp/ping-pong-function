package function

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPing(t *testing.T) {
	tests := map[string]struct {
		method             string
		endpoint           string
		responseStatusCode int
		responseBody       string
	}{
		"successful test case": {method: "GET", endpoint: "/ping", responseStatusCode: 200, responseBody: `{"message": "pong"}`},
		"failure test case 1":  {method: "GET", endpoint: "/blah", responseStatusCode: 400, responseBody: "invalid request"},
		"failure test case 2":  {method: "POST", endpoint: "/ping", responseStatusCode: 400, responseBody: "invalid request"},
	}

	for testName, test := range tests {
		t.Logf("running: %v", testName)
		request, err := http.NewRequest(test.method, test.endpoint, nil)
		if err != nil {
			t.Fatal(err)
		}

		handler := http.HandlerFunc(Ping)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, request)

		if w.Result().StatusCode != test.responseStatusCode {
			t.Errorf("status code mis-match, got: %v; want: %v", w.Code, test.responseStatusCode)
		}

		body, err := ioutil.ReadAll(w.Result().Body)
		if err != nil {
			t.Fatal(err)
		}
		defer w.Result().Body.Close()

		if strings.TrimSpace(string(body)) != test.responseBody {
			t.Errorf("response body mis-match, got: %v; want: %v", string(body), test.responseBody)
		}
	}
}
