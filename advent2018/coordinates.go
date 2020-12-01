package advent2018

import (
	"bufio"
	"errors"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"
)

type grid struct {
	coordinates []point
	min, max    point
}

func (g grid) largestFiniteArea() int {
	closestPointsByLandmark := map[point][]point{}

	for x := g.min.x; x <= g.max.x; x++ {
		for y := g.min.y; y <= g.max.y; y++ {
			a := point{x, y}
			closest, err := findClosestCoordinate(a, g.coordinates)

			if err != nil {
				continue
			}

			closestPointsByLandmark[closest] = append(closestPointsByLandmark[closest], a)
		}
	}

	finites := []point{}

	for _, p := range g.coordinates {
		if g.isFinite(closestPointsByLandmark[p]) {
			finites = append(finites, p)
		}
	}

	largestArea := math.MinInt64

	sort.Slice(finites, func(i, j int) bool {
		return len(closestPointsByLandmark[finites[i]]) > len(closestPointsByLandmark[finites[j]])
	})

	for _, f := range finites {
		if len(closestPointsByLandmark[f]) > largestArea {
			largestArea = len(closestPointsByLandmark[f])
		}
	}

	return largestArea
}

// a given point on the grid is considered finite if it non of it's closest points are on the boundary
func (g grid) isFinite(closestPoints []point) bool {
	for _, p := range closestPoints {
		if p.x == g.min.x || p.x == g.max.x || p.y == g.min.y || p.y == g.max.y {
			return false
		}
	}

	return true
}

func (g grid) regionMinimizedByConstraint(constraint int) int {
	region := []point{}

	for x := g.min.x; x <= g.max.x; x++ {
		for y := g.min.y; y <= g.max.y; y++ {
			var cummulativeDistance int

			a := point{x, y}

			for _, b := range g.coordinates {
				cummulativeDistance += a.manhattanDistance(b)
			}

			if cummulativeDistance < constraint {
				region = append(region, a)
			}
		}
	}

	return len(region)
}

func findClosestCoordinate(a point, coordinates []point) (point, error) {
	coordinatesByDistance := map[int][]point{}

	min := math.MaxInt64

	for _, c := range coordinates {
		d := a.manhattanDistance(c)

		if d < min {
			min = d
		}

		coordinatesByDistance[d] = append(coordinatesByDistance[d], c)
	}

	if len(coordinatesByDistance[min]) > 1 {
		return point{}, errors.New("no ties allowed")
	}

	return coordinatesByDistance[min][0], nil
}

type point struct {
	x, y int
}

func (p point) manhattanDistance(other point) int {
	return abs(p.x-other.x) + abs(p.y-other.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func FindLargestFiniteArea(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	var points []point

	for scanner.Scan() {
		points = append(points, parsePoint(scanner.Text()))
	}

	min, max := findMinMax(points)
	g := grid{points, min, max}

	return g.largestFiniteArea()
}

func FindRegionAreaMinimizedByConstraint(r io.Reader, constraint int) int {
	scanner := bufio.NewScanner(r)

	var points []point

	for scanner.Scan() {
		points = append(points, parsePoint(scanner.Text()))
	}

	min, max := findMinMax(points)
	g := grid{points, min, max}

	return g.regionMinimizedByConstraint(constraint)
}

func parsePoint(rawPoint string) point {
	parts := strings.Split(rawPoint, ", ")

	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])

	return point{x, y}
}

func findMinMax(points []point) (point, point) {
	var min, max point
	minx, miny, maxx, maxy := math.MaxInt64, math.MaxInt64, math.MinInt64, math.MinInt64

	for _, point := range points {
		if point.x < minx {
			minx = point.x
		}
		if point.y < miny {
			miny = point.y
		}
		if point.x > maxx {
			maxx = point.x
		}
		if point.y > maxy {
			maxy = point.y
		}
	}

	min.x = minx
	min.y = miny
	max.x = maxx
	max.y = maxy

	return min, max
}
