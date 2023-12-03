package advent2023

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"unicode"

	"github.com/willmadison/advent/internal/location"
)

type PartNumber struct {
	Number       int
	DigitLocales []location.Coordinate
}

func (p PartNumber) IsAdjacentToASymbol(symbolLocations map[location.Coordinate]struct{}) bool {
	for _, locale := range p.DigitLocales {
		for _, neighbor := range locale.Neighbors() {
			if _, present := symbolLocations[neighbor]; present {
				return true
			}
		}
	}

	return false
}

func (p PartNumber) IsAdjacentToLocation(coordinate location.Coordinate) bool {
	neighboringLocations := map[location.Coordinate]struct{}{}

	for _, locale := range p.DigitLocales {
		for _, neighbor := range locale.Neighbors() {
			neighboringLocations[neighbor] = struct{}{}
		}
	}

	_, adjacent := neighboringLocations[coordinate]

	return adjacent
}

func FindPartNumbers(r io.Reader) ([]PartNumber, map[location.Coordinate]struct{}) {
	var potentialPartNumbers []PartNumber

	symbolLocations := map[location.Coordinate]struct{}{}
	gearLocations := map[location.Coordinate]struct{}{}

	scanner := bufio.NewScanner(r)

	row := -1

	for scanner.Scan() {
		row++

		var buf bytes.Buffer
		digitCoordinates := []location.Coordinate{}

		for col, character := range scanner.Text() {
			switch {
			case unicode.IsDigit(character):
				buf.WriteRune(character)
				digitCoordinates = append(digitCoordinates, location.Coordinate{Row: row, Col: col})
			case character == '.':
				if buf.Len() > 0 {
					number, _ := strconv.Atoi(buf.String())
					partNumber := PartNumber{Number: number, DigitLocales: digitCoordinates}
					potentialPartNumbers = append(potentialPartNumbers, partNumber)

					buf.Reset()
					digitCoordinates = []location.Coordinate{}
				}
			default:
				if buf.Len() > 0 {
					number, _ := strconv.Atoi(buf.String())
					partNumber := PartNumber{Number: number, DigitLocales: digitCoordinates}
					potentialPartNumbers = append(potentialPartNumbers, partNumber)

					buf.Reset()
					digitCoordinates = []location.Coordinate{}
				}
				coordinate := location.Coordinate{Row: row, Col: col}

				symbolLocations[coordinate] = struct{}{}

				if character == '*' {
					gearLocations[coordinate] = struct{}{}
				}
			}
		}

		if buf.Len() > 0 {
			number, _ := strconv.Atoi(buf.String())
			partNumber := PartNumber{Number: number, DigitLocales: digitCoordinates}
			potentialPartNumbers = append(potentialPartNumbers, partNumber)
		}
	}

	var partNumbers []PartNumber

	for _, p := range potentialPartNumbers {
		if p.IsAdjacentToASymbol(symbolLocations) {
			partNumbers = append(partNumbers, p)
		}
	}

	return partNumbers, gearLocations
}

func DetermineTotalGearRatio(partNumbers []PartNumber, gearLocations map[location.Coordinate]struct{}) int {
	var totalGearRatio int

	for loc := range gearLocations {
		adjacentParts := findAdjacentParts(partNumbers, loc)

		if len(adjacentParts) == 2 {
			totalGearRatio += adjacentParts[0].Number * adjacentParts[1].Number
		}
	}

	return totalGearRatio
}

func findAdjacentParts(partNumbers []PartNumber, coordinate location.Coordinate) []PartNumber {
	var adjacentParts []PartNumber

	for _, part := range partNumbers {
		if part.IsAdjacentToLocation(coordinate) {
			adjacentParts = append(adjacentParts, part)
		}
	}

	return adjacentParts
}
