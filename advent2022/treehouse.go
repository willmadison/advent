package advent2022

import (
	"bufio"
	"io"

	"github.com/willmadison/advent/internal/location"
)

type Tree int

func parseTreeGrid(r io.Reader) (map[location.Coordinate]Tree, int, int) {
	treesByLocation := map[location.Coordinate]Tree{}

	scanner := bufio.NewScanner(r)

	var row int
	var maxCol int

	for scanner.Scan() {
		line := scanner.Text()

		for col, char := range line {
			treesByLocation[location.Coordinate{Row: row, Col: col}] = Tree(char - '0')
			if col > maxCol {
				maxCol = col
			}
		}

		row++
	}

	maxRow := row - 1

	return treesByLocation, maxRow, maxCol
}

func CountVisibleTrees(r io.Reader) (int, error) {
	treesByLocation, maxRow, maxCol := parseTreeGrid(r)

	visible := map[location.Coordinate]bool{}

	for r := 0; r <= maxRow; r++ {
		maxHeight := Tree(-1)
		for c := 0; c <= maxCol; c++ {
			coord := location.Coordinate{Row: r, Col: c}
			height := treesByLocation[coord]
			if height > maxHeight {
				visible[coord] = true
				maxHeight = height
			}
		}

		maxHeight = Tree(-1)
		for c := maxCol; c >= 0; c-- {
			coord := location.Coordinate{Row: r, Col: c}
			height := treesByLocation[coord]
			if height > maxHeight {
				visible[coord] = true
				maxHeight = height
			}
		}
	}

	for c := 0; c <= maxCol; c++ {
		maxHeight := Tree(-1)
		for r := 0; r <= maxRow; r++ {
			coord := location.Coordinate{Row: r, Col: c}
			height := treesByLocation[coord]
			if height > maxHeight {
				visible[coord] = true
				maxHeight = height
			}
		}

		maxHeight = Tree(-1)
		for r := maxRow; r >= 0; r-- {
			coord := location.Coordinate{Row: r, Col: c}
			height := treesByLocation[coord]
			if height > maxHeight {
				visible[coord] = true
				maxHeight = height
			}
		}
	}

	return len(visible), nil
}

func GetMaxScenicScore(r io.Reader) (int, error) {
	treesByLocation, maxRow, maxCol := parseTreeGrid(r)
	maxScore := 0

	for r := 0; r <= maxRow; r++ {
		for c := 0; c <= maxCol; c++ {
			coord := location.Coordinate{Row: r, Col: c}
			height := treesByLocation[coord]

			// Calculate viewing distance in each direction
			upScore := 0
			for row := r - 1; row >= 0; row-- {
				upScore++
				if treesByLocation[location.Coordinate{Row: row, Col: c}] >= height {
					break
				}
			}

			downScore := 0
			for row := r + 1; row <= maxRow; row++ {
				downScore++
				if treesByLocation[location.Coordinate{Row: row, Col: c}] >= height {
					break
				}
			}

			leftScore := 0
			for col := c - 1; col >= 0; col-- {
				leftScore++
				if treesByLocation[location.Coordinate{Row: r, Col: col}] >= height {
					break
				}
			}

			rightScore := 0
			for col := c + 1; col <= maxCol; col++ {
				rightScore++
				if treesByLocation[location.Coordinate{Row: r, Col: col}] >= height {
					break
				}
			}

			scenicScore := upScore * downScore * leftScore * rightScore
			if scenicScore > maxScore {
				maxScore = scenicScore
			}
		}
	}

	return maxScore, nil
}
