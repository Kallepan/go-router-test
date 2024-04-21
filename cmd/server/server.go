package main

import (
	"net/http"

	v1 "github.com/kallepan/go-router-test/api/v1"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/hello", v1.HelloHandler)
	mux.HandleFunc("GET /v1/goodbye", v1.GoodbyeHandler)
	mux.HandleFunc("POST /v1/create", v1.CreateHandler)

	http.ListenAndServe(":8080", mux)
}
