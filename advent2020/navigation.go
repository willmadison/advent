package advent2020

import (
	"bufio"
	"io"
	"strconv"

	"github.com/willmadison/advent/internal/location"
)

type Action string

const (
	North   = Action("N")
	South   = Action("S")
	East    = Action("E")
	West    = Action("W")
	Left    = Action("L")
	Right   = Action("R")
	Forward = Action("F")
)

type Orientation int

const (
	Unknown Orientation = iota
	Northbound
	Eastbound
	Southbound
	Westbound
)

type NavigationInstruction struct {
	Action Action
	Value  int
}

type Ship struct {
	Location    location.Point
	Orientation Orientation
	Waypoint    location.Point
}

func (s *Ship) Navigate(instructions []NavigationInstruction) {
	if s.Orientation == Unknown {
		s.Orientation = Eastbound
	}

	for _, instruction := range instructions {
		switch instruction.Action {
		case North:
			s.Location.Y += instruction.Value
		case East:
			s.Location.X += instruction.Value
		case South:
			s.Location.Y -= instruction.Value
		case West:
			s.Location.X -= instruction.Value
		case Forward:
			switch s.Orientation {
			case Northbound:
				s.Location.Y += instruction.Value
			case Eastbound:
				s.Location.X += instruction.Value
			case Southbound:
				s.Location.Y -= instruction.Value
			case Westbound:
				s.Location.X -= instruction.Value
			}
		case Left:
			turns := instruction.Value / 90
			var i int

			for i < turns {
				switch s.Orientation {
				case Northbound:
					s.Orientation = Westbound
				case Eastbound:
					s.Orientation = Northbound
				case Southbound:
					s.Orientation = Eastbound
				case Westbound:
					s.Orientation = Southbound
				}
				i++
			}

		case Right:
			turns := instruction.Value / 90
			var i int

			for i < turns {
				switch s.Orientation {
				case Northbound:
					s.Orientation = Eastbound
				case Eastbound:
					s.Orientation = Southbound
				case Southbound:
					s.Orientation = Westbound
				case Westbound:
					s.Orientation = Northbound
				}
				i++
			}
		}
	}
}

func (s *Ship) NavigateByWaypoint(instructions []NavigationInstruction) {
	if s.Orientation == Unknown {
		s.Orientation = Eastbound
		s.Waypoint = location.Point{X: 10, Y: 1}
	}

	for _, instruction := range instructions {
		switch instruction.Action {
		case North:
			s.Waypoint.Y += instruction.Value
		case East:
			s.Waypoint.X += instruction.Value
		case South:
			s.Waypoint.Y -= instruction.Value
		case West:
			s.Waypoint.X -= instruction.Value
		case Forward:
			s.Location.X += instruction.Value * s.Waypoint.X
			s.Location.Y += instruction.Value * s.Waypoint.Y
		case Left:
			turns := instruction.Value / 90
			var i int

			for i < turns {
				s.Waypoint.Rotate90(location.Counterclockwise)
				i++
			}

		case Right:
			turns := instruction.Value / 90
			var i int

			for i < turns {
				s.Waypoint.Rotate90(location.Clockwise)
				i++
			}
		}
	}

}

func ParseNavigationInstructions(r io.Reader) []NavigationInstruction {
	var instructions []NavigationInstruction

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		instruction := parseInstruction(scanner.Text())
		instructions = append(instructions, instruction)
	}

	return instructions
}

func parseInstruction(rawInstruction string) NavigationInstruction {
	action := rawInstruction[0:1]
	value, _ := strconv.Atoi(rawInstruction[1:])

	return NavigationInstruction{Action: Action(action), Value: value}
}
