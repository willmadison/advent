package advent2024

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/willmadison/advent/internal/location"
)

func CountXmases(r io.Reader) (int, error) {
	puzzle, err := parsePuzzle(r)

	if err != nil {
		return 0, err
	}

	count := doCountXmases(puzzle)

	return count, nil
}

func CountXs(r io.Reader) (int, error) {
	puzzle, err := parsePuzzle(r)

	if err != nil {
		return 0, err
	}

	count := doCountXs(puzzle)

	return count, nil
}

func parsePuzzle(r io.Reader) ([]string, error) {
	var puzzle []string

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		puzzle = append(puzzle, line)
	}

	return puzzle, nil
}

func doCountXmases(puzzle []string) int {
	var totalCount int

	found := map[string]struct{}{}

	for row, values := range puzzle {
		for col := range values {
			totalCount += countXmases(location.Coordinate{Row: row, Col: col}, found, puzzle)
		}
	}

	return totalCount
}

func doCountXs(puzzle []string) int {
	var totalCount int

	rows := len(puzzle)
	cols := len(puzzle[0])

	for row, values := range puzzle {
		for col, value := range values {
			coordinate := location.Coordinate{Row: row, Col: col}

			if value == 'A' {
				neighbors := coordinate.Neighbors()

				northeastNeighbor := neighbors[location.Northeast]
				southwestNeighbor := neighbors[location.Southwest]
				northwestNeighbor := neighbors[location.Northwest]
				southeastNeighbor := neighbors[location.Southeast]

				if northeastNeighbor.InBounds(rows, cols) && southwestNeighbor.InBounds(rows, cols) &&
					northwestNeighbor.InBounds(rows, cols) && southeastNeighbor.InBounds(rows, cols) {

					if ((puzzle[northeastNeighbor.Row][northeastNeighbor.Col] == 'M' &&
						puzzle[southwestNeighbor.Row][southwestNeighbor.Col] == 'S') ||
						(puzzle[northeastNeighbor.Row][northeastNeighbor.Col] == 'S' &&
							puzzle[southwestNeighbor.Row][southwestNeighbor.Col] == 'M')) &&
						((puzzle[northwestNeighbor.Row][northwestNeighbor.Col] == 'M' &&
							puzzle[southeastNeighbor.Row][southeastNeighbor.Col] == 'S') ||
							(puzzle[northwestNeighbor.Row][northwestNeighbor.Col] == 'S' &&
								puzzle[southeastNeighbor.Row][southeastNeighbor.Col] == 'M')) {
						totalCount++
					}
				}

			}
		}
	}

	return totalCount
}

func countXmases(startingPoint location.Coordinate, found map[string]struct{}, puzzle []string) int {
	var totalCount int

	var cardinalities [][]location.Coordinate

	cardinalities = append(cardinalities, startingPoint.WithNextN(3, location.North))
	cardinalities = append(cardinalities, startingPoint.WithNextN(3, location.Northeast))
	cardinalities = append(cardinalities, startingPoint.WithNextN(3, location.East))
	cardinalities = append(cardinalities, startingPoint.WithNextN(3, location.Southeast))
	cardinalities = append(cardinalities, startingPoint.WithNextN(3, location.South))
	cardinalities = append(cardinalities, startingPoint.WithNextN(3, location.Southwest))
	cardinalities = append(cardinalities, startingPoint.WithNextN(3, location.West))
	cardinalities = append(cardinalities, startingPoint.WithNextN(3, location.Northwest))

	for _, coordinates := range cardinalities {
		key := fmt.Sprintf("%+v", coordinates)

		if _, seen := found[key]; !seen {
			if spellsXmas(coordinates, puzzle) {
				totalCount++
				found[key] = struct{}{}
			}
		}

	}

	return totalCount
}

func spellsXmas(coordinates []location.Coordinate, puzzle []string) bool {
	for _, coordinate := range coordinates {
		if coordinate.Row < 0 || coordinate.Col < 0 ||
			coordinate.Row >= len(puzzle) || coordinate.Col >= len(puzzle[0]) {
			return false
		}
	}

	var buf strings.Builder

	for _, coordinate := range coordinates {
		buf.WriteByte(puzzle[coordinate.Row][coordinate.Col])
	}

	if buf.String() == "XMAS" {
		return true
	}

	buf.Reset()

	for i := len(coordinates) - 1; i >= 0; i-- {
		coordinate := coordinates[i]
		buf.WriteByte(puzzle[coordinate.Row][coordinate.Col])
	}

	return buf.String() == "XMAS"
}
