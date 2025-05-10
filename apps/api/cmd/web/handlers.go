package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/redis/go-redis/v9"
	"snippetbox.gentiluomo.dev/internal/models"
)

type HandleGetApiHealthStatusOutput struct {
	Body struct {
		Message     string `json:"message"`
		DBStatus    string `json:"dbStatus"`
		CacheStatus string `json:"cacheStatus"`
	}
}

func (app *application) handleGetApiHealthStatus(ctx context.Context, input *struct{}) (*HandleGetApiHealthStatusOutput, error) {
	resp := &HandleGetApiHealthStatusOutput{}

	var greeting string
	err := app.db.QueryRow(ctx, "SELECT 'Hello, Sir.'").Scan(&greeting)
	if err != nil {
		app.logger.Error("QueryRow failed", "err", err)
		resp.Body.Message = "There is a problem with the database"
		resp.Body.DBStatus = "inactive"
		resp.Body.CacheStatus = "unchecked"

		return resp, huma.Error503ServiceUnavailable("The fault is ours, brother")
	}

	err = app.cache.Ping(ctx).Err()
	if err != nil {
		app.logger.Error("Unable to connect to redis", "err", err)
		resp.Body.Message = "There is a problem with the cache"
		resp.Body.DBStatus = "active"
		resp.Body.CacheStatus = "inactive"

		return resp, huma.Error503ServiceUnavailable("The fault is ours, brother")
	}

	resp.Body.Message = "All services are up and running"
	resp.Body.DBStatus = "active"
	resp.Body.CacheStatus = "active"

	return resp, nil
}

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
	Body             models.Snippet
	ShowCreatedToast http.Cookie `header:"Set-Cookie"`
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

	_, err = app.cache.Get(ctx, "show-snippet-created-toast").Result()

	if err == redis.Nil {
		app.logger.Info("No key `show-snippet-created-toast` found")
	} else if err != nil {
		app.logger.Error("Something went wrong trying to get value for key `show-snippet-created-toast`", "error", err)
	} else {
		app.cache.Del(ctx, "show-snippet-created-toast")
		resp.ShowCreatedToast = http.Cookie{
			Name:     "show-created-toast",
			Value:    "true",
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		}
	}

	resp.Body = snippet

	return resp, nil
}

type HandlePostSnippetsInput struct {
	Body struct {
		Title      string `json:"title" doc:"Title of snippet" maxLength:"100"`
		Content    string `json:"content" doc:"Content of snippet"`
		Expires_at int    `json:"expires_at" doc:"Number of days until the snippet expires" minimum:"1" maximum:"365" enum:"1,7,365"`
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

	err = app.cache.Set(ctx, "show-snippet-created-toast", "true", 5*time.Minute).Err()
	if err != nil {
		app.logger.Error("Something went wrong trying to store creation message for new snippet", "input", input, "error", err)
	}

	resp.Url = fmt.Sprintf("/snippets/%d", id)

	return resp, nil
}
