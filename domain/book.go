package domain

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID
	Title     string
	Author    string
	ISBN      string
	CreatedAt time.Time
}

func NewBook(title, author, isbn string) Book {
	return Book{
		ID:        uuid.New(),
		Title:     title,
		Author:    author,
		ISBN:      isbn,
		CreatedAt: time.Now(),
	}
}
