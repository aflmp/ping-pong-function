package main

import (
	"log"
	"net/http"

	function "github.com/aflmp/ping-pong-function/cloud-function"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", function.Ping)
	log.Println("server listening on 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
