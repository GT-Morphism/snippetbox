package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/danielgtaylor/huma/v2"
	"snippetbox.gentiluomo.dev/internal/models"
)

type HandleGetSnippetsOutput struct {
	Body []models.Snippet
}

func (app *application) handleGetSnippets(ctx context.Context, input *struct{}) (*HandleGetSnippetsOutput, error) {
	resp := &HandleGetSnippetsOutput{}
	snippets, err := app.snippets.Latest()
	if err != nil {
		resp.Body = []models.Snippet{}
		app.logger.Error("Something went wrong trying to fetch snippets", "error", err)
		return resp, huma.Error500InternalServerError("The fault is ours, brother.")
	}

	resp.Body = snippets

	return resp, nil
}

type HandleGetSnippetByIdOutput struct {
	Body models.Snippet
}

func (app *application) handleGetSnippetById(ctx context.Context, input *struct {
	ID int `path:"id" minimum:"1" doc:"id of snippet to show"`
},
) (*HandleGetSnippetByIdOutput, error) {
	resp := &HandleGetSnippetByIdOutput{}

	snippet, err := app.snippets.Get(input.ID)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			return nil, huma.Error404NotFound("Brother, no sinppet for your id")
		} else {
			app.logger.Error("Something went wrong trying to fetch snippet by id", "id", input.ID, "error", err)
			return resp, huma.Error500InternalServerError("The fault is ours, brother.")
		}
	}

	resp.Body = snippet

	return resp, nil
}

type HandlePostSnippetsInput struct {
	Body struct {
		Title      string `json:"title" doc:"Title of snippet"`
		Content    string `json:"content" doc:"Content of snippet"`
		Expires_at int    `json:"expires_at" doc:"Number of days until the snippet expires" minimum:"0" maximum:"30"`
	}
}

type HandlePostSnippetsResponse struct {
	Url string `header:"Location" doc:"The URL of the newly created snippet to which the user will be redirected to"`
}

func (app *application) handlePostSnippets(ctx context.Context, input *HandlePostSnippetsInput) (*HandlePostSnippetsResponse, error) {
	resp := &HandlePostSnippetsResponse{}

	id, err := app.snippets.Insert(input.Body.Title, input.Body.Content, input.Body.Expires_at)
	if err != nil {
		app.logger.Error("Something went wrong trying to create a new snippet", "input", input, "error", err)
		return nil, huma.Error500InternalServerError("The fault is ours, brother.")
	}

	resp.Url = fmt.Sprintf("/snippets/%d", id)

	return resp, nil
}
