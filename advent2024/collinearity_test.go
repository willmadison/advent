package advent2024_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2024"
)

func TestFindAntennae(t *testing.T) {
	given := strings.NewReader(`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`)

	antennae, dimensions := advent2024.FindAntennae(given)

	assert.Equal(t, 2, len(antennae))
	assert.Equal(t, 12, dimensions.Row)
	assert.Equal(t, 12, dimensions.Col)
}

func TestFindAntinodes(t *testing.T) {
	given := strings.NewReader(`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`)

	antennae, dimensions := advent2024.FindAntennae(given)

	antinodes := advent2024.FindAntinodes(antennae, dimensions)

	assert.Equal(t, 14, len(antinodes))
}

func TestFindResonantAntinodes(t *testing.T) {
	given := strings.NewReader(`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`)

	antennae, dimensions := advent2024.FindAntennae(given)

	antinodes := advent2024.FindResonantAntinodes(antennae, dimensions)

	assert.Equal(t, 34, len(antinodes))
}
