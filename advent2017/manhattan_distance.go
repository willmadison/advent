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
	for current != cellNumber && moves < l-2 {
		//Move down
		point.y--
		moves++
		current--
	}

	return point
}

var valuesByLocation map[Point]int

func FirstCellLargerThan(target int) int {
	valuesByLocation = map[Point]int{
		Point{0, 0}: 1,
		Point{1, 0}: 1,
	}

	currentLocation := Point{1, 0}
	value := valuesByLocation[currentLocation]

	l := 3
	for value <= target {
		moves := 0
		for value <= target && moves < l-2 {
			//Move up
			currentLocation = Point{currentLocation.x, currentLocation.y + 1}
			value = calculateValueAt(currentLocation)
			valuesByLocation[currentLocation] = value

			moves++
		}

		moves = 0
		for value <= target && moves < l-1 {
			///Move left
			currentLocation = Point{currentLocation.x - 1, currentLocation.y}
			value = calculateValueAt(currentLocation)
			valuesByLocation[currentLocation] = value

			moves++
		}

		moves = 0
		for value <= target && moves < l-1 {
			///Move down
			currentLocation = Point{currentLocation.x, currentLocation.y - 1}
			value = calculateValueAt(currentLocation)
			valuesByLocation[currentLocation] = value

			moves++
		}

		moves = 0
		for value <= target && moves < l {
			///Move right
			currentLocation = Point{currentLocation.x + 1, currentLocation.y}
			value = calculateValueAt(currentLocation)
			valuesByLocation[currentLocation] = value

			moves++
		}

		if value <= target {
			l += 2
		}
	}

	return value
}

func calculateValueAt(point Point) int {
	neighboringPoints := getNeighbors(point)

	var sum int

	for _, n := range neighboringPoints {
		sum += valuesByLocation[n]
	}

	return sum
}

func getNeighbors(point Point) []Point {
	neighbors := []Point{}

	neighbors = append(neighbors, Point{point.x + 1, point.y})     // Right neighbor
	neighbors = append(neighbors, Point{point.x + 1, point.y + 1}) // Upper right neighbor
	neighbors = append(neighbors, Point{point.x, point.y + 1})     // Upper neighbor
	neighbors = append(neighbors, Point{point.x - 1, point.y + 1}) // Upper left neighbor
	neighbors = append(neighbors, Point{point.x - 1, point.y})     // Left neighbor
	neighbors = append(neighbors, Point{point.x - 1, point.y - 1}) // Bottom left neighbor
	neighbors = append(neighbors, Point{point.x, point.y - 1})     // Bottom neighbor
	neighbors = append(neighbors, Point{point.x + 1, point.y - 1}) // Bottom right neighbor

	return neighbors
}
