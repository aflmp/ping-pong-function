package function

import (
	"fmt"
	"net/http"
)

// Ping is an http handler that responds with pong message
func Ping(w http.ResponseWriter, r *http.Request) {
	validRequest := r.URL.Path == "/ping" && r.Method == http.MethodGet
	if !validRequest {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	w.Header().Add("content-type", "application/json")
	fmt.Fprintln(w, `{"message": "pong"}`)
}
