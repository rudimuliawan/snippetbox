package models

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (s *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := s.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *SnippetModel) Get(id int) (Snippet, error) {
	stmt := `SELECT id, title, content, created, expires 
	FROM snippets WHERE id = ? AND expires > UTC_TIMESTAMP()`

	row := s.DB.QueryRow(stmt, id)

	var snippet Snippet

	err := row.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNotFound
		} else {
			return Snippet{}, err
		}
	}

	return snippet, nil
}

func (s *SnippetModel) Latest() ([]Snippet, error) {
	stmt := `SELECT * FROM snippets WHERE expires > UTC_TIMESTAMP() ORDER BY created`

	rows, err := s.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	var snippets []Snippet

	for rows.Next() {
		var snippet Snippet
		err := rows.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, snippet)
	}

	return snippets, nil
}
