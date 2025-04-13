package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func getSnippetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Brother, I need a proper integer id; also positive, brother."))
		return
	}

	msg := fmt.Sprintf("Here you desired snippet with ID %d, brother.", id)
	w.Write([]byte(msg))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("GET /snippets/{id}", getSnippetById)

	log.Print("Starting sever on :4000")
	log.Print("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
