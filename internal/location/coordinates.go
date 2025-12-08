package location

import "slices"

type Coordinate struct {
	Row, Col int
}

func (c Coordinate) Neighbors() []Coordinate {
	return []Coordinate{
		{Row: c.Row - 1, Col: c.Col},
		{Row: c.Row - 1, Col: c.Col + 1},
		{Row: c.Row, Col: c.Col + 1},
		{Row: c.Row + 1, Col: c.Col + 1},
		{Row: c.Row + 1, Col: c.Col},
		{Row: c.Row + 1, Col: c.Col - 1},
		{Row: c.Row, Col: c.Col - 1},
		{Row: c.Row - 1, Col: c.Col - 1},
	}
}

func (c Coordinate) WithNextN(n int, direction CardinalDirection) []Coordinate {
	var coordinates []Coordinate

	var delta Coordinate

	switch direction {
	case North:
		delta = Coordinate{Row: -1, Col: 0}
	case Northeast:
		delta = Coordinate{Row: -1, Col: 1}
	case East:
		delta = Coordinate{Row: 0, Col: 1}
	case Southeast:
		delta = Coordinate{Row: 1, Col: 1}
	case South:
		delta = Coordinate{Row: 1, Col: 0}
	case Southwest:
		delta = Coordinate{Row: 1, Col: -1}
	case West:
		delta = Coordinate{Row: 0, Col: -1}
	case Northwest:
		delta = Coordinate{Row: -1, Col: -1}
	}

	coordinate := Coordinate{Row: c.Row, Col: c.Col}

	coordinates = append(coordinates, coordinate)

	for len(coordinates) < n+1 {
		coordinate = Coordinate{Row: coordinate.Row + delta.Row, Col: coordinate.Col + delta.Col}
		coordinates = append(coordinates, coordinate)
	}

	slices.SortFunc(coordinates, func(a, b Coordinate) int {
		if a.Row != b.Row {
			return a.Row - b.Row
		}

		return a.Col - b.Col
	})

	return coordinates
}

func (c Coordinate) InBounds(rows, cols int) bool {
	return c.Row >= 0 && c.Row < rows && c.Col >= 0 && c.Col < cols
}

func (c Coordinate) Delta(other Coordinate) Coordinate {
	return Coordinate{Row: c.Row - other.Row, Col: c.Col - other.Col}
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

func (p Point) EuclideanDistance(other Point) int {
	return (p.X-other.X)*(p.X-other.X) + (p.Y-other.Y)*(p.Y-other.Y) + (p.Z-other.Z)*(p.Z-other.Z)
}

type CardinalDirection uint

const (
	North CardinalDirection = iota
	Northeast
	East
	Southeast
	South
	Southwest
	West
	Northwest
)

func (d CardinalDirection) Reverse() CardinalDirection {
	return (d + 2) % 4
}

type Vector struct {
	Point
	Direction CardinalDirection
}

var Vectors = []Vector{
	{Point: Point{X: 0, Y: 1}, Direction: East},
	{Point: Point{X: 0, Y: -1}, Direction: West},
	{Point: Point{X: 1, Y: 0}, Direction: South},
	{Point: Point{X: -1, Y: 0}, Direction: North},
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

func (p Point) Add(other Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
