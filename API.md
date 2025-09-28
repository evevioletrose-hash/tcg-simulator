# API Documentation

This document provides detailed API documentation for the TCG Simulator's internal package.

## Core Types

### Game

The `Game` struct orchestrates all game mechanics and maintains global state.

```go
type Game struct {
    Field   *Field    // The game board
    Players []*Player // All players in the game
    UC      int       // Universe Capacity: shared resource pool
}
```

#### Constructor

**`NewGame() *Game`**

Creates a new game with default configuration:
- 5x3 board (5 rows, 3 columns)
- 2 players with home rows at positions 0 and 4
- Universe Capacity of 20 points

#### Methods

**`Run()`**

Executes the main game loop. Currently provides basic initialization; expansion planned for:
- Turn management and phase progression
- Player input handling and command processing
- Field rendering and game state display
- Win condition checking and game resolution

### Field

The `Field` struct represents the spatial game board as a grid of rooms.

```go
type Field struct {
    Grid       [][]*Room // 2D grid of rooms [row][col]
    Rows, Cols int       // Board dimensions
}
```

#### Constructor

**`NewField(rows, cols int) *Field`**

Creates a new game board with specified dimensions. All rooms start neutral (ControlledBy: -1).

**Parameters:**
- `rows`: Number of rows in the grid
- `cols`: Number of columns in the grid

**Common Configurations:**
- `NewField(5, 3)`: Standard tactical board
- `NewField(7, 5)`: Large strategic board  
- `NewField(3, 3)`: Quick encounter board

### Room

The `Room` struct represents a single location on the board with tactical and resource properties.

```go
type Room struct {
    Row, Col     int         // Grid coordinates
    RPV          int         // Room Point Value (counts against UC)
    RS           int         // Room Space (internal capacity)
    Footprints   []int       // Blocked adjacent columns
    Exits        []Direction // Valid movement directions
    Occupants    []*Card     // Cards in this room
    ControlledBy int         // -1=neutral, 0/1=player index
}
```

#### Methods

**`GetTotalObjectSize() int`**

Returns the sum of Object Size values for all occupants in the room.

**`CanAccommodate(card *Card) bool`**

Checks if a card can be placed without exceeding Room Space limits.

**Parameters:**
- `card`: The card to potentially place

**Returns:** `true` if the card fits, `false` otherwise

**`IsControlledBy(playerIndex int) bool`**

Checks if the room is controlled by the specified player.

**`IsNeutral() bool`**

Returns `true` if the room is uncontrolled (ControlledBy == -1).

**`HasExit(dir Direction) bool`**

Checks if the room has an exit in the specified direction.

**`BlocksColumn(col int) bool`**

Returns `true` if the room's footprint blocks the specified column.

### Card

The `Card` struct represents game pieces with stats, abilities, and tactical properties.

```go
type Card struct {
    ID         string      // Unique identifier
    Name       string      // Display name
    OS         int         // Object Size (space consumption)
    Archetypes []Archetype // Combat specializations
    Tags       []string    // Flexible categorization
    Effects    []Effect    // Special abilities
    Owner      int         // Player index (0 or 1)
}
```

#### Methods

**`HasArchetype(archetype Archetype) bool`**

Checks if the card has the specified archetype.

**`HasTag(tag string) bool`**

Check if the card has the specified tag.

**`GetCombatModifier(opponent *Card, room *Room) int`**

Calculates combat bonuses based on archetypes and context. Currently provides basic Soldier bonus; will be expanded for:
- Room-based bonuses
- Archetype interactions
- Equipment effects

**`IsOwnedBy(playerIndex int) bool`**

Returns `true` if the card is owned by the specified player.

### Player

The `Player` struct represents a game participant with resources and positioning.

```go
type Player struct {
    Name    string  // Display name
    HomeRow int     // Spawn row for units
    Hand    []*Card // Available cards
}
```

#### Constructor

**`NewPlayer(name string, homeRow int) *Player`**

Creates a new player with specified name and home row position.

#### Methods

**`AddToHand(card *Card)`**

Adds a card to the player's hand.

**`RemoveFromHand(card *Card) bool`**

Removes a specific card from the hand. Returns `true` if found and removed.

**`HasCardInHand(card *Card) bool`**

Checks if the player has the specified card in hand.

**`GetHandSize() int`**

Returns the current number of cards in hand.

**`GetControlledRooms(field *Field, playerIndex int) []*Room`**

Returns all rooms on the field controlled by this player.

**`CanSpawnInHomeRow(field *Field, card *Card) *Room`**

Returns the first available room in the home row that can accommodate the card, or `nil` if none available.

## Resolution System

### DiceResult

Represents the outcome of a d20 roll in contested actions.

```go
type DiceResult struct {
    Roll   int    // Final roll result (d20 + modifiers)
    Band   string // Success level: "FAIL", "OK", "CRIT"
    Margin int    // Difference vs opponent (can be negative)
}
```

### Resolution Functions

**`OpposedD20(modA, modB int) (DiceResult, DiceResult)`**

Performs contested d20 rolls between two participants.

**Parameters:**
- `modA`: Modifier for participant A
- `modB`: Modifier for participant B

**Returns:**
- `DiceResult` for participant A
- `DiceResult` for participant B (with negative margin)

**Resolution Bands:**
- **FAIL**: Margin < 2 (no significant effect)
- **OK**: Margin ≥ 2 (basic success)
- **CRIT**: Margin ≥ 2 AND natural 20 (enhanced success)

**`ResolveCombat(attacker, defender *Card, room *Room) (DiceResult, DiceResult, bool)`**

Handles combat between two cards using opposed d20 resolution.

**Parameters:**
- `attacker`: The initiating card
- `defender`: The target card  
- `room`: The room where combat occurs

**Returns:**
- `DiceResult` for attacker
- `DiceResult` for defender
- `bool` indicating if defender was eliminated

**`RollD20() int`**

Simple utility function for basic d20 rolls (1-20). For contested actions, use `OpposedD20()`.

## Constants and Enums

### Archetype

```go
type Archetype string

const (
    Soldier Archetype = "Soldier" // Steady combat specialist
    Mage    Archetype = "Mage"    // Magical damage dealer
    Rogue   Archetype = "Rogue"   // Tactical manipulator
)
```

### Direction

```go
type Direction int

const (
    North Direction = iota // Toward row 0
    East                   // Toward higher columns
    South                  // Toward higher rows
    West                   // Toward lower columns
)
```

## Usage Examples

### Basic Game Setup

```go
// Create a new game
game := NewGame()

// Access the 5x3 field
field := game.Field

// Get players
player1 := game.Players[0] // Home row 0
player2 := game.Players[1] // Home row 4

// Check Universe Capacity
fmt.Printf("Available UC: %d\n", game.UC) // Prints: 20
```

### Room Management

```go
// Create a room with tactical properties
room := &Room{
    Row:          2,
    Col:          1,
    RPV:          4,           // Consumes 4 UC
    RS:           4,           // Provides 4 room space
    Footprints:   []int{0, 2}, // Blocks columns 0 and 2
    Exits:        []Direction{North, South, East},
    ControlledBy: 0,           // Controlled by player 0
}

// Check capacity
soldier := &Card{OS: 2}
scout := &Card{OS: 1}

if room.CanAccommodate(soldier) {
    room.Occupants = append(room.Occupants, soldier)
}
if room.CanAccommodate(scout) {
    room.Occupants = append(room.Occupants, scout)
}

// Total OS: 3, Room Space: 4, so both fit
```

### Combat Resolution

```go
// Create combatants
attacker := &Card{
    Archetypes: []Archetype{Soldier},
    OS:         2,
}

defender := &Card{
    Archetypes: []Archetype{Mage},
    OS:         1,
}

// Resolve combat
attackResult, defendResult, eliminated := ResolveCombat(attacker, defender, room)

if attackResult.Margin >= 2 {
    fmt.Printf("Attacker wins with %s result!\n", attackResult.Band)
    if eliminated {
        fmt.Println("Defender eliminated!")
    }
}
```

### Player Operations

```go
player := NewPlayer("Alice", 0)

// Add cards to hand
card1 := &Card{Name: "Elite Soldier", OS: 2}
card2 := &Card{Name: "Scout", OS: 1}

player.AddToHand(card1)
player.AddToHand(card2)

// Check spawn opportunities
if spawnRoom := player.CanSpawnInHomeRow(field, card1); spawnRoom != nil {
    fmt.Printf("Can spawn at (%d, %d)\n", spawnRoom.Row, spawnRoom.Col)
}
```

## Extension Points

The API is designed for extensibility in several areas:

### Effect System

The `Effect` struct is currently minimal but designed for expansion:

```go
type Effect struct {
    Description string
    // Future: Type, Timing, Target, Modifiers, etc.
}
```

### Room Customization

Rooms can be extended with:
- Tag systems for archetype bonuses
- Environmental effects and hazards
- Dynamic properties that change during gameplay

### Combat System

The combat resolution can be enhanced with:
- Equipment and enhancement modifiers
- Chained reactions and triggered abilities
- Environmental factors and positioning bonuses

This API provides a solid foundation for tactical gameplay while remaining extensible for future enhancements.