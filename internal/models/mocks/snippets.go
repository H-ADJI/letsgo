package mocks

import (
	"time"

	"github.com/H-ADJI/letsgo/internal/models"
)

var mockSnippet = models.Snippet{
	ID:      1,
	Title:   "A test title ",
	Content: "A test content ....",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnippetModel struct{}

func (s SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 2, nil
}
func (s SnippetModel) Get(id int) (models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return models.Snippet{}, models.ErrNoRecords
	}
}
func (s SnippetModel) Latest() ([]models.Snippet, error) {
	return []models.Snippet{mockSnippet}, nil
}
