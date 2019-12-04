package advent2019

import (
	"bufio"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/willmadison/advent"
)

type path struct {
	stepCountsByPoint map[advent.Point]int
	steps             []advent.Point
}

func (p path) stepsTo(point advent.Point) int {
	return p.stepCountsByPoint[point]
}

func (p path) intersect(other path) []advent.Point {
	var intersection []advent.Point
	var origin advent.Point

	for point, _ := range other.stepCountsByPoint {
		if _, present := p.stepCountsByPoint[point]; present && point != origin {
			intersection = append(intersection, point)
		}
	}

	return intersection
}

func FindNearestIntersection(input io.Reader) advent.Point {
	firstPath, secondPath := derivePaths(input)
	intersections := firstPath.intersect(secondPath)

	var origin advent.Point
	sort.Slice(intersections, func(i, j int) bool {
		return intersections[i].ManhattanDistance(origin) < intersections[j].ManhattanDistance(origin)
	})

	return intersections[0]
}

func FindMinimalTotalSteps(input io.Reader) int {
	firstPath, secondPath := derivePaths(input)

	intersections := firstPath.intersect(secondPath)

	minSteps := math.MaxInt32

	for _, intersection := range intersections {
		steps := firstPath.stepsTo(intersection)
		steps += secondPath.stepsTo(intersection)

		if steps < minSteps {
			minSteps = steps
		}
	}

	return minSteps
}

func derivePaths(r io.Reader) (path, path) {
	rawPaths := extractRawPaths(r)
	return derivePath(rawPaths[0]), derivePath(rawPaths[1])
}

func extractRawPaths(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	var rawPaths []string

	for scanner.Scan() {
		rawPaths = append(rawPaths, scanner.Text())
	}

	return rawPaths
}

type cardinalDirection rune

const (
	north cardinalDirection = 'U'
	east  cardinalDirection = 'R'
	south cardinalDirection = 'D'
	west  cardinalDirection = 'L'
)

type vector struct {
	direction cardinalDirection
	magnitude int
}

func toVector(s string) vector {
	direction := cardinalDirection(s[0])
	magnitiude, _ := strconv.Atoi(s[1:])

	return vector{direction, magnitiude}
}

func derivePath(rawPath string) path {
	movements := strings.Split(rawPath, ",")

	steps := []advent.Point{}
	p := map[advent.Point]int{}

	var currentLocation advent.Point

	for _, rawMovement := range movements {
		movement := toVector(rawMovement)

		switch movement.direction {
		case north:
			for y := currentLocation.Y; y < currentLocation.Y+movement.magnitude; y++ {
				point := advent.Point{X: currentLocation.X, Y: y}
				steps = append(steps, point)
			}

			currentLocation.Y += movement.magnitude
		case east:
			for x := currentLocation.X; x < currentLocation.X+movement.magnitude; x++ {
				point := advent.Point{X: x, Y: currentLocation.Y}
				steps = append(steps, point)
			}

			currentLocation.X += movement.magnitude
		case south:
			for y := currentLocation.Y; y > currentLocation.Y-movement.magnitude; y-- {
				point := advent.Point{X: currentLocation.X, Y: y}
				steps = append(steps, point)
			}

			currentLocation.Y -= movement.magnitude
		case west:
			for x := currentLocation.X; x > currentLocation.X-movement.magnitude; x-- {
				point := advent.Point{X: x, Y: currentLocation.Y}
				steps = append(steps, point)
			}

			currentLocation.X -= movement.magnitude
		}
	}

	for count, step := range steps {
		if _, present := p[step]; !present {
			p[step] = count
		}
	}

	return path{stepCountsByPoint: p, steps: steps}
}
