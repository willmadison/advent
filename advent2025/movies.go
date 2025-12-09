package advent2025

import (
	"bufio"
	"fmt"
	"io"

	"github.com/willmadison/advent/internal/containers"
	"github.com/willmadison/advent/internal/location"
)

func FindMaxiumAreaRectangle(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	var redTiles []location.Coordinate

	for scanner.Scan() {
		line := scanner.Text()

		coordinate, err := parseCoordinate(line)
		if err != nil {
			return 0, err
		}

		redTiles = append(redTiles, coordinate)
	}

	maxArea := 0

	for i := 0; i < len(redTiles); i++ {
		for j := i + 1; j < len(redTiles); j++ {
			c1, c2 := redTiles[i], redTiles[j]

			width := abs(c2.Col-c1.Col) + 1
			height := abs(c2.Row-c1.Row) + 1
			area := width * height

			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea, nil
}

func FindMaxiumAreaRectangleWithConstraints(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	var redTiles []location.Coordinate

	for scanner.Scan() {
		line := scanner.Text()

		coordinate, err := parseCoordinate(line)
		if err != nil {
			return 0, err
		}

		redTiles = append(redTiles, coordinate)
	}

	rowSet := make(map[int]bool)
	colSet := make(map[int]bool)

	for _, tile := range redTiles {
		rowSet[tile.Row] = true
		colSet[tile.Col] = true
	}

	rows := make([]int, 0, len(rowSet))
	for row := range rowSet {
		rows = append(rows, row)
	}
	cols := make([]int, 0, len(colSet))
	for col := range colSet {
		cols = append(cols, col)
	}

	sortInts(rows)
	sortInts(cols)

	rowToIdx := make(map[int]int)
	colToIdx := make(map[int]int)
	for i, row := range rows {
		rowToIdx[row] = i
	}
	for i, col := range cols {
		colToIdx[col] = i
	}

	compressedRows := len(rows)
	compressedCols := len(cols)
	compressedGrid := make([][]bool, compressedRows)
	for i := range compressedGrid {
		compressedGrid[i] = make([]bool, compressedCols)
	}

	for i := 0; i < len(redTiles); i++ {
		next := (i + 1) % len(redTiles)
		markPathInCompressedGrid(redTiles[i], redTiles[next], compressedGrid, rowToIdx, colToIdx)
	}

	fillCompressedExterior(compressedGrid, compressedRows, compressedCols)

	prefixSum := make([][]int, compressedRows+1)
	for i := range prefixSum {
		prefixSum[i] = make([]int, compressedCols+1)
	}

	for r := 1; r <= compressedRows; r++ {
		for c := 1; c <= compressedCols; c++ {
			val := 0
			if compressedGrid[r-1][c-1] {
				val = 1
			}
			prefixSum[r][c] = val + prefixSum[r-1][c] + prefixSum[r][c-1] - prefixSum[r-1][c-1]
		}
	}

	maxArea := 0

	sortedByRow := make([]location.Coordinate, len(redTiles))
	copy(sortedByRow, redTiles)
	sortCoordinates(sortedByRow)

	for i := 0; i < len(sortedByRow); i++ {
		c1 := sortedByRow[i]
		r1 := rowToIdx[c1.Row] + 1
		c1Idx := colToIdx[c1.Col] + 1

		for j := i + 1; j < len(sortedByRow); j++ {
			c2 := sortedByRow[j]

			theoreticalMaxArea := abs(c2.Row-c1.Row+1) * abs(c2.Col-c1.Col+1)
			if theoreticalMaxArea <= maxArea {
				continue
			}

			r2 := rowToIdx[c2.Row] + 1
			c2Idx := colToIdx[c2.Col] + 1

			minR := min(r1, r2)
			maxR := max(r1, r2)
			minC := min(c1Idx, c2Idx)
			maxC := max(c1Idx, c2Idx)

			compressedArea := (maxR - minR + 1) * (maxC - minC + 1)
			sum := prefixSum[maxR][maxC] - prefixSum[minR-1][maxC] - prefixSum[maxR][minC-1] + prefixSum[minR-1][minC-1]

			if sum == compressedArea {
				actualMinRow := rows[minR-1]
				actualMaxRow := rows[maxR-1]
				actualMinCol := cols[minC-1]
				actualMaxCol := cols[maxC-1]

				width := actualMaxCol - actualMinCol + 1
				height := actualMaxRow - actualMinRow + 1
				area := width * height

				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea, nil
}

func sortCoordinates(coords []location.Coordinate) {
	if len(coords) < 20 {
		for i := 1; i < len(coords); i++ {
			key := coords[i]
			j := i - 1
			for j >= 0 && (coords[j].Row > key.Row || (coords[j].Row == key.Row && coords[j].Col > key.Col)) {
				coords[j+1] = coords[j]
				j--
			}
			coords[j+1] = key
		}
		return
	}

	quickSortCoords(coords, 0, len(coords)-1)
}

func quickSortCoords(coords []location.Coordinate, low, high int) {
	if low < high {
		pi := partitionCoords(coords, low, high)
		quickSortCoords(coords, low, pi-1)
		quickSortCoords(coords, pi+1, high)
	}
}

func partitionCoords(coords []location.Coordinate, low, high int) int {
	pivot := coords[high]
	i := low - 1

	for j := low; j < high; j++ {
		if coords[j].Row < pivot.Row || (coords[j].Row == pivot.Row && coords[j].Col < pivot.Col) {
			i++
			coords[i], coords[j] = coords[j], coords[i]
		}
	}
	coords[i+1], coords[high] = coords[high], coords[i+1]
	return i + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseCoordinate(line string) (location.Coordinate, error) {
	var coord location.Coordinate

	_, err := fmt.Sscanf(line, "%d,%d", &coord.Col, &coord.Row)
	if err != nil {
		return location.Coordinate{}, err
	}

	return coord, nil
}

func sortInts(arr []int) {
	if len(arr) < 20 {
		for i := 1; i < len(arr); i++ {
			key := arr[i]
			j := i - 1
			for j >= 0 && arr[j] > key {
				arr[j+1] = arr[j]
				j--
			}
			arr[j+1] = key
		}
		return
	}

	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func markPathInCompressedGrid(start, end location.Coordinate, grid [][]bool, rowToIdx, colToIdx map[int]int) {
	startR, startC := rowToIdx[start.Row], colToIdx[start.Col]
	endR, endC := rowToIdx[end.Row], colToIdx[end.Col]

	if startR == endR {
		minC := min(startC, endC)
		maxC := max(startC, endC)
		for c := minC; c <= maxC; c++ {
			grid[startR][c] = true
		}
	} else if startC == endC {
		minR := min(startR, endR)
		maxR := max(startR, endR)
		for r := minR; r <= maxR; r++ {
			grid[r][startC] = true
		}
	}
}

func fillCompressedExterior(grid [][]bool, rows, cols int) {
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	type pos struct{ r, c int }
	queue := containers.NewQueue[pos]()

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if (r == 0 || r == rows-1 || c == 0 || c == cols-1) && !grid[r][c] {
				visited[r][c] = true
				queue.Enqueue(pos{r, c})
			}
		}
	}

	directions := []pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for queue.Size() > 0 {
		curr, _ := queue.Dequeue()

		for _, dir := range directions {
			nr, nc := curr.r+dir.r, curr.c+dir.c
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && !visited[nr][nc] && !grid[nr][nc] {
				visited[nr][nc] = true
				queue.Enqueue(pos{nr, nc})
			}
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if !visited[r][c] && !grid[r][c] {
				grid[r][c] = true
			}
		}
	}
}
