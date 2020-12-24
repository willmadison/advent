package advent2020

import (
	"bufio"
	"io"

	"github.com/willmadison/advent/internal/location"
)

type HexDirection string

const (
	Northward HexDirection = "n"
	Northeast              = "ne"
	Eastward               = "e"
	Southeast              = "se"
	Southward              = "s"
	Southwest              = "sw"
	Westward               = "w"
	Northwest              = "nw"
)

func ParseDirections(value string) []HexDirection {
	var directions []HexDirection

	for i := 0; i < len(value); {
		var direction HexDirection

		switch rune(value[i]) {
		case 'w':
			direction = Westward
			i++
		case 'e':
			direction = Eastward
			i++
		case 's':
			if i+1 < len(value) {
				next := rune(value[i+1])

				if next == 'e' {
					direction = Southeast
					i += 2
				} else if next == 'w' {
					direction = Southwest
					i += 2
				} else {
					direction = Southward
					i++
				}
			}
		case 'n':
			if i+1 < len(value) {
				next := rune(value[i+1])

				if next == 'e' {
					direction = Northeast
					i += 2
				} else if next == 'w' {
					direction = Northwest
					i += 2
				} else {
					direction = Northward
					i++
				}
			}
		}

		directions = append(directions, direction)
	}

	return directions
}

type Floor struct {
	AllDirections [][]HexDirection
	blackTiles    map[location.Point]struct{}
}

func (f *Floor) Follow(directions []HexDirection) {
	var location location.Point

	for _, d := range directions {
		switch d {
		case Northward, Northeast:
			location.Y++
		case Eastward:
			location.X++
		case Southeast:
			location.X++
			location.Y--
		case Southward, Southwest:
			location.Y--
		case Westward:
			location.X--
		case Northwest:
			location.X--
			location.Y++
		}
	}

	if _, isBlack := f.blackTiles[location]; isBlack {
		delete(f.blackTiles, location)
	} else {
		f.blackTiles[location] = struct{}{}
	}
}

func (f Floor) GetBlackCount() int {
	return len(f.blackTiles)
}

func (f *Floor) Rotate() {
	tiles := map[location.Point]struct{}{}

	for tile := range f.blackTiles {
		tiles[tile] = struct{}{}

		for _, n := range tile.Neighbors() {
			if _, present := tiles[n]; !present {
				tiles[n] = struct{}{}
			}
		}
	}

	toBeFlippedBlack := map[location.Point]struct{}{}
	toBeFlippedWhite := map[location.Point]struct{}{}

	for tile := range tiles {
		var blackNeighbors int

		for _, n := range tile.Neighbors() {
			if _, isBlack := f.blackTiles[n]; isBlack {
				blackNeighbors++
			}
		}

		if _, isBlack := f.blackTiles[tile]; isBlack {
			if blackNeighbors == 0 || blackNeighbors > 2 {
				toBeFlippedWhite[tile] = struct{}{}
			}
		} else {
			if blackNeighbors == 2 {
				toBeFlippedBlack[tile] = struct{}{}
			}
		}
	}

	for blackFlip := range toBeFlippedBlack {
		f.blackTiles[blackFlip] = struct{}{}
	}

	for whiteFlip := range toBeFlippedWhite {
		delete(f.blackTiles, whiteFlip)
	}
}

func ParseAllDirections(r io.Reader) Floor {
	var floor Floor
	floor.blackTiles = map[location.Point]struct{}{}

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		directions := ParseDirections(scanner.Text())
		floor.AllDirections = append(floor.AllDirections, directions)
	}

	return floor
}
