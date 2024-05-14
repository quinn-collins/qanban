package models

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type CardModelInterface interface {
	Insert(title, content string) (int, error)
	Get(id int) (*Card, error)
	GetAll() ([]*Card, error)
}

type Card struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CardModel struct {
	DB *pgxpool.Pool
}

func (m *CardModel) Insert(title string, content string) (int, error) {
	stmt := `INSERT INTO cards (title, content) VALUES ($1, $2) RETURNING id;`

	var id int
	err := m.DB.QueryRow(context.Background(), stmt, title, content).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *CardModel) Get(id int) (*Card, error) {
	stmt := `SELECT id, title, content, created_at, updated_at FROM cards WHERE id = $1;`

	var c Card
	err := m.DB.QueryRow(context.Background(), stmt, id).Scan(&c.ID, &c.Title, &c.Content, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &c, nil
}

func (m *CardModel) GetAll() ([]*Card, error) {
	stmt := `SELECT id, title, content, created_at, updated_at FROM cards;`

	var cs []*Card
	rows, err := m.DB.Query(context.Background(), stmt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	for rows.Next() {
		var c Card
		err = rows.Scan(&c.ID, &c.Title, &c.Content, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return nil, err
		}

		cs = append(cs, &c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cs, nil
}
