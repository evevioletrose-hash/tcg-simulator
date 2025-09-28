package internal

// Direction represents the four cardinal directions for movement and room exits.
// Rooms can have exits in multiple directions, creating complex movement paths
// and strategic positioning opportunities.
type Direction int

const (
	North Direction = iota // Toward row 0 (typically toward Player 1's home)
	East                   // Toward higher column numbers (rightward)
	South                  // Toward higher row numbers (typically toward Player 2's home)
	West                   // Toward lower column numbers (leftward)
)

// Field represents the game board as a spatial grid of rooms.
// The field serves as the primary tactical space where:
//   - Room control determines expansion opportunities
//   - Spatial positioning affects combat and movement
//   - Lane control through footprints creates strategic depth
type Field struct {
	Grid       [][]*Room // 2D grid of rooms [row][col]
	Rows, Cols int       // Board dimensions
}

// NewField creates a new game board with the specified dimensions.
// Each position is initialized with a neutral room at coordinates (row, col).
// 
// Common configurations:
//   - 5x3: Standard tactical board with clear lanes
//   - 7x5: Larger strategic board for complex games
//   - 3x3: Quick tactical encounters
//
// All rooms start neutral (ControlledBy: -1) and must be claimed through gameplay.
func NewField(rows, cols int) *Field {
	grid := make([][]*Room, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]*Room, cols)
		for j := 0; j < cols; j++ {
			grid[i][j] = &Room{
				Row:          i,
				Col:          j,
				ControlledBy: -1, // -1 = neutral, 0/1 = player index
			}
		}
	}
	return &Field{
		Grid: grid,
		Rows: rows,
		Cols: cols,
	}
}

// Room represents a single location on the game board.
// Rooms are the fundamental unit of spatial control and resource management.
//
// Key mechanics:
//   - RPV (Room Point Value) counts against Universe Capacity
//   - RS (Room Space) = RPV, determining internal capacity
//   - Object Size limits what can occupy the room
//   - Footprints block adjacent expansion for lane control
//   - Exits enable movement and expansion pathways
type Room struct {
	Row, Col int // Grid coordinates (0-indexed)

	// Resource System
	RPV int // Room Point Value: counts against Universe Capacity AND determines Room Space
	RS  int // Room Space: internal capacity for occupants (typically equal to RPV)

	// Spatial Control
	Footprints []int // Blocked adjacent columns on this row for lane control
	Exits      []Direction

	// Occupancy
	Occupants    []*Card // Cards currently in this room
	ControlledBy int     // -1 = neutral, 0/1 = player index
}

// GetTotalObjectSize calculates the total Object Size of all occupants in this room.
// This is used to validate placement against Room Space limits.
//
// Returns the sum of all occupant OS values.
func (r *Room) GetTotalObjectSize() int {
	total := 0
	for _, card := range r.Occupants {
		total += card.OS
	}
	return total
}

// CanAccommodate checks if a card can be placed in this room.
// Placement is valid if adding the card's Object Size wouldn't exceed Room Space.
//
// Parameters:
//   - card: The card to potentially place
//
// Returns true if the card can fit, false otherwise.
func (r *Room) CanAccommodate(card *Card) bool {
	return r.GetTotalObjectSize()+card.OS <= r.RS
}

// IsControlledBy checks if this room is controlled by the specified player.
//
// Parameters:
//   - playerIndex: Player index to check (0 or 1)
//
// Returns true if the room is controlled by the player, false if neutral or enemy-controlled.
func (r *Room) IsControlledBy(playerIndex int) bool {
	return r.ControlledBy == playerIndex
}

// IsNeutral checks if this room is currently uncontrolled.
//
// Returns true if the room is neutral (-1), false if controlled by any player.
func (r *Room) IsNeutral() bool {
	return r.ControlledBy == -1
}

// HasExit checks if this room has an exit in the specified direction.
//
// Parameters:
//   - dir: Direction to check for an exit
//
// Returns true if an exit exists in that direction.
func (r *Room) HasExit(dir Direction) bool {
	for _, exit := range r.Exits {
		if exit == dir {
			return true
		}
	}
	return false
}

// BlocksColumn checks if this room's footprint blocks the specified column.
// Footprints prevent opponent expansion and maintain lane integrity.
//
// Parameters:
//   - col: Column index to check
//
// Returns true if this room blocks that column on its row.
func (r *Room) BlocksColumn(col int) bool {
	for _, blockedCol := range r.Footprints {
		if blockedCol == col {
			return true
		}
	}
	return false
}
