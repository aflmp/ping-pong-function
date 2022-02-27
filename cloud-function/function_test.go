package function

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPing(t *testing.T) {
	tests := []struct {
		name               string
		method             string
		endpoint           string
		responseStatusCode int
		responseBody       string
	}{
		{
			name:               "test-success",
			method:             "GET",
			endpoint:           "/ping",
			responseStatusCode: 200,
			responseBody:       `{"message": "pong"}`,
		},
		{
			name:               "test-failure-1",
			method:             "GET",
			endpoint:           "/blah",
			responseStatusCode: 400,
			responseBody:       "invalid request",
		},
		{
			name:               "test-failure-2",
			method:             "POST",
			endpoint:           "/ping",
			responseStatusCode: 400,
			responseBody:       "invalid request",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, err := http.NewRequest(test.method, test.endpoint, nil)
			if err != nil {
				t.Fatal(err)
			}

			handler := http.HandlerFunc(Ping)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, request)

			if w.Result().StatusCode != test.responseStatusCode {
				t.Errorf("status code mismatch, got: %v; want: %v", w.Code, test.responseStatusCode)
			}

			body, err := ioutil.ReadAll(w.Result().Body)
			if err != nil {
				t.Fatal(err)
			}
			defer w.Result().Body.Close()

			if strings.TrimSpace(string(body)) != test.responseBody {
				t.Errorf("response body mismatch, got: %v; want: %v", string(body), test.responseBody)
			}
		})
	}
}
