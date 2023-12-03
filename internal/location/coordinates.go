package location

type Coordinate struct {
	Row, Col int
}

func (c Coordinate) Neighbors() []Coordinate {
	return []Coordinate{
		{Row: c.Row - 1, Col: c.Col - 1},
		{Row: c.Row - 1, Col: c.Col},
		{Row: c.Row - 1, Col: c.Col + 1},
		{Row: c.Row, Col: c.Col + 1},
		{Row: c.Row + 1, Col: c.Col + 1},
		{Row: c.Row + 1, Col: c.Col},
		{Row: c.Row + 1, Col: c.Col - 1},
		{Row: c.Row, Col: c.Col - 1},
	}
}

type Slope struct {
	Rise, Run int
}

type Point struct {
	X, Y, Z, W int
}

func (p Point) Neighbors() []Point {
	return []Point{
		{X: p.X + 1, Y: p.Y},
		{X: p.X - 1, Y: p.Y},
		{X: p.X, Y: p.Y + 1},
		{X: p.X, Y: p.Y - 1},
		{X: p.X + 1, Y: p.Y - 1},
		{X: p.X - 1, Y: p.Y + 1},
	}
}

func (p Point) ManhattanDistance(other Point) int {
	return abs(p.X-other.X) + abs(p.Y-other.Y)
}

type Direction int

const (
	_ Direction = iota
	Clockwise
	Counterclockwise
)

func (p *Point) Rotate90(direction Direction) {
	switch direction {
	case Clockwise:
		p.X, p.Y = p.Y, -p.X
	case Counterclockwise:
		p.X, p.Y = -p.Y, p.X
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
