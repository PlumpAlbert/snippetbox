package models

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *pgxpool.Pool
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	query := fmt.Sprintf(
		`INSERT INTO snippets (title, content, created, expires)
		VALUES(@title, @content, NOW(), NOW() + INTERVAL '%d days')
		RETURNING id`,
		expires,
	)
	args := pgx.NamedArgs{
		"title":   title,
		"content": content,
	}

	row := m.DB.QueryRow(
		context.Background(),
		query,
		args,
	)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
