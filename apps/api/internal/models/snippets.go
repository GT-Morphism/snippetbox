package models

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
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
	query := `
		SELECT id, title, content, created_at, expires_at FROM snippets
		WHERE expires_at > CURRENT_TIMESTAMP AND id = $1;
	`

	var s Snippet

	err := m.DB.QueryRow(context.Background(), query, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created_at, &s.Expires_at)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	query := `
		SELECT id, title, content, created_at, expires_at FROM snippets
		WHERE expires_at > CURRENT_TIMESTAMP
		ORDER BY id DESC LIMIT 10
	`

	rows, err := m.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var snippets []Snippet

	for rows.Next() {
		var s Snippet

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created_at, &s.Expires_at)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
