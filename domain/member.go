package domain

import (
	"time"

	"github.com/google/uuid"
)

type Member struct {
	ID       uuid.UUID
	Name     string
	Email    string
	JoinedAt time.Time
}

func NewMember(name, email string) Member {
	return Member{
		ID:       uuid.New(),
		Name:     name,
		Email:    email,
		JoinedAt: time.Now(),
	}
}
