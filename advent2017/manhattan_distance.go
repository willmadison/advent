package advent2017

func ManhattanDistanceTo(cellNumber int) int {
	point := toCartesianCoordinate(cellNumber)
	return point.manhattanDistanceFromOrigin()
}

type Point struct {
	x, y int
}

func (p Point) manhattanDistanceFromOrigin() int {
	if p.x < 0 {
		p.x = -p.x
	}

	if p.y < 0 {
		p.y = -p.y
	}

	return p.x + p.y
}

func toCartesianCoordinate(cellNumber int) Point {
	var point Point

	l := 1
	var breadth int

	for l*l < cellNumber {
		l += 2
		breadth++
	}

	if breadth > 0 {
		point.x, point.y = breadth, -breadth
	}

	current := l * l

	var moves int

	for current != cellNumber && moves < l-1 {
		//Move left
		point.x--
		moves++
		current--
	}

	moves = 0
	for current != cellNumber && moves < l-1 {
		//Move up
		point.y++
		moves++
		current--
	}

	moves = 0
	for current != cellNumber && moves < l-1 {
		//Move right
		point.x++
		moves++
		current--
	}

	moves = 0
	for current != cellNumber && moves < l-1 {
		//Move down
		point.y--
		moves++
		current--
	}

	return point
}
