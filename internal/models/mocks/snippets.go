package mocks

import (
	"time"

	"phoenixixixix/snippetbox/internal/models"
)

var mockSnippet = &models.Snippet{
	ID:        1,
	Title:     "Snippet Mock",
	Content:   "Content of mocked snippet",
	CreatedAt: time.Now(),
	ExpiresAt: time.Now(),
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 2, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}