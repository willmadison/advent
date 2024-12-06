package advent2024

import (
	"bufio"
	"errors"
	"io"

	"github.com/willmadison/advent/internal/location"
)

type Guard struct {
	Location    location.Coordinate
	Orientation location.CardinalDirection
}

func (g Guard) Inbounds(dimensions location.Coordinate) bool {
	return g.Location.Row >= 0 && g.Location.Row < dimensions.Row &&
		g.Location.Col >= 0 && g.Location.Col < dimensions.Col
}

func (g Guard) CanAdvance(obstacles map[location.Coordinate]struct{}) bool {
	nextLocation := g.previewAdvancement()

	_, obstaclePresent := obstacles[nextLocation]

	return !obstaclePresent
}

func (g Guard) previewAdvancement() location.Coordinate {
	var nextLocation location.Coordinate

	switch g.Orientation {
	case location.North:
		nextLocation = location.Coordinate{Row: g.Location.Row - 1, Col: g.Location.Col}
	case location.East:
		nextLocation = location.Coordinate{Row: g.Location.Row, Col: g.Location.Col + 1}
	case location.South:
		nextLocation = location.Coordinate{Row: g.Location.Row + 1, Col: g.Location.Col}
	case location.West:
		nextLocation = location.Coordinate{Row: g.Location.Row, Col: g.Location.Col - 1}
	}

	return nextLocation
}

func (g *Guard) Advance() {
	g.Location = g.previewAdvancement()
}

func (g *Guard) TurnRight() {
	switch g.Orientation {
	case location.North:
		g.Orientation = location.East
	case location.East:
		g.Orientation = location.South
	case location.South:
		g.Orientation = location.West
	case location.West:
		g.Orientation = location.North
	}
}

func ParseMap(r io.Reader) (map[location.Coordinate]struct{}, Guard, location.Coordinate, error) {
	obstacles := map[location.Coordinate]struct{}{}
	guard := Guard{Orientation: location.North}

	var row, columns int

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		if columns == 0 {
			columns = len(line)
		}

		for col, char := range line {
			switch char {
			case '^':
				guard.Location = location.Coordinate{Row: row, Col: col}
			case '#':
				obstacle := location.Coordinate{Row: row, Col: col}
				obstacles[obstacle] = struct{}{}
			}
		}

		row++
	}

	return obstacles, guard, location.Coordinate{Row: row, Col: columns}, nil
}

func SimulatePatrol(guard Guard, obstacles map[location.Coordinate]struct{}, dimensions location.Coordinate) (map[location.Coordinate]struct{}, error) {
	patrolledLocations := map[location.Coordinate]struct{}{}
	patrolledLocations[guard.Location] = struct{}{}

	patrolledLocationOrientations := map[location.Coordinate]map[location.CardinalDirection]struct{}{}

	for guard.Inbounds(dimensions) {
		if guard.CanAdvance(obstacles) {
			guard.Advance()

			if guard.Inbounds(dimensions) {
				patrolledLocations[guard.Location] = struct{}{}

				if _, present := patrolledLocationOrientations[guard.Location]; !present {
					patrolledLocationOrientations[guard.Location] = map[location.CardinalDirection]struct{}{}
				}

				if _, present := patrolledLocationOrientations[guard.Location][guard.Orientation]; !present {
					patrolledLocationOrientations[guard.Location][guard.Orientation] = struct{}{}
				} else {
					return nil, errors.New("cycle detected")
				}
			}
		} else {
			guard.TurnRight()
		}
	}

	return patrolledLocations, nil
}

func IntroduceObstacles(guard Guard, obstacles map[location.Coordinate]struct{}, dimensions location.Coordinate) int {
	var loops int

	initialPosition := guard.Location
	initialOrientation := guard.Orientation

	for row := 0; row < dimensions.Row; row++ {
		for col := 0; col < dimensions.Col; col++ {
			potentialObstacle := location.Coordinate{Row: row, Col: col}

			if _, obstaclePresent := obstacles[potentialObstacle]; obstaclePresent || guard.Location == potentialObstacle {
				continue
			}

			obstacles[potentialObstacle] = struct{}{}

			_, err := SimulatePatrol(guard, obstacles, dimensions)

			if err != nil {
				loops++
			}

			guard.Location = initialPosition
			guard.Orientation = initialOrientation

			delete(obstacles, potentialObstacle)
		}
	}

	return loops
}
