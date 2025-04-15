package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippets/{id}", app.handleGetSnippetById)
	mux.HandleFunc("POST /snippets", app.handlePostSnippets)

	return mux
}
