package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Snippet struct {
	ID         int
	Title      string
	Content    string
	Created_at time.Time
	Expires_at time.Time
}

type SnippetModel struct {
	DB *pgxpool.Pool
}

func (m *SnippetModel) Insert(title, content string, expires_at int) (int, error) {
	query := `
    INSERT INTO snippets (title, content, expires_at)
    VALUES ($1, $2, CURRENT_TIMESTAMP + INTERVAL '1 day' * $3)
    RETURNING id;`

	var lastId int

	err := m.DB.QueryRow(context.Background(), query, title, content, expires_at).Scan(&lastId)
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (m *SnippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
