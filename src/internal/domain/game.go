package domain

import "github.com/google/uuid"

type Game struct {
	ID    string
	Board Board
}

func NewGame() *Game {
	return &Game{
		ID:    uuid.New().String(),
		Board: Board{},
	}
}
