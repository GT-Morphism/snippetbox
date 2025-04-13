package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Write([]byte("Hello from Snippetbox"))
}

func handleGetSnippetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Brother, I need a proper integer id; also positive, brother."))
		return
	}

	fmt.Fprintf(w, "Here you desired snippet with ID %d, brother.", id)
}

func handlePostSnippets(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet..."))
}
