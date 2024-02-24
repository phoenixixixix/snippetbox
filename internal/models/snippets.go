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
	return 0, nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
