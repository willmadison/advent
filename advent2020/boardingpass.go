package advent2020

import (
	"bufio"
	"io"
	"sort"
)

func FindSeatLocation(boardingPass string) (int, int) {
	var row, column int

	rowMin := 0
	rowMax := 127

	rowLocator := boardingPass[0:7]

	for _, c := range rowLocator {
		midpoint := (rowMin + rowMax) / 2
		switch rune(c) {
		case 'F':
			rowMax = midpoint
		default:
			rowMin = midpoint + 1
		}
	}

	row = rowMax

	columnMin := 0
	columnMax := 7

	columnLocator := boardingPass[7:]

	for _, c := range columnLocator {
		midpoint := (columnMin + columnMax) / 2
		switch rune(c) {
		case 'L':
			columnMax = midpoint
		default:
			columnMin = midpoint + 1
		}
	}

	column = columnMin

	return row, column
}

func FindMaxSeatID(r io.Reader) int {
	var maxSeatID int

	allSeatIDs := findAllSeatIDs(r)

	for _, seatID := range allSeatIDs {
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	return maxSeatID
}

func FindMissingSeatID(r io.Reader) int {
	var missingSeatID int

	allSeatIDs := findAllSeatIDs(r)

	sort.Ints(allSeatIDs)

	for i := 0; i < len(allSeatIDs)-1; i++ {
		nextSeatId := allSeatIDs[i] + 1
		if allSeatIDs[i+1] != nextSeatId {
			missingSeatID = nextSeatId
			break
		}
	}

	return missingSeatID
}

func findAllSeatIDs(r io.Reader) []int {
	var seatIDs []int

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		boardingPass := scanner.Text()
		row, col := FindSeatLocation(boardingPass)
		seatID := row*8 + col

		seatIDs = append(seatIDs, seatID)
	}

	return seatIDs
}
