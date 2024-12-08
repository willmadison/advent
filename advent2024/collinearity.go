package advent2024

import (
	"bufio"
	"io"

	"github.com/willmadison/advent/internal/location"
)

type Antenna struct {
	Frequency rune
	Location  location.Coordinate
}

func FindAntennae(r io.Reader) (map[rune][]Antenna, location.Coordinate) {
	antennae := map[rune][]Antenna{}

	var row, columns int

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		if columns == 0 {
			columns = len(line)
		}

		for col, char := range line {
			switch char {
			case '.':
			default:
				frequency := rune(char)
				if _, present := antennae[frequency]; !present {
					antennae[rune(char)] = []Antenna{}
				}

				location := location.Coordinate{Row: row, Col: col}
				antenna := Antenna{Frequency: frequency, Location: location}

				antennae[frequency] = append(antennae[frequency], antenna)
			}
		}

		row++
	}

	return antennae, location.Coordinate{Row: row, Col: columns}
}

func FindAntinodes(antennae map[rune][]Antenna, dimensions location.Coordinate) map[location.Coordinate]struct{} {
	antinodes := map[location.Coordinate]struct{}{}

	for _, antennas := range antennae {
		for i, antenna := range antennas {
			for j := 0; j < len(antennas); j++ {
				if j != i {
					delta := antenna.Location.Delta(antennas[j].Location)
					antinode := location.Coordinate{Row: antenna.Location.Row + delta.Row, Col: antenna.Location.Col + delta.Col}

					if antinode.InBounds(dimensions.Row, dimensions.Col) {
						if _, present := antinodes[antinode]; !present {
							antinodes[antinode] = struct{}{}
						}
					}
				}
			}
		}
	}

	return antinodes
}

func FindResonantAntinodes(antennae map[rune][]Antenna, dimensions location.Coordinate) map[location.Coordinate]struct{} {
	antinodes := map[location.Coordinate]struct{}{}

	for _, antennas := range antennae {
		for i, antenna := range antennas {
			for j := 0; j < len(antennas); j++ {
				if j != i {
					delta := antenna.Location.Delta(antennas[j].Location)
					antinode := location.Coordinate{Row: antenna.Location.Row, Col: antenna.Location.Col}

					for antinode.InBounds(dimensions.Row, dimensions.Col) {
						if _, present := antinodes[antinode]; !present {
							antinodes[antinode] = struct{}{}
						}

						antinode = location.Coordinate{Row: antinode.Row + delta.Row, Col: antinode.Col + delta.Col}
					}
				}
			}
		}
	}

	return antinodes
}
