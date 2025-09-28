// Package internal implements the core game mechanics for the TCG Simulator.
// This package provides the foundation for a strategic trading card game featuring
// spatial board control, resource management, and dice-based combat resolution.
package internal

import (
	"fmt"
	"math/rand"
)

// Game represents the main game state and orchestrates all game mechanics.
// The Game manages the shared Universe Capacity (UC), field state, and player interactions.
type Game struct {
	Field   *Field    // The game board containing all rooms and spatial relationships
	Players []*Player // All players in the game (typically 2)
	UC      int       // Universe Capacity: shared point pool limiting total room complexity
}

// NewGame creates a new game instance with default settings.
// Default configuration:
//   - 5x3 game board (5 rows, 3 columns)
//   - 2 players with home rows at opposite ends (rows 0 and 4)
//   - Universe Capacity of 20 points
//
// The UC limit of 20 forces strategic decisions about room deployment,
// as each room's RPV (Room Point Value) counts against this shared pool.
func NewGame() *Game {
	// Example: 5x3 board, UC 20
	field := NewField(5, 3)
	players := []*Player{
		NewPlayer("Player 1", 0), // Home row 0 (top)
		NewPlayer("Player 2", 4), // Home row 4 (bottom)
	}
	return &Game{
		Field:   field,
		Players: players,
		UC:      20, // Universe Capacity: shared resource pool
	}
}

// Run executes the main game loop and handles player interactions.
// Currently provides basic initialization; will be expanded to include:
//   - Turn management and phase progression
//   - Player input handling and command processing  
//   - Field rendering and game state display
//   - Win condition checking and game resolution
func (g *Game) Run() {
	fmt.Println("Game start!")
	// TODO: Add game loop and CLI input handling here
	// e.g., drawing the field, prompting for commands, resolving actions
}

// RollD20 generates a random d20 roll (1-20).
// This is a utility function for simple dice rolls.
// For contested actions, use OpposedD20() which provides full resolution context.
func RollD20() int {
	return rand.Intn(20) + 1
}
