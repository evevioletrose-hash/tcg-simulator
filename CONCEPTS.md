# Core Game Concepts

This document provides detailed explanations of the fundamental mechanics that drive the TCG Simulator.

## Resource Management

### Universe Capacity (UC)
Universe Capacity is the foundational resource constraint that shapes all strategic decisions. Unlike traditional TCGs where each player has individual resources, UC is a **shared global pool** that both players draw from.

#### Key Properties:
- **Shared Resource**: Both players compete for the same UC pool
- **Fixed Pool**: Default 20 points total across the entire game
- **Strategic Tension**: Every room one player builds reduces opportunities for the opponent
- **Escalation Control**: Prevents runaway board states where one player dominates

#### Strategic Implications:
```
Turn 1: Player A builds Room (RPV 3) → 17 UC remaining
Turn 2: Player B builds Room (RPV 5) → 12 UC remaining  
Turn 3: Player A builds Room (RPV 8) → 4 UC remaining
Turn 4: Player B can only build small rooms (RPV ≤ 4)
```

### Room Point Value (RPV) & Room Space (RS)
RPV serves a **dual purpose** in the resource economy:

1. **UC Consumption**: Each room's RPV counts against the global UC limit
2. **Internal Capacity**: RPV directly determines the room's RS (Room Space)

This creates a natural balance where powerful rooms (high RPV) both:
- Consume more of the shared resource pool
- Provide more internal capacity for units

#### Design Philosophy:
The RPV=RS equation ensures that room investment scales predictably. A player can't create a cheap room with massive internal capacity, nor a expensive room with tiny capacity.

### Object Size (OS) System
Object Size creates **tactical depth within rooms** by limiting what can coexist in the same space.

#### Placement Rules:
- Total OS of occupants ≤ Room Space (RS)
- Players must choose between quantity vs. quality
- Larger units consume more space but typically have more power

#### Example Scenarios:
```
Room with RS=4 can accommodate:
- 4 units of OS=1 each (swarm strategy)
- 2 units of OS=2 each (balanced approach)  
- 1 unit of OS=4 (single powerful unit)
- 1 unit of OS=3 + 1 unit of OS=1 (mixed strategy)
```

## Spatial Mechanics  

### Footprints & Lane Control
Footprints extend a room's influence beyond its grid position, enabling **strategic lane denial**.

#### Mechanics:
- Rooms can specify which adjacent columns they "block"
- Blocked columns cannot be developed by opponents on that row
- Creates strategic chokepoints and territorial control
- Enables defensive positioning and offensive pressure

#### Tactical Applications:
```
Row 2: [Room A] [  ?  ] [Room B]
       Footprint: blocks column 1

Player cannot place room at (2,1) due to Room A's footprint
Forces expansion around the blocked lane
Creates predictable movement channels
```

### Spawn Distance & Movement
The spawn distance system eliminates traditional "summoning sickness" through **forward momentum mechanics**.

#### Core Rules:
- All units spawn in their player's **home row**
- Movement **toward the opponent** removes movement restrictions
- Encourages aggressive, forward-thinking gameplay
- Creates natural battlefield progression

#### Strategic Flow:
1. **Spawn Phase**: Unit enters in home row (limited actions)
2. **Advance Phase**: Moving toward opponent activates full capabilities  
3. **Engagement Phase**: Units meet in contested middle ground
4. **Control Phase**: Victor claims territory and expands options

### Exit System & Map Control
Rooms define movement paths through their **exit configuration**, creating a tactical movement puzzle.

#### Exit Mechanics:
- Rooms specify which directions have valid exits
- Movement requires matching exit/entrance pairs
- Room control opens or closes movement paths
- Strategic positioning can cut off opponent expansion

#### Territory Control:
Controlling rooms provides:
- **Expansion Opportunities**: Access to adjacent uncontrolled spaces
- **Movement Paths**: Ability to traverse the battlefield
- **Defensive Positions**: Control over key strategic locations
- **Resource Denial**: Preventing opponent access to valuable positions

## Resolution System

### Opposed d20 Mechanics
The resolution system creates **meaningful differentiation** between outcome levels while maintaining tactical uncertainty.

#### Three-Band System:
- **FAIL** (Margin < 2): No significant effect, resources wasted
- **OK** (Margin ≥ 2): Standard success, basic effects trigger  
- **CRIT** (Margin ≥ 2 + Natural 20): Enhanced success, powerful effects

#### Design Philosophy:
The margin requirement (≥2) prevents **lucky single-point victories** from having major impact. Players must achieve **decisive wins** to secure meaningful advantages.

### Modifier Sources
Combat effectiveness comes from multiple stacking sources:

#### Archetype Bonuses:
- **Soldier**: Consistent combat bonuses (+1-2 typically)
- **Mage**: Armor penetration and magical effects
- **Rogue**: Situational bonuses and roll manipulation

#### Environmental Factors:
- **Room Tags**: Military rooms boost Soldiers, Arcane rooms boost Mages
- **Positional Advantage**: High ground, flanking, etc.
- **Equipment Effects**: Weapons, armor, magical items

#### Dynamic Modifiers:
- **Card Effects**: Temporary bonuses from played cards
- **Status Effects**: Ongoing conditions affecting performance
- **Timing Windows**: Reactive bonuses and penalties

### Margin & Effect Scaling
The margin system creates **graduated success levels**:

#### Standard Effects (Margin ≥ 2):
- Remove OS=1 enemy units
- Force repositioning
- Gain temporary advantages

#### Enhanced Effects (Margin ≥ 4):
- Remove larger enemy units  
- Claim room control
- Trigger powerful card abilities

#### Critical Effects (Natural 20 + Margin ≥ 2):
- Maximum damage potential
- Bypass normal restrictions
- Chain reaction opportunities

## Archetype System

### Combat Triangle
The three archetypes form a **strategic ecosystem** with distinct roles:

#### Soldier: The Reliable Foundation
- **Philosophy**: Steady performance over flashy plays
- **Mechanics**: Consistent combat bonuses, defensive positioning
- **Counters**: Magical attacks that bypass physical defenses
- **Excels Against**: Direct physical confrontations

#### Mage: The Glass Cannon
- **Philosophy**: High risk, high reward magical assault
- **Mechanics**: Armor penetration, powerful spell effects
- **Counters**: Fast physical attacks before spells activate
- **Excels Against**: Heavily armored opponents

#### Rogue: The Tactical Manipulator  
- **Philosophy**: Control probability and positioning
- **Mechanics**: Roll manipulation, superior mobility
- **Counters**: Overwhelming force that can't be outmaneuvered
- **Excels Against**: Predictable strategies and static defenses

### Hybrid Strategies
Cards can have **multiple archetypes**, enabling complex strategic combinations:

#### Common Hybrids:
- **Soldier/Mage**: Battle mages with magical armor and weapons
- **Rogue/Soldier**: Elite scouts with combat training  
- **Mage/Rogue**: Trickster mages with illusion and mobility

#### Strategic Trade-offs:
- Hybrid cards sacrifice specialization for versatility
- More expensive to field (higher OS typically)
- Require more complex positioning and timing
- Reward skilled tactical play

## Future Expansion Systems

### Manipulation Hooks
Planned **timing windows** for interactive gameplay:

#### Reroll System:
- Limited rerolls per turn/game
- Resource cost for second chances
- Counter-reroll mechanics for opponents

#### Modifier Windows:
- React to opponent's rolls with +/- adjustments
- Stack timing for maximum impact
- Resource management for modifier usage

### Advanced Effects
The Effect system will expand to include:

#### Triggered Abilities:
- **On Enter**: Effects when cards enter play
- **On Combat**: Effects during resolution
- **On Control**: Effects when claiming rooms

#### Persistent Effects:
- **Auras**: Ongoing bonuses to nearby units
- **Terrain**: Permanent battlefield modifications
- **Conditions**: Status effects lasting multiple turns

This foundation provides rich tactical depth while maintaining accessibility for new players.