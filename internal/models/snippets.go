package models

import "database/sql"

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
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
