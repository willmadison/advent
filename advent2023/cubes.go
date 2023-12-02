package advent2023

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Game struct {
	Number int
	Rounds []Round
}

type Round struct {
	Cubes        []Cube
	CubesByColor map[Color][]Cube
}

func NewRound(cubes []Cube) Round {
	round := Round{Cubes: cubes}
	round.CubesByColor = map[Color][]Cube{}

	for _, c := range round.Cubes {

		if _, seen := round.CubesByColor[c.Color]; !seen {
			round.CubesByColor[c.Color] = []Cube{}
		}

		round.CubesByColor[c.Color] = append(round.CubesByColor[c.Color], c)
	}

	return round
}

type Color int

const (
	Red Color = iota
	Blue
	Green
)

var colorsByName = map[string]Color{
	"red":   Red,
	"blue":  Blue,
	"green": Green,
}

type Cube struct {
	Color Color
}

type Constraint func(Game) bool

func LimitedColors(maxRed, maxBlue, maxGreen int) Constraint {
	return func(game Game) bool {
		allRoundsInColorBound := true

		for _, round := range game.Rounds {
			if len(round.CubesByColor[Red]) > maxRed ||
				len(round.CubesByColor[Blue]) > maxBlue ||
				len(round.CubesByColor[Green]) > maxGreen {
				allRoundsInColorBound = false
				break
			}
		}

		return allRoundsInColorBound
	}
}

func FindPossibleGames(r io.Reader, constraints ...Constraint) []Game {
	games := parseGames(r)

	var possibleGames []Game

	for _, game := range games {
		possible := true

		for _, constraint := range constraints {
			possible = constraint(game)

			if !possible {
				break
			}
		}

		if !possible {
			continue
		}

		possibleGames = append(possibleGames, game)
	}

	return possibleGames
}

func FindMinimumViableGameConfigs(r io.Reader) []map[Color]int {
	games := parseGames(r)

	configs := []map[Color]int{}

	for _, game := range games {
		configs = append(configs, determineMVGC(game))
	}

	return configs
}

func determineMVGC(game Game) map[Color]int {
	maxRed, maxBlue, maxGreen := len(game.Rounds[0].CubesByColor[Red]), len(game.Rounds[0].CubesByColor[Blue]), len(game.Rounds[0].CubesByColor[Green])

	for _, round := range game.Rounds {
		if len(round.CubesByColor[Red]) > maxRed {
			maxRed = len(round.CubesByColor[Red])
		}

		if len(round.CubesByColor[Blue]) > maxBlue {
			maxBlue = len(round.CubesByColor[Blue])
		}

		if len(round.CubesByColor[Green]) > maxGreen {
			maxGreen = len(round.CubesByColor[Green])
		}
	}

	config := map[Color]int{
		Red:   maxRed,
		Blue:  maxBlue,
		Green: maxGreen,
	}

	return config
}

func parseGames(r io.Reader) []Game {
	scanner := bufio.NewScanner(r)

	var games []Game

	for scanner.Scan() {
		game := parseGameLine(scanner.Text())
		games = append(games, game)
	}

	return games
}

func parseGameLine(line string) Game {
	game := Game{}

	rawGameParts := strings.Split(line, ": ")

	rawNumber := strings.TrimPrefix(rawGameParts[0], "Game ")
	game.Number, _ = strconv.Atoi(rawNumber)

	game.Rounds = parseRounds(rawGameParts[1])

	return game
}

func parseRounds(rawRounds string) []Round {
	var rounds []Round

	rawIndividualRounds := strings.Split(rawRounds, "; ")

	for _, rawRound := range rawIndividualRounds {
		rounds = append(rounds, NewRound(parseCubes(rawRound)))
	}

	return rounds
}

func parseCubes(rawCubes string) []Cube {
	var cubes []Cube

	rawCubeParts := strings.Split(rawCubes, ", ")

	for _, cubePart := range rawCubeParts {
		rawCubeMeta := strings.Fields(cubePart)

		count, _ := strconv.Atoi(rawCubeMeta[0])
		color := colorsByName[rawCubeMeta[1]]

		for i := 0; i < count; i++ {
			cubes = append(cubes, Cube{Color: color})
		}
	}

	return cubes
}
