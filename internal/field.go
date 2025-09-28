package internal

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Field struct {
	Grid [][]*Room
	Rows, Cols int
}

func NewField(rows, cols int) *Field {
	grid := make([][]*Room, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]*Room, cols)
		for j := 0; j < cols; j++ {
			grid[i][j] = &Room{
				Row: i,
				Col: j,
			}
		}
	}
	return &Field{
		Grid:  grid,
		Rows:  rows,
		Cols:  cols,
	}
}

type Room struct {
	Row, Col   int
	RPV        int      // Room Point Value
	RS         int      // Room Space
	Footprints []int    // Blocked adjacent columns (for lane control)
	Exits      []Direction
	Occupants  []*Card
	ControlledBy int    // -1 = neutral, 0/1 = player index
}
