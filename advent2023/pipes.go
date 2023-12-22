package advent2023

import (
	"bufio"
	"io"

	"golang.org/x/exp/slices"

	"github.com/willmadison/advent/internal/location"
)

type Pipe int

type PipeLocale struct {
	pipe     Pipe
	location location.Point
}

const (
	Vertical Pipe = iota
	Horizontal
	NortheastConnector
	NorthwestConnector
	SoutheastConnector
	SouthwestConnector
)

var pipesByRune = map[rune]Pipe{
	'|': Vertical,
	'-': Horizontal,
	'L': NortheastConnector,
	'J': NorthwestConnector,
	'F': SoutheastConnector,
	'7': SouthwestConnector,
}

var directionsByPipe = map[Pipe][]location.CardinalDirection{
	Vertical:           {location.North, location.South},
	Horizontal:         {location.East, location.West},
	NortheastConnector: {location.North, location.East},
	NorthwestConnector: {location.North, location.West},
	SoutheastConnector: {location.South, location.East},
	SouthwestConnector: {location.South, location.West},
}

type Situation struct {
	PipeLocales      []PipeLocale
	PipesByLocation  map[location.Point]PipeLocale
	StartingPosition location.Point
	CurrentPosition  location.Point
	Orientation      location.CardinalDirection
	Steps            int
}

func (s *Situation) Advance() {
	for _, v := range location.Vectors {
		if v.Direction == s.Orientation.Reverse() {
			continue
		}

		var waysForward []location.CardinalDirection

		if s.CurrentPosition != s.StartingPosition {
			waysForward = directionsByPipe[s.PipesByLocation[s.CurrentPosition].pipe]
		} else {
			waysForward = []location.CardinalDirection{
				location.North,
				location.East,
				location.South,
				location.West,
			}
		}

		if s.CurrentPosition != s.StartingPosition && !slices.Contains(waysForward, v.Direction) {
			continue
		}

		nextPosition := s.CurrentPosition.Add(v.Point)
		nextPipeLocale := s.PipesByLocation[nextPosition]
		waysForwardFromNext := directionsByPipe[nextPipeLocale.pipe]

		if slices.Contains(waysForwardFromNext, v.Direction.Reverse()) {
			s.Orientation = v.Direction
			s.CurrentPosition = nextPosition
			s.Steps++
			return
		}
	}
}

func DistanceToFurthestPipeFromStart(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	var startingLocation location.Point

	row, col := 0, 0

	var pipeLocales []PipeLocale
	pipesByLocation := map[location.Point]PipeLocale{}

	for scanner.Scan() {
		line := scanner.Text()

		col = 0

		for _, c := range line {
			switch c {
			case '.':
				col++
				continue
			case 'S':
				startingLocation = location.Point{X: row, Y: col}
			default:
				p := pipesByRune[c]
				pipeLocale := PipeLocale{pipe: p, location: location.Point{X: row, Y: col}}
				pipesByLocation[pipeLocale.location] = pipeLocale
				pipeLocales = append(pipeLocales, pipeLocale)
			}

			col++
		}

		row++
	}

	var startingVector location.Vector

	for _, v := range location.Vectors {
		nextLocale := startingLocation.Add(v.Point)
		pipeLocale, ok := pipesByLocation[nextLocale]

		if ok && slices.Contains(directionsByPipe[pipeLocale.pipe], v.Direction.Reverse()) {
			startingVector = v
			break
		}

	}

	situation := Situation{
		PipeLocales:      pipeLocales,
		PipesByLocation:  pipesByLocation,
		StartingPosition: startingLocation,
		CurrentPosition:  startingLocation,
		Orientation:      startingVector.Direction,
	}

	for {
		situation.Advance()

		if situation.CurrentPosition == situation.StartingPosition {
			break
		}
	}

	return situation.Steps / 2
}
