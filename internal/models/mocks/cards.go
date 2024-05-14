package mocks

import (
	"github.com/quinn-collins/qanban/internal/models"
	"time"
)

var mockCard = &models.Card{
	ID:        1,
	Title:     "An old silent pond",
	Content:   "An old silent pond...",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

type CardModel struct{}

func (m *CardModel) Insert(title, content string) (int, error) {
	return 2, nil
}

func (m *CardModel) Get(id int) (*models.Card, error) {
	switch id {
	case 1:
		return mockCard, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *CardModel) GetAll() ([]*models.Card, error) {
	return []*models.Card{mockCard}, nil
}
