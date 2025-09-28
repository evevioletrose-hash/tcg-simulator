# TCG Simulator

A strategic Trading Card Game simulator featuring spatial board control, resource management, and tactical combat using dice-based resolution.

## Table of Contents

- [Core Concepts](#core-concepts)
- [Resolution System](#resolution-system)
- [Archetypes](#archetypes)
- [Getting Started](#getting-started)
- [Game Components](#game-components)
- [Examples](#examples)

## Core Concepts

### Universe Capacity (UC)
**Universe Capacity** is the shared point pool that limits the total complexity of rooms active on the board. Each room's **Room Point Value (RPV)** counts against this global limit, creating strategic decisions about board development vs. individual room power.

- Default UC: 20 points
- Shared across all players
- Prevents runaway board states
- Forces strategic choices between quantity and quality

### Room Point Value (RPV) & Room Space (RS)
**Room Point Value** serves dual purposes:
1. **Consumes Universe Capacity**: Each room's RPV counts against the global UC limit
2. **Determines Room Space**: RPV directly sets the room's internal capacity (RS)

**Room Space** determines how many and what size objects can occupy a room:
- Objects have **Object Size (OS)** values
- Total OS of occupants cannot exceed room's RS
- Provides tactical depth for unit placement and room design

### Footprints
Rooms can block adjacent slots on their row to maintain clean tactical lanes:
- Prevents opponent expansion into key positions
- Creates strategic chokepoints
- Enables lane control tactics
- Adjacency blocking is configurable per room

### Spawn Distance & Movement
**Spawn Distance** mechanics eliminate traditional "summoning sickness":
- Actors enter play in their player's **home row**
- **Movement forward** (toward opponent) removes movement restrictions
- Encourages aggressive, forward-thinking play
- Creates natural flow toward confrontation

### Map-Based Gameplay
The game emphasizes **territorial control**:
- **Rooms have exits** defining movement paths
- **Room control** enables expansion into new board slots
- Victory through strategic positioning and expansion
- Board state directly impacts available actions

## Resolution System

### Opposed d20 Rolls
Combat and contests use **opposed d20 rolls** with three outcome bands:

#### Result Bands
- **FAIL**: Roll difference < 2 (no effect)
- **OK**: Roll difference ≥ 2 (basic effect)  
- **CRIT**: Roll difference ≥ 2 AND natural 20 (enhanced effect)

#### Modifiers
Resolution can be modified by:
- **Room rules**: Tag-based bonuses (e.g., Soldier +2 in Military rooms)
- **Card effects**: Direct roll modifiers (ROLL_MOD effects)
- **Archetype interactions**: Type advantages/disadvantages

#### Margin System
**Win by ≥2** to secure meaningful effects:
- **Margin**: Difference between winning and losing rolls
- **Effect Threshold**: Margin ≥2 required for actions like removing OS=1 enemy units
- Prevents lucky single-point victories from having major impact
- Requires decisive victories for significant effects

#### Manipulation Hooks *(Future Feature)*
Planned timing windows for:
- **Rerolls**: Second chances on important rolls
- **Modifiers**: +/- adjustments during resolution
- Creates interactive, reactive gameplay moments

## Archetypes

The game features three core **archetypes** that define strategic approaches:

### Soldier
- **Specialty**: Steady attack modifiers
- **Strength**: Reliable combat performance
- **Weakness**: Vulnerable to magical attacks
- **Playstyle**: Direct confrontation and board control

### Mage  
- **Specialty**: Bypasses armor and defenses
- **Strength**: Ignores traditional protection
- **Weakness**: Fragile to physical damage
- **Playstyle**: Glass cannon with powerful but risky tactics

### Rogue
- **Specialty**: Roll manipulation and mobility
- **Strength**: Controls probability and positioning
- **Weakness**: Moderate stats require tactical finesse
- **Playstyle**: Trickery, positioning, and calculated risks

### Archetype Interactions
- **Tags on cards**: Enable room-specific bonuses
- **Room keying**: Rooms can provide bonuses to specific archetypes
- **Strategic depth**: Mixing archetypes vs. specialization decisions

## Getting Started

### Prerequisites
- Go 1.21 or later

### Installation
```bash
# Clone the repository
git clone https://github.com/evevioletrose-hash/tcg-simulator.git
cd tcg-simulator

# Build the project
go build ./...

# Run the simulator
go run cmd/tcg-sim/main.go
```

### Basic Usage
The simulator currently provides a foundation with:
- 5x3 game board
- Two-player setup with home rows
- Universe Capacity of 20 points
- Core data structures for cards, rooms, and players

## Game Components

### Field Structure
- **Grid-based board**: Configurable dimensions (default 5x3)
- **Directional movement**: North, East, South, West exits
- **Room control**: Neutral (-1) or player-controlled (0/1)

### Card System
- **Object Size (OS)**: Determines space consumption
- **Multiple Archetypes**: Cards can have multiple type tags
- **Effect System**: Expandable framework for card abilities
- **Tag System**: Flexible categorization for interactions

### Player Management
- **Home Row**: Starting position for new units
- **Hand Management**: Card collection and play
- **Expandable**: Ready for deck, discard, and other zones

## Examples

### Basic Room Setup
```go
// Create a room with RPV=3 (consumes 3 UC, provides 3 RS)
room := &Room{
    RPV: 3,           // Consumes 3 Universe Capacity
    RS:  3,           // Provides 3 Room Space
    Footprints: []int{1, 2}, // Blocks columns 1 and 2 on this row
}

// Place units up to RS limit
soldier := &Card{OS: 2} // Consumes 2 Room Space
scout := &Card{OS: 1}   // Consumes 1 Room Space
// Total OS: 3, exactly fills RS: 3
```

### Dice Resolution Example
```go
// Combat between modified rolls
resultA, resultB := OpposedD20(+2, +1) // Soldier vs Mage modifiers

// Check for significant victory
if resultA.Margin >= 2 {
    // Soldier wins with margin ≥2, can remove OS=1 enemy
    if resultA.Band == "CRIT" {
        // Natural 20 with margin ≥2, enhanced effect
    }
}
```

### Universe Capacity Management
```go
game := NewGame() // UC: 20
// Room deployments must stay within UC limit:
// Room A: RPV 5 (15 UC remaining)
// Room B: RPV 8 (7 UC remaining)  
// Room C: RPV 7 (0 UC remaining)
// No more rooms can be placed until some are removed
```

---

**Status**: Core mechanics implemented, game loop and advanced features in development.

**Contributing**: This project implements a novel approach to TCG design focusing on spatial tactics and resource management. Contributions welcome!
