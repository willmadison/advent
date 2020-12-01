package advent

type Point struct {
	X, Y int
}

func (p Point) ManhattanDistance(other Point) int {
	return abs(p.X-other.X) + abs(p.Y-other.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
