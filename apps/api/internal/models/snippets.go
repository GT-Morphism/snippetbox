package models

import (
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
	return 0, nil
}

func (m *SnippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
