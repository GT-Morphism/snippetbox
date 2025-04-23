package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"snippetbox.gentiluomo.dev/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Write([]byte("Hello from Snippetbox"))
}

func (app *application) handleGetSnippets(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"snippets": snippets}, nil)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) handleGetSnippetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		app.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Brother, I need a proper integer id; also positive, brother."))
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	fmt.Fprintf(w, "Here you desired snippet with ID %d, brother: %+v", id, snippet)
}

func (app *application) handlePostSnippets(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires_at := 7

	id, err := app.snippets.Insert(title, content, expires_at)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippets/%d", id), http.StatusSeeOther)
}
