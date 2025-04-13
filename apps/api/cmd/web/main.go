package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)

	mux.HandleFunc("GET /snippets/{id}", handleGetSnippetById)
	mux.HandleFunc("POST /snippets", handlePostSnippets)

	log.Print("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
