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

// package function

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
// )

// // count is a global variable, but only shared within a function instance.
// var count = 0

// func init() {
// 	functions.HTTP("ExecutionCount", ExecutionCount)
// }

// // ExecutionCount is an HTTP Cloud Function that counts how many times it
// // is executed within a specific instance.
// func ExecutionCount(w http.ResponseWriter, r *http.Request) {
// 	count++

// 	// Note: the total function invocation count across
// 	// all instances may not be equal to this value!
// 	fmt.Fprintf(w, "Instance execution count: %d", count)
// }
