package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Write([]byte("Hello from Snippetbox"))
}

func (app *application) handleGetSnippetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		app.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Brother, I need a proper integer id; also positive, brother."))
		return
	}

	fmt.Fprintf(w, "Here you desired snippet with ID %d, brother.", id)
}

func (app *application) handlePostSnippets(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet..."))
}
