package advent2020

import (
	"bufio"
	"io"

	"github.com/willmadison/advent/internal/location"
)

type CubeStatus int

const (
	_ CubeStatus = iota
	Active
	Inactive
)

type Cube struct {
	CellsByLocation map[location.Point]CubeCell
}

func (c *Cube) RunCycle() {
	activeCubes := map[location.Point]CubeCell{}
	inactiveCube := map[location.Point]CubeCell{}

	shouldBeActivated := []location.Point{}
	shouldBeDeactivated := []location.Point{}

	additions := map[location.Point]struct{}{}

	for location, cell := range c.CellsByLocation {
		if _, recentlyAdded := additions[location]; !recentlyAdded {
			neighbors := cell.NeighborLocations()

			for _, neighbor := range neighbors {
				if _, mapped := c.CellsByLocation[neighbor]; !mapped {
					additionalCell := CubeCell{neighbor, Inactive}
					c.CellsByLocation[neighbor] = additionalCell
					additions[neighbor] = struct{}{}
				}
			}
		}
	}

	for location, cell := range c.CellsByLocation {
		if cell.Status == Active {
			activeCubes[location] = cell
		} else {
			inactiveCube[location] = cell
		}
	}

	for _, cell := range c.CellsByLocation {
		neighbors := cell.NeighborLocations()

		var activeNeighbors int

		for _, neighbor := range neighbors {
			if _, isActive := activeCubes[neighbor]; isActive {
				activeNeighbors++
			}
		}

		if cell.Status == Active && activeNeighbors != 2 && activeNeighbors != 3 {
			shouldBeDeactivated = append(shouldBeDeactivated, cell.Coordinate)
		} else if cell.Status == Inactive && activeNeighbors == 3 {
			shouldBeActivated = append(shouldBeActivated, cell.Coordinate)
		}
	}

	for _, location := range shouldBeActivated {
		cell := c.CellsByLocation[location]
		cell.Status = Active
		c.CellsByLocation[location] = cell
	}

	for _, location := range shouldBeDeactivated {
		cell := c.CellsByLocation[location]
		cell.Status = Inactive
		c.CellsByLocation[location] = cell
	}
}

func (c Cube) CellCountByStatus(status CubeStatus) int {
	var count int

	for _, cell := range c.CellsByLocation {
		if cell.Status == status {
			count++
		}
	}

	return count
}

type CubeCell struct {
	Coordinate location.Point
	Status     CubeStatus
}

func (c CubeCell) NeighborLocations() []location.Point {
	var neighbors []location.Point

	for deltaX := -1; deltaX <= 1; deltaX++ {
		for deltaY := -1; deltaY <= 1; deltaY++ {
			for deltaZ := -1; deltaZ <= 1; deltaZ++ {
				for deltaW := -1; deltaW <= 1; deltaW++ {
					if deltaX == 0 && deltaY == 0 && deltaZ == 0 && deltaW == 0 {
						continue
					}
					neighbor := location.Point{c.Coordinate.X + deltaX, c.Coordinate.Y + deltaY, c.Coordinate.Z + deltaZ, c.Coordinate.W + deltaW}
					neighbors = append(neighbors, neighbor)
				}
			}
		}
	}

	return neighbors
}

func ParseConwayCube(r io.Reader) Cube {
	var cubeCells []CubeCell

	scanner := bufio.NewScanner(r)

	var y int
	for scanner.Scan() {
		cells := parseCubeRow(scanner.Text(), y)
		cubeCells = append(cubeCells, cells...)
		y++
	}

	cellsByLocation := map[location.Point]CubeCell{}

	for _, cell := range cubeCells {
		func(c CubeCell) {
			cellsByLocation[cell.Coordinate] = c
		}(cell)
	}

	return Cube{CellsByLocation: cellsByLocation}
}

func parseCubeRow(row string, y int) []CubeCell {
	var cells []CubeCell

	for x, rawStatus := range row {
		status := Inactive

		if rune(rawStatus) == '#' {
			status = Active
		}

		cells = append(cells, CubeCell{location.Point{X: x, Y: y, Z: 0, W: 0}, status})
	}

	return cells
}
