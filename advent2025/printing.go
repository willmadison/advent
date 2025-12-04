package advent2025

import (
	"bufio"
	"io"

	"github.com/willmadison/advent/internal/location"
)

func parsePaperRollLocales(r io.Reader) (map[location.Coordinate]struct{}, error) {
	var row int

	scanner := bufio.NewScanner(r)

	paperRollsByLocation := map[location.Coordinate]struct{}{}

	for scanner.Scan() {
		rawRow := scanner.Text()

		for col, v := range rawRow {
			value := rune(v)

			if value == '.' {
				continue
			}

			switch value {
			case '@':
				paperRollsByLocation[location.Coordinate{Row: row, Col: col}] = struct{}{}
			}
		}

		row++
	}

	return paperRollsByLocation, nil
}

func FindAccessiblePaperRolls(r io.Reader) ([]location.Coordinate, error) {
	paperRollsByLocation, err := parsePaperRollLocales(r)

	if err != nil {
		return nil, err
	}

	return doFindAccessiblePaperRolls(paperRollsByLocation)
}

func doFindAccessiblePaperRolls(paperRollsByLocation map[location.Coordinate]struct{}) ([]location.Coordinate, error) {
	var accessiblePaperRolls []location.Coordinate

	for paperRoll := range paperRollsByLocation {
		var numNeighboringRolls int

		for _, neighbor := range paperRoll.Neighbors() {
			if _, present := paperRollsByLocation[neighbor]; present {
				numNeighboringRolls++
			}
		}

		if numNeighboringRolls < 4 {
			accessiblePaperRolls = append(accessiblePaperRolls, paperRoll)
		}
	}

	return accessiblePaperRolls, nil
}

func RemoveAllAccessiblePaperRolls(r io.Reader) (int, error) {
	paperRollsByLocation, err := parsePaperRollLocales(r)

	if err != nil {
		return 0, err
	}

	var paperRollsRemoved int

	accessiblePaperRolls, err := doFindAccessiblePaperRolls(paperRollsByLocation)

	if err != nil {
		return 0, err
	}

	for len(accessiblePaperRolls) > 0 {
		for _, accessiblePaperRoll := range accessiblePaperRolls {
			delete(paperRollsByLocation, accessiblePaperRoll)
			paperRollsRemoved++
		}

		accessiblePaperRolls, err = doFindAccessiblePaperRolls(paperRollsByLocation)

		if err != nil {
			return 0, err
		}
	}

	return paperRollsRemoved, nil
}
