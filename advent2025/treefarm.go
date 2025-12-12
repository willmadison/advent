package advent2025

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/willmadison/advent/internal/location"
)

type presentShape struct {
	index         int
	occupiedCells map[location.Coordinate]struct{}
}

type giftRegion struct {
	dimensions             location.Coordinate
	giftShapeQuotasByIndex map[int]int
}

func (g giftRegion) canFitAllShapesInQuota(presentShapesByIndex map[int]presentShape) bool {
	type shapeWithSize struct {
		shape presentShape
		size  int
	}

	var shapesWithSizes []shapeWithSize
	totalShapeArea := 0
	for idx, count := range g.giftShapeQuotasByIndex {
		shape := presentShapesByIndex[idx]
		shapeArea := len(shape.occupiedCells)
		for i := 0; i < count; i++ {
			shapesWithSizes = append(shapesWithSizes, shapeWithSize{shape, shapeArea})
			totalShapeArea += shapeArea
		}
	}

	gridArea := g.dimensions.Row * g.dimensions.Col
	if totalShapeArea > gridArea {
		return false
	}

	slices.SortFunc(shapesWithSizes, func(a, b shapeWithSize) int {
		return b.size - a.size
	})

	var shapesToPlace []presentShape
	for _, sws := range shapesWithSizes {
		shapesToPlace = append(shapesToPlace, sws.shape)
	}

	grid := make([][]bool, g.dimensions.Row)
	for i := range grid {
		grid[i] = make([]bool, g.dimensions.Col)
	}

	return tryPlaceShapes(grid, shapesToPlace, presentShapesByIndex, 0, gridArea-totalShapeArea)
}

func tryPlaceShapes(grid [][]bool, shapesToPlace []presentShape, allShapes map[int]presentShape, shapeIndex int, remainingFreeArea int) bool {
	if shapeIndex >= len(shapesToPlace) {
		return true
	}

	shape := shapesToPlace[shapeIndex]
	rotations := getUniqueRotations(shape)

	startRow, startCol := findFirstEmptyCell(grid)
	if startRow == -1 {
		return false
	}

	for row := startRow; row < len(grid); row++ {
		colStart := 0
		if row == startRow {
			colStart = startCol
		}
		for col := colStart; col < len(grid[0]); col++ {
			if !grid[row][col] {
				for _, rotated := range rotations {
					if canPlaceShape(grid, rotated, row, col) {
						placeShape(grid, rotated, row, col, true)
						if tryPlaceShapes(grid, shapesToPlace, allShapes, shapeIndex+1, remainingFreeArea) {
							return true
						}
						placeShape(grid, rotated, row, col, false)
					}
				}
			}
		}
	}

	return false
}

func findFirstEmptyCell(grid [][]bool) (int, int) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if !grid[row][col] {
				return row, col
			}
		}
	}
	return -1, -1
}

func getUniqueRotations(shape presentShape) []presentShape {
	rotations := []presentShape{shape}
	seen := make(map[string]struct{})
	seen[shapeSignature(shape)] = struct{}{}

	current := shape
	for i := 0; i < 3; i++ {
		current = rotate90(current)
		sig := shapeSignature(current)
		if _, exists := seen[sig]; !exists {
			rotations = append(rotations, current)
			seen[sig] = struct{}{}
		}
	}

	return rotations
}

func shapeSignature(shape presentShape) string {
	var coords []location.Coordinate
	for coord := range shape.occupiedCells {
		coords = append(coords, coord)
	}

	slices.SortFunc(coords, func(a, b location.Coordinate) int {
		if a.Row != b.Row {
			return a.Row - b.Row
		}
		return a.Col - b.Col
	})

	var sig strings.Builder
	for _, coord := range coords {
		sig.WriteString(strconv.Itoa(coord.Row))
		sig.WriteString(",")
		sig.WriteString(strconv.Itoa(coord.Col))
		sig.WriteString(";")
	}
	return sig.String()
}

func getAllRotations(shape presentShape) []presentShape {
	rotations := []presentShape{shape}

	current := shape
	for i := 0; i < 3; i++ {
		current = rotate90(current)
		rotations = append(rotations, current)
	}

	return rotations
}

func rotate90(shape presentShape) presentShape {
	newCells := make(map[location.Coordinate]struct{})
	for coord := range shape.occupiedCells {
		newCoord := location.Coordinate{Row: coord.Col, Col: -coord.Row}
		newCells[newCoord] = struct{}{}
	}

	minRow, minCol := 1000000, 1000000
	for coord := range newCells {
		if coord.Row < minRow {
			minRow = coord.Row
		}
		if coord.Col < minCol {
			minCol = coord.Col
		}
	}

	normalized := make(map[location.Coordinate]struct{})
	for coord := range newCells {
		normalized[location.Coordinate{
			Row: coord.Row - minRow,
			Col: coord.Col - minCol,
		}] = struct{}{}
	}

	return presentShape{
		index:         shape.index,
		occupiedCells: normalized,
	}
}

func canPlaceShape(grid [][]bool, shape presentShape, startRow, startCol int) bool {
	for coord := range shape.occupiedCells {
		row := startRow + coord.Row
		col := startCol + coord.Col

		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
			return false
		}

		if grid[row][col] {
			return false
		}
	}
	return true
}

func placeShape(grid [][]bool, shape presentShape, startRow, startCol int, place bool) {
	for coord := range shape.occupiedCells {
		row := startRow + coord.Row
		col := startCol + coord.Col
		grid[row][col] = place
	}
}

func CountGiftRegionFits(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)

	presentShapesByIndex := make(map[int]presentShape)
	var giftRegions []giftRegion

	var currentShapeIndex *int
	var currentShapeLines []string

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if currentShapeIndex != nil {
				shape := parseShape(*currentShapeIndex, currentShapeLines)
				presentShapesByIndex[*currentShapeIndex] = shape
				currentShapeIndex = nil
				currentShapeLines = nil
			}
			continue
		}

		if strings.HasSuffix(line, ":") && len(line) <= 3 {
			if currentShapeIndex != nil {
				shape := parseShape(*currentShapeIndex, currentShapeLines)
				presentShapesByIndex[*currentShapeIndex] = shape
			}

			idx, _ := strconv.Atoi(strings.TrimSuffix(line, ":"))
			currentShapeIndex = &idx
			currentShapeLines = nil
			continue
		}

		if strings.Contains(line, "x") && strings.Contains(line, ":") {
			parts := strings.Split(line, ":")
			dimParts := strings.Split(parts[0], "x")
			cols, _ := strconv.Atoi(dimParts[0])
			rows, _ := strconv.Atoi(dimParts[1])

			quotas := make(map[int]int)
			counts := strings.Fields(parts[1])
			for i, countStr := range counts {
				count, _ := strconv.Atoi(countStr)
				if count > 0 {
					quotas[i] = count
				}
			}

			giftRegions = append(giftRegions, giftRegion{
				dimensions:             location.Coordinate{Row: rows, Col: cols},
				giftShapeQuotasByIndex: quotas,
			})
			continue
		}

		if currentShapeIndex != nil {
			currentShapeLines = append(currentShapeLines, line)
		}
	}

	if currentShapeIndex != nil {
		shape := parseShape(*currentShapeIndex, currentShapeLines)
		presentShapesByIndex[*currentShapeIndex] = shape
	}

	fitCount := 0
	for _, region := range giftRegions {
		if region.canFitAllShapesInQuota(presentShapesByIndex) {
			fitCount++
		}
	}

	return fitCount, nil
}

func parseShape(index int, lines []string) presentShape {
	occupiedCells := make(map[location.Coordinate]struct{})

	for row, line := range lines {
		for col, ch := range line {
			if ch == '#' {
				occupiedCells[location.Coordinate{Row: row, Col: col}] = struct{}{}
			}
		}
	}

	return presentShape{
		index:         index,
		occupiedCells: occupiedCells,
	}
}
