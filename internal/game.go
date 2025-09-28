package internal

import (
	"fmt"
	"math/rand"
)

type Game struct {
	Field   *Field
	Players []*Player
	UC      int // Universe Capacity
}

func NewGame() *Game {
	// Example: 5x3 board, UC 20
	field := NewField(5, 3)
	players := []*Player{
		NewPlayer("Player 1", 0),
		NewPlayer("Player 2", 4),
	}
	return &Game{
		Field:   field,
		Players: players,
		UC:      20,
	}
}

func (g *Game) Run() {
	fmt.Println("Game start!")
	// Add game loop and CLI input handling here
	// e.g., drawing the field, prompting for commands, resolving actions
}

func RollD20() int {
	return rand.Intn(20) + 1
}
