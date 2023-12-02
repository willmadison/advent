package advent2023_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2023"
)

func TestFindPossibleGames(t *testing.T) {
	given := strings.NewReader(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`)

	expected := 3

	games := advent2023.FindPossibleGames(given, advent2023.LimitedColors(12, 14, 13))

	assert.Equal(t, expected, len(games))

	expected = 8

	var sum int

	for _, game := range games {
		sum += game.Number
	}

	assert.Equal(t, expected, sum)
}

func TestFindMinimumViableGameConfigs(t *testing.T) {
	given := strings.NewReader(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`)

	expected := 2286

	configs := advent2023.FindMinimumViableGameConfigs(given)

	var totalPower int

	for _, config := range configs {
		totalPower += config[advent2023.Red] * config[advent2023.Blue] * config[advent2023.Green]
	}

	assert.Equal(t, expected, totalPower)
}
