# Examples and Use Cases

This document provides practical examples demonstrating the TCG Simulator's core mechanics through realistic gameplay scenarios.

## Table of Contents

- [Basic Gameplay Flow](#basic-gameplay-flow)
- [Resource Management Examples](#resource-management-examples)
- [Combat Scenarios](#combat-scenarios)
- [Spatial Tactics](#spatial-tactics)
- [Archetype Strategies](#archetype-strategies)
- [Advanced Combinations](#advanced-combinations)

## Basic Gameplay Flow

### Game Initialization

```go
package main

import (
    "fmt"
    "github.com/evevioletrose-hash/tcg-simulator/internal"
)

func main() {
    // Create a new game with default settings
    game := NewGame()
    
    fmt.Printf("Game created with %d UC available\n", game.UC)
    fmt.Printf("Board size: %dx%d\n", game.Field.Rows, game.Field.Cols)
    fmt.Printf("Player 1 home row: %d\n", game.Players[0].HomeRow)
    fmt.Printf("Player 2 home row: %d\n", game.Players[1].HomeRow)
}

// Output:
// Game created with 20 UC available
// Board size: 5x3
// Player 1 home row: 0
// Player 2 home row: 4
```

### First Turn Setup

```go
func firstTurnExample() {
    game := NewGame()
    player1 := game.Players[0]
    
    // Create a basic soldier card
    soldier := &Card{
        ID:         "soldier_001",
        Name:       "Infantry Unit",
        OS:         2,
        Archetypes: []Archetype{Soldier},
        Tags:       []string{"military", "basic"},
        Owner:      0,
    }
    
    // Add to player's hand
    player1.AddToHand(soldier)
    
    // Find spawn location
    spawnRoom := player1.CanSpawnInHomeRow(game.Field, soldier)
    if spawnRoom != nil {
        fmt.Printf("Can spawn soldier at (%d, %d)\n", spawnRoom.Row, spawnRoom.Col)
        
        // Place the soldier (in a real game, this would be a method call)
        spawnRoom.Occupants = append(spawnRoom.Occupants, soldier)
        fmt.Printf("Soldier deployed! Room space used: %d/%d\n", 
                   spawnRoom.GetTotalObjectSize(), spawnRoom.RS)
    }
}
```

## Resource Management Examples

### Universe Capacity Planning

```go
func universeCapacityExample() {
    game := NewGame()
    startingUC := game.UC // 20
    
    // Player 1's strategy: Few powerful rooms
    room1 := &Room{
        Row: 1, Col: 1,
        RPV: 8,  // High-value room
        RS:  8,  // Large capacity
        ControlledBy: 0,
    }
    
    room2 := &Room{
        Row: 2, Col: 0,
        RPV: 6,  // Medium-value room
        RS:  6,
        ControlledBy: 0,
    }
    
    usedUC := room1.RPV + room2.RPV // 14
    remainingUC := startingUC - usedUC // 6
    
    fmt.Printf("Player 1 strategy: %d UC used, %d remaining\n", usedUC, remainingUC)
    
    // Player 2's counter-strategy: Many small rooms
    room3 := &Room{Row: 3, Col: 0, RPV: 2, RS: 2, ControlledBy: 1}
    room4 := &Room{Row: 3, Col: 1, RPV: 2, RS: 2, ControlledBy: 1}
    room5 := &Room{Row: 3, Col: 2, RPV: 2, RS: 2, ControlledBy: 1}
    
    player2UC := room3.RPV + room4.RPV + room5.RPV // 6
    finalUC := remainingUC - player2UC // 0
    
    fmt.Printf("Player 2 counter: %d UC used, %d remaining\n", player2UC, finalUC)
    fmt.Println("Universe Capacity exhausted - no more room deployment possible!")
}
```

### Room Space Optimization

```go
func roomSpaceExample() {
    // Room with RS=5 - different deployment strategies
    room := &Room{
        RPV: 5,
        RS:  5,
        Occupants: []*Card{},
    }
    
    // Strategy 1: Swarm approach
    scouts := []*Card{
        {Name: "Scout A", OS: 1},
        {Name: "Scout B", OS: 1},
        {Name: "Scout C", OS: 1},
        {Name: "Scout D", OS: 1},
        {Name: "Scout E", OS: 1},
    }
    
    totalOS := 0
    for _, scout := range scouts {
        if room.CanAccommodate(scout) {
            room.Occupants = append(room.Occupants, scout)
            totalOS += scout.OS
        }
    }
    fmt.Printf("Swarm strategy: %d units, %d/%d space used\n", 
               len(room.Occupants), totalOS, room.RS)
    
    // Strategy 2: Elite approach
    room.Occupants = []*Card{} // Reset
    elite := &Card{Name: "Elite Champion", OS: 4}
    support := &Card{Name: "Support Unit", OS: 1}
    
    room.Occupants = append(room.Occupants, elite, support)
    fmt.Printf("Elite strategy: %d units, %d/%d space used\n",
               len(room.Occupants), elite.OS+support.OS, room.RS)
}
```

## Combat Scenarios

### Basic Combat Resolution

```go
func basicCombatExample() {
    // Create opposing forces
    attacker := &Card{
        Name:       "Veteran Soldier",
        Archetypes: []Archetype{Soldier},
        OS:         2,
    }
    
    defender := &Card{
        Name:       "Apprentice Mage", 
        Archetypes: []Archetype{Mage},
        OS:         1,
    }
    
    room := &Room{} // Neutral room
    
    // Resolve combat
    attackResult, defendResult, eliminated := ResolveCombat(attacker, defender, room)
    
    fmt.Printf("Combat Results:\n")
    fmt.Printf("Attacker rolled %d (%s), margin: %d\n", 
               attackResult.Roll, attackResult.Band, attackResult.Margin)
    fmt.Printf("Defender rolled %d (%s), margin: %d\n",
               defendResult.Roll, defendResult.Band, defendResult.Margin)
    
    if attackResult.Margin >= 2 {
        fmt.Println("Attacker wins decisively!")
        if eliminated {
            fmt.Println("Defender eliminated (OS=1 with decisive loss)")
        }
    } else if defendResult.Margin >= 2 {
        fmt.Println("Defender wins decisively!")
    } else {
        fmt.Println("Inconclusive result - no major effects")
    }
}
```

### Archetype Advantage Example

```go
func archetypeAdvantageExample() {
    // Soldier vs Mage scenario
    soldier := &Card{
        Name:       "Heavy Infantry",
        Archetypes: []Archetype{Soldier},
    }
    
    mage := &Card{
        Name:       "Battle Mage",
        Archetypes: []Archetype{Mage},
    }
    
    // In this example, we'll simulate the advantage system
    soldierMod := soldier.GetCombatModifier(mage, nil) // +1 for Soldier
    mageMod := mage.GetCombatModifier(soldier, nil)    // 0 base
    
    fmt.Printf("Soldier modifier: +%d (steady combat bonus)\n", soldierMod)
    fmt.Printf("Mage modifier: +%d (no bonus vs Soldier)\n", mageMod)
    
    // Simulate multiple combats to show statistical advantage
    soldierWins := 0
    mageWins := 0
    ties := 0
    
    for i := 0; i < 100; i++ {
        aResult, dResult := OpposedD20(soldierMod, mageMod)
        if aResult.Margin >= 2 {
            soldierWins++
        } else if dResult.Margin >= 2 {
            mageWins++
        } else {
            ties++
        }
    }
    
    fmt.Printf("Results over 100 combats:\n")
    fmt.Printf("Soldier decisive wins: %d\n", soldierWins)
    fmt.Printf("Mage decisive wins: %d\n", mageWins)
    fmt.Printf("Inconclusive: %d\n", ties)
}
```

## Spatial Tactics

### Footprint Control

```go
func footprintExample() {
    field := NewField(5, 3)
    
    // Player 1 places a room with strategic footprint
    stronghold := &Room{
        Row:          2,
        Col:          1,
        RPV:          6,
        RS:           6,
        Footprints:   []int{0, 2}, // Blocks columns 0 and 2 on row 2
        ControlledBy: 0,
    }
    field.Grid[2][1] = stronghold
    
    fmt.Println("Player 1 places stronghold at (2,1)")
    fmt.Println("Footprint blocks columns 0 and 2 on row 2")
    
    // Player 2 tries to expand
    for col := 0; col < field.Cols; col++ {
        targetRoom := field.Grid[2][col]
        if stronghold.BlocksColumn(col) {
            fmt.Printf("Column %d: BLOCKED by stronghold footprint\n", col)
        } else if targetRoom.ControlledBy == 0 {
            fmt.Printf("Column %d: Occupied by Player 1\n", col)
        } else {
            fmt.Printf("Column %d: Available for Player 2\n", col)
        }
    }
    
    // Output:
    // Column 0: BLOCKED by stronghold footprint
    // Column 1: Occupied by Player 1  
    // Column 2: BLOCKED by stronghold footprint
}
```

### Movement and Exits

```go
func movementExample() {
    field := NewField(5, 3)
    
    // Create a movement path
    roomA := &Room{
        Row:   1,
        Col:   1, 
        Exits: []Direction{South, East},
    }
    
    roomB := &Room{
        Row:   2,
        Col:   1,
        Exits: []Direction{North, South, West},
    }
    
    roomC := &Room{
        Row:   1,
        Col:   2,
        Exits: []Direction{West, South},
    }
    
    field.Grid[1][1] = roomA
    field.Grid[2][1] = roomB  
    field.Grid[1][2] = roomC
    
    // Check valid movement from roomA
    fmt.Println("Movement options from Room A (1,1):")
    
    if roomA.HasExit(South) {
        targetRoom := field.Grid[2][1] // roomB
        if targetRoom.HasExit(North) {
            fmt.Println("- Can move South to Room B (2,1)")
        }
    }
    
    if roomA.HasExit(East) {
        targetRoom := field.Grid[1][2] // roomC
        if targetRoom.HasExit(West) {
            fmt.Println("- Can move East to Room C (1,2)")
        }
    }
    
    if roomA.HasExit(North) {
        fmt.Println("- Cannot move North (no exit)")
    } else {
        fmt.Println("- North movement blocked (no exit)")
    }
}
```

## Archetype Strategies

### Soldier Tactics

```go
func soldierStrategy() {
    fmt.Println("=== SOLDIER STRATEGY EXAMPLE ===")
    
    // Soldier deck focused on consistent pressure
    soldiers := []*Card{
        {
            Name:       "Frontline Infantry",
            OS:         2,
            Archetypes: []Archetype{Soldier},
            Tags:       []string{"military", "infantry"},
        },
        {
            Name:       "Elite Guard",
            OS:         3,
            Archetypes: []Archetype{Soldier},
            Tags:       []string{"military", "elite"},
        },
        {
            Name:       "Field Commander",
            OS:         2,
            Archetypes: []Archetype{Soldier},
            Tags:       []string{"military", "leader"},
        },
    }
    
    // Optimal room for soldiers: Military-tagged rooms
    militaryBase := &Room{
        RPV:  5,
        RS:   5,
        Tags: []string{"military"}, // Would provide Soldier bonuses
    }
    
    fmt.Println("Soldier Strategy:")
    fmt.Println("- Deploy multiple reliable units")
    fmt.Println("- Control military-tagged rooms for bonuses")
    fmt.Println("- Maintain steady pressure on multiple fronts")
    
    totalOS := 0
    for _, soldier := range soldiers {
        totalOS += soldier.OS
    }
    fmt.Printf("Total force OS: %d (fits in RS=%d room)\n", totalOS, militaryBase.RS)
}
```

### Mage Glass Cannon

```go
func mageStrategy() {
    fmt.Println("=== MAGE STRATEGY EXAMPLE ===")
    
    // Mage approach: High impact, fragile units
    mageForce := []*Card{
        {
            Name:       "Arcane Destroyer",
            OS:         1, // Fragile but powerful
            Archetypes: []Archetype{Mage},
            Tags:       []string{"arcane", "destruction"},
        },
        {
            Name:       "Support Wizard", 
            OS:         1,
            Archetypes: []Archetype{Mage},
            Tags:       []string{"arcane", "support"},
        },
    }
    
    // Mages prefer arcane-tagged rooms
    arcaneLibrary := &Room{
        RPV:  4,
        RS:   4,
        Tags: []string{"arcane"}, // Would provide Mage bonuses
    }
    
    fmt.Println("Mage Strategy:")
    fmt.Println("- Deploy glass cannon units (low OS, high impact)")
    fmt.Println("- Target enemy armor and defenses")
    fmt.Println("- Avoid direct confrontation with Soldiers")
    fmt.Println("- Control arcane rooms for magical bonuses")
    
    // Can fit 4 OS=1 mages in RS=4 room - maximum flexibility
    fmt.Printf("Can deploy %d mages in RS=%d arcane room\n", 
               arcaneLibrary.RS, arcaneLibrary.RS)
}
```

### Rogue Manipulation

```go
func rogueStrategy() {
    fmt.Println("=== ROGUE STRATEGY EXAMPLE ===")
    
    rogueTeam := []*Card{
        {
            Name:       "Master Thief",
            OS:         1,
            Archetypes: []Archetype{Rogue},
            Tags:       []string{"stealth", "mobility"},
        },
        {
            Name:       "Shadow Assassin",
            OS:         2,
            Archetypes: []Archetype{Rogue},
            Tags:       []string{"stealth", "lethal"},
        },
    }
    
    // Rogues excel with multiple small rooms for mobility
    hidingSpots := []*Room{
        {RPV: 2, RS: 2, Tags: []string{"stealth"}},
        {RPV: 2, RS: 2, Tags: []string{"stealth"}},
        {RPV: 2, RS: 2, Tags: []string{"stealth"}},
    }
    
    fmt.Println("Rogue Strategy:")
    fmt.Println("- Deploy mobile, low-OS units")
    fmt.Println("- Control multiple small rooms for positioning")
    fmt.Println("- Use roll manipulation to win key combats")
    fmt.Println("- Avoid prolonged direct confrontation")
    
    totalRPV := 0
    for _, room := range hidingSpots {
        totalRPV += room.RPV
    }
    fmt.Printf("Total RPV for mobility network: %d\n", totalRPV)
}
```

## Advanced Combinations

### Hybrid Archetype Card

```go
func hybridExample() {
    // Battle Mage: Soldier/Mage hybrid
    battleMage := &Card{
        Name:       "Spellsword Champion",
        OS:         3, // Higher cost for versatility
        Archetypes: []Archetype{Soldier, Mage},
        Tags:       []string{"military", "arcane", "elite"},
    }
    
    fmt.Println("Hybrid Unit Example:")
    fmt.Printf("Unit: %s\n", battleMage.Name)
    fmt.Printf("Archetypes: %v\n", battleMage.Archetypes)
    fmt.Printf("Object Size: %d (premium cost for versatility)\n", battleMage.OS)
    
    // Check archetype bonuses
    if battleMage.HasArchetype(Soldier) {
        fmt.Println("- Gets Soldier combat bonuses")
    }
    if battleMage.HasArchetype(Mage) {
        fmt.Println("- Can bypass armor like a Mage")
    }
    
    // Can benefit from both military AND arcane rooms
    if battleMage.HasTag("military") && battleMage.HasTag("arcane") {
        fmt.Println("- Benefits from both Military and Arcane room bonuses")
    }
}
```

### Complex Board State

```go
func complexScenario() {
    fmt.Println("=== COMPLEX BOARD SCENARIO ===")
    
    game := NewGame()
    
    // Set up a mid-game board state
    // Player 1 (Soldier strategy): Control center with footprints
    centerFortress := &Room{
        Row:          2,
        Col:          1,
        RPV:          8,
        RS:           8,
        Footprints:   []int{0, 2},
        ControlledBy: 0,
        Occupants: []*Card{
            {Name: "Elite Commander", OS: 4, Archetypes: []Archetype{Soldier}},
            {Name: "Guard Squad", OS: 4, Archetypes: []Archetype{Soldier}},
        },
    }
    
    // Player 2 (Rogue strategy): Multiple small outposts
    outpost1 := &Room{
        Row: 1, Col: 0, RPV: 3, RS: 3, ControlledBy: 1,
        Occupants: []*Card{
            {Name: "Infiltrator", OS: 1, Archetypes: []Archetype{Rogue}},
            {Name: "Scout", OS: 2, Archetypes: []Archetype{Rogue}},
        },
    }
    
    outpost2 := &Room{
        Row: 3, Col: 2, RPV: 3, RS: 3, ControlledBy: 1,
        Occupants: []*Card{
            {Name: "Saboteur", OS: 1, Archetypes: []Archetype{Rogue}},
            {Name: "Thief", OS: 1, Archetypes: []Archetype{Rogue}},
        },
    }
    
    // Calculate resource usage
    usedUC := centerFortress.RPV + outpost1.RPV + outpost2.RPV
    remainingUC := game.UC - usedUC
    
    fmt.Printf("Board State:\n")
    fmt.Printf("Player 1 Fortress: RPV %d, %d units (OS %d total)\n",
               centerFortress.RPV, len(centerFortress.Occupants), 8)
    fmt.Printf("Player 2 Outpost 1: RPV %d, %d units (OS %d total)\n",
               outpost1.RPV, len(outpost1.Occupants), 3)
    fmt.Printf("Player 2 Outpost 2: RPV %d, %d units (OS %d total)\n",
               outpost2.RPV, len(outpost2.Occupants), 2)
    fmt.Printf("Universe Capacity: %d used, %d remaining\n", usedUC, remainingUC)
    
    fmt.Println("\nStrategic Analysis:")
    fmt.Println("- Player 1: Centralized power, lane control via footprints")
    fmt.Println("- Player 2: Distributed positioning, mobility advantage")
    fmt.Println("- Remaining UC allows for 1-2 small rooms or tactical adjustments")
}
```

These examples demonstrate the rich tactical depth available in the TCG Simulator while showing how the core mechanics create meaningful strategic choices. Each system reinforces the others to create emergent gameplay complexity from simple, well-defined rules.