package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMuxHandler(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(write http.ResponseWriter, request *http.Request) {
		fmt.Fprint(write, "Hello, world!")
	})

	mux.HandleFunc("/about", func(write http.ResponseWriter, request *http.Request) {
		fmt.Fprint(write, "this testing with mux handler")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
