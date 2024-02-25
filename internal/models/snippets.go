package models

import (
	"database/sql"
	"errors"
)

type Snippet struct {
	ID        int
	Title     string
	Content   string
	CreatedAt string
	ExpiresAt string
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expiresAt int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created_at, expires_at)
  VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expiresAt)
	if err != nil {
		return 0, nil
	}

	id, err := result.LastInsertId() // LastInsertId() method is not supported by PostgreSQL
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created_at, expires_at FROM snippets
  WHERE expires_at > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.CreatedAt, &s.ExpiresAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
