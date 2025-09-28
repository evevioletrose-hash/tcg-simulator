package internal

// Archetype defines the strategic role and combat specialization of cards.
// Archetypes determine combat modifiers, vulnerabilities, and special abilities.
// Cards can have multiple archetypes for hybrid strategies.
type Archetype string

const (
	// Soldier archetype: Reliable combat specialists
	// - Strength: Steady attack modifiers and defensive positioning
	// - Weakness: Vulnerable to magical attacks that bypass armor
	// - Playstyle: Direct confrontation and board control
	Soldier Archetype = "Soldier"

	// Mage archetype: Magical damage dealers
	// - Strength: Bypasses armor and conventional defenses
	// - Weakness: Fragile to physical damage and direct assault
	// - Playstyle: Glass cannon tactics with powerful but risky abilities
	Mage Archetype = "Mage"

	// Rogue archetype: Tactical manipulators
	// - Strength: Roll manipulation and superior mobility
	// - Weakness: Moderate base stats require tactical finesse
	// - Playstyle: Trickery, positioning, and calculated risks
	Rogue Archetype = "Rogue"
)

// Card represents a single game piece with stats, abilities, and tactical properties.
// Cards are the primary units that occupy rooms and engage in combat.
//
// Key systems:
//   - Object Size (OS) determines space consumption in rooms
//   - Archetypes provide combat bonuses and strategic identity
//   - Tags enable flexible interactions with room rules and effects
//   - Effects define special abilities and triggered actions
type Card struct {
	ID   string // Unique identifier for this card instance
	Name string // Display name for the card

	// Space System
	OS int // Object Size: how much Room Space this card consumes

	// Strategic Identity
	Archetypes []Archetype // Combat specializations (can have multiple)
	Tags       []string    // Flexible categorization for room interactions

	// Abilities
	Effects []Effect // Special abilities and triggered effects

	// Ownership
	Owner int // Player index (0 or 1) who owns this card
}

// Effect represents a special ability or triggered action on a card.
// Effects define how cards interact with the game state beyond basic combat.
//
// Future expansions will include:
//   - Timing hooks (when the effect triggers)
//   - Targeting systems (what the effect affects)
//   - Duration and persistence mechanics
//   - Cost and resource requirements
type Effect struct {
	Description string // Human-readable description of the effect
	// TODO: Expand for hooks, modifiers, etc.
	// Type        EffectType // Category of effect (combat, movement, etc.)
	// Timing      []Hook     // When this effect can trigger
	// Target      TargetType // What this effect can target
	// Modifiers   []Modifier // Stat adjustments or special rules
}

// HasArchetype checks if this card has the specified archetype.
// Used for determining combat bonuses and room-based advantages.
//
// Parameters:
//   - archetype: The archetype to check for
//
// Returns true if the card has this archetype.
func (c *Card) HasArchetype(archetype Archetype) bool {
	for _, a := range c.Archetypes {
		if a == archetype {
			return true
		}
	}
	return false
}

// HasTag checks if this card has the specified tag.
// Tags enable flexible interactions with room rules and other game effects.
//
// Parameters:
//   - tag: The tag to check for
//
// Returns true if the card has this tag.
func (c *Card) HasTag(tag string) bool {
	for _, t := range c.Tags {
		if t == tag {
			return true
		}
	}
	return false
}

// GetCombatModifier calculates combat bonuses based on archetypes and context.
// This method will be expanded to include:
//   - Room-based bonuses (e.g., Soldiers in Military rooms)
//   - Archetype advantages/disadvantages
//   - Equipment and enhancement effects
//
// Parameters:
//   - opponent: The opposing card (for archetype interactions)
//   - room: The room where combat occurs (for environmental bonuses)
//
// Returns the total combat modifier for this card.
func (c *Card) GetCombatModifier(opponent *Card, room *Room) int {
	modifier := 0
	
	// Base archetype modifiers (placeholder - will be expanded)
	if c.HasArchetype(Soldier) {
		modifier += 1 // Soldiers get steady combat bonus
	}
	
	// Future: Add room-based bonuses, archetype interactions, etc.
	
	return modifier
}

// IsOwnedBy checks if this card is owned by the specified player.
//
// Parameters:
//   - playerIndex: Player index to check (0 or 1)
//
// Returns true if the card is owned by the player.
func (c *Card) IsOwnedBy(playerIndex int) bool {
	return c.Owner == playerIndex
}
