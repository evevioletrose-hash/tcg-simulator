package internal

type Player struct {
	Name    string
	HomeRow int
	Hand    []*Card
	// TODO: add deck, discard, etc.
}

func NewPlayer(name string, homeRow int) *Player {
	return &Player{
		Name:    name,
		HomeRow: homeRow,
		Hand:    []*Card{}, // Start with an empty hand
	}
}
