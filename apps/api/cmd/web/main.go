package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)

	mux.HandleFunc("GET /snippets/{id}", handleGetSnippetById)
	mux.HandleFunc("POST /snippets", handlePostSnippets)

	log.Printf("Starting server on %s", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
