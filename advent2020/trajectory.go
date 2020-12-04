package advent2020

import (
	"bufio"
	"io"

	"github.com/willmadison/advent/internal/location"
)

type TrajectoryMap struct {
	Trees         map[location.Coordinate]struct{}
	Rows, Columns int
}

func NewTrajectoryMap(r io.Reader) TrajectoryMap {
	var treeCoordinates []location.Coordinate

	scanner := bufio.NewScanner(r)

	var row int
	var columns int

	for scanner.Scan() {
		trajectoryRow := scanner.Text()

		if columns == 0 {
			columns = len(trajectoryRow)
		}

		coordinates := findTrees(row, trajectoryRow)
		treeCoordinates = append(treeCoordinates, coordinates...)
		row++
	}

	trees := map[location.Coordinate]struct{}{}

	for _, c := range treeCoordinates {
		trees[c] = struct{}{}
	}

	return TrajectoryMap{Trees: trees, Rows: row, Columns: columns}
}

func findTrees(rowNumber int, trajectoryRow string) []location.Coordinate {
	var treeCoordinates []location.Coordinate

	tree := '#'

	for col, c := range trajectoryRow {
		if rune(c) == tree {
			coordinate := location.Coordinate{rowNumber, col}
			treeCoordinates = append(treeCoordinates, coordinate)
		}
	}

	return treeCoordinates
}

func CountEncounteredTrees(tm TrajectoryMap, slope location.Slope) int {
	current := location.Coordinate{0, 0}

	var treesEncountered int

	for current.Row < tm.Rows {
		current.Col += slope.Run
		current.Row += slope.Rise

		locale := location.Coordinate{current.Row, current.Col % tm.Columns}

		if _, present := tm.Trees[locale]; present {
			treesEncountered++
		}
	}

	return treesEncountered
}
