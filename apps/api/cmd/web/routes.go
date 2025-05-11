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

	apiConfig := huma.DefaultConfig("Snippetbox API", "0.0.0")
	apiConfig.Info.Contact = &huma.Contact{
		Name:  "Giuseppe",
		Email: "don@example.com",
	}
	apiConfig.Info.License = &huma.License{
		Name:       "MIT License",
		Identifier: "MIT",
		URL:        "https://mit-license.org/",
	}
	apiConfig.Info.Description = `## Purpose
The documentation has two goals:
- Let new developers quickly get a high-level overview of the project's services
- Let developers be forced to provide proper documentation
`

	api := humago.New(mux, apiConfig)
	api.UseMiddleware(commonHeaders, app.logRequest, recoverPanic(api))

	huma.Register(api, huma.Operation{
		OperationID: "get-api-health-status",
		Method:      http.MethodGet,
		Path:        "/health",
		Summary:     "Show health status of API",
		Description: "This endpoint allows the retrieval of the API's current health status",
	}, app.handleGetApiHealthStatus)

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
		DefaultStatus: http.StatusCreated,
	}, app.handlePostSnippets)

	huma.Register(api, huma.Operation{
		OperationID: "get-snippet-by-id",
		Method:      http.MethodGet,
		Path:        "/snippets/{id}",
		Summary:     "Get snippet by id",
		Description: "This endpoint provides details of a snippet given its id",
		Tags:        []string{"Snippets"},
	}, app.handleGetSnippetById)

	huma.Register(api, huma.Operation{
		OperationID: "post-users",
		Method:      http.MethodPost,
		Path:        "/users",
		Summary:     "Create new user",
		Description: "This endpoint allows to sign up new users",
		Tags:        []string{"Users"},
	}, app.handlePostUsers)

	huma.Register(api, huma.Operation{
		OperationID: "auth-login",
		Method:      http.MethodPost,
		Path:        "/auth/login",
		Summary:     "Authenticate and login the user",
		Description: "This endpoint allows registered users to login",
		Tags:        []string{"Authentication"},
	}, app.handleAuthLogin)

	huma.Register(api, huma.Operation{
		OperationID: "auth-logout",
		Method:      http.MethodPost,
		Path:        "/auth/logout",
		Summary:     "Logout the user",
		Description: "This endpoint allows authenticated users to log out",
		Tags:        []string{"Authentication"},
	}, app.handleAuthLogout)

	return mux
}
