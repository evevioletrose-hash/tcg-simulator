package internal

type Archetype string

const (
	Soldier Archetype = "Soldier"
	Mage    Archetype = "Mage"
	Rogue   Archetype = "Rogue"
)

type Card struct {
	ID         string
	Name       string
	OS         int          // Object Size
	Archetypes []Archetype
	Tags       []string
	Effects    []Effect
	Owner      int
}

type Effect struct {
	Description string
	// TODO: Expand for hooks, modifiers, etc.
}
