package main

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)

type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	api := humago.New(mux, huma.DefaultConfig("Snippetbox API", "0.0.0"))

	huma.Register(api, huma.Operation{
		OperationID: "get-snippets",
		Method:      http.MethodGet,
		Path:        "/snippets",
		Summary:     "Get last 10 snippets",
		Description: "This endpoint provides the last 10 snippets sorted desc by id",
		Tags:        []string{"Snippets"},
	}, app.handleGetSnippets)

	huma.Register(api, huma.Operation{
		OperationID:   "post-snippets",
		Method:        http.MethodPost,
		Path:          "/snippets",
		Summary:       "Create new snippet",
		Description:   "This endpoint allows for the creation of a new snippet",
		Tags:          []string{"Snippets"},
		DefaultStatus: http.StatusSeeOther,
	}, app.handlePostSnippets)

	huma.Register(api, huma.Operation{
		OperationID: "get-snippet-by-id",
		Method:      http.MethodGet,
		Path:        "/snippets/{id}",
		Summary:     "Get snippet by id",
		Description: "This endpoint provides details of a snippet given its id",
		Tags:        []string{"Snippets"},
	}, app.handleGetSnippetById)

	return mux
}
