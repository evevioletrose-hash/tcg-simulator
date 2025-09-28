package internal

// Player represents a participant in the game with their resources and positioning.
// Players manage cards, control territory, and compete for board dominance.
//
// Key mechanics:
//   - Home Row: Starting position where new units spawn
//   - Hand Management: Cards available for play
//   - Territorial Control: Rooms controlled provide expansion opportunities
type Player struct {
	Name    string  // Display name for the player
	HomeRow int     // Row index where this player's units spawn (0-4 typically)
	Hand    []*Card // Cards currently available to play
	// TODO: add deck, discard, resources, etc.
}

// NewPlayer creates a new player with the specified name and home row position.
// 
// Home Row Strategy:
//   - Row 0: "Northern" player, typically Player 1
//   - Row 4: "Southern" player, typically Player 2  
//   - Units spawn in home row, must move forward to engage
//   - Forward movement eliminates "summoning sickness"
//
// Parameters:
//   - name: Display name for the player
//   - homeRow: Row index where this player's units will spawn
//
// Returns a new Player instance ready for gameplay.
func NewPlayer(name string, homeRow int) *Player {
	return &Player{
		Name:    name,
		HomeRow: homeRow,
		Hand:    []*Card{}, // Start with an empty hand
	}
}

// AddToHand adds a card to this player's hand.
// Used for drawing cards, gaining cards from effects, etc.
//
// Parameters:
//   - card: The card to add to the hand
func (p *Player) AddToHand(card *Card) {
	p.Hand = append(p.Hand, card)
}

// RemoveFromHand removes a specific card from this player's hand.
// Used when playing cards, discarding, or losing cards to effects.
//
// Parameters:
//   - card: The card to remove from the hand
//
// Returns true if the card was found and removed, false otherwise.
func (p *Player) RemoveFromHand(card *Card) bool {
	for i, handCard := range p.Hand {
		if handCard == card {
			// Remove card by shifting slice
			p.Hand = append(p.Hand[:i], p.Hand[i+1:]...)
			return true
		}
	}
	return false
}

// HasCardInHand checks if the player has a specific card in their hand.
//
// Parameters:
//   - card: The card to look for
//
// Returns true if the card is in the player's hand.
func (p *Player) HasCardInHand(card *Card) bool {
	for _, handCard := range p.Hand {
		if handCard == card {
			return true
		}
	}
	return false
}

// GetHandSize returns the current number of cards in the player's hand.
//
// Returns the count of cards in hand.
func (p *Player) GetHandSize() int {
	return len(p.Hand)
}

// GetControlledRooms finds all rooms on the field controlled by this player.
// Used for calculating territorial advantage, expansion opportunities, etc.
//
// Parameters:
//   - field: The game field to search
//   - playerIndex: This player's index (0 or 1)
//
// Returns a slice of all rooms controlled by this player.
func (p *Player) GetControlledRooms(field *Field, playerIndex int) []*Room {
	var controlled []*Room
	
	for row := 0; row < field.Rows; row++ {
		for col := 0; col < field.Cols; col++ {
			room := field.Grid[row][col]
			if room.IsControlledBy(playerIndex) {
				controlled = append(controlled, room)
			}
		}
	}
	
	return controlled
}

// CanSpawnInHomeRow checks if this player can place a unit in their home row.
// Validates room availability and capacity in the home row.
//
// Parameters:
//   - field: The game field to check
//   - card: The card to potentially spawn
//
// Returns the first available room in home row, or nil if none available.
func (p *Player) CanSpawnInHomeRow(field *Field, card *Card) *Room {
	if p.HomeRow < 0 || p.HomeRow >= field.Rows {
		return nil // Invalid home row
	}
	
	// Check each column in the home row
	for col := 0; col < field.Cols; col++ {
		room := field.Grid[p.HomeRow][col]
		if room.CanAccommodate(card) {
			return room
		}
	}
	
	return nil // No available space in home row
}
