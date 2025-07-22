package models

import (
	"database/sql"
	"errors"
	"log/slog"
	"os"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippedModel struct {
	DB *sql.DB
}

func (m *SnippedModel) Insert(title string, content string, expires int) (int, error) {

	return 0, nil
}

func (m *SnippedModel) Get(id int) (Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippetbox.snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	var s Snippet

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		}

		return Snippet{}, err
	}
	slog.New(slog.NewJSONHandler(os.Stdout, nil)).Error("%s", s.Title)
	return s, nil
}

func (m *SnippedModel) Latest() ([]Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippetbox.snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY  id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippets []Snippet
	for rows.Next() {
		var s Snippet

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	return snippets, nil
}
