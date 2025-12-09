package advent2025

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/willmadison/advent/internal/containers/graphs"
	"github.com/willmadison/advent/internal/location"
)

type Circuit []location.Point

func getCircuits(uf *graphs.UnionFind, points []location.Point) []Circuit {
	groups := uf.GetGroups()

	circuits := make([]Circuit, 0, len(groups))
	for _, indices := range groups {
		circuit := make(Circuit, len(indices))
		for i, idx := range indices {
			circuit[i] = points[idx]
		}
		circuits = append(circuits, circuit)
	}

	return circuits
}

type edge struct {
	i, j     int
	distance int
}

func FindNLargestCircuits(r io.Reader, n int, numConnections int) ([]Circuit, error) {
	scanner := bufio.NewScanner(r)

	var points []location.Point

	for scanner.Scan() {
		line := scanner.Text()

		point, err := parsePoint(line)
		if err != nil {
			return nil, err
		}

		points = append(points, point)
	}

	if len(points) == 0 {
		return []Circuit{}, nil
	}

	var edges []edge
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := points[i].EuclideanDistance(points[j])
			edges = append(edges, edge{i, j, dist})
		}
	}

	slices.SortFunc(edges, func(a, b edge) int {
		return a.distance - b.distance
	})

	uf := graphs.NewUnionFind(len(points))

	for i := 0; i < numConnections && i < len(edges); i++ {
		e := edges[i]
		uf.Union(e.i, e.j)
	}

	circuits := getCircuits(uf, points)

	slices.SortFunc(circuits, func(a, b Circuit) int {
		return len(b) - len(a)
	})

	if n > len(circuits) {
		n = len(circuits)
	}

	return circuits[:n], nil
}

func FindLastConnectionToUnify(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)

	var points []location.Point

	for scanner.Scan() {
		line := scanner.Text()

		point, err := parsePoint(line)
		if err != nil {
			return 0, err
		}

		points = append(points, point)
	}

	if len(points) == 0 {
		return 0, nil
	}

	var edges []edge
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := points[i].EuclideanDistance(points[j])
			edges = append(edges, edge{i, j, dist})
		}
	}

	slices.SortFunc(edges, func(a, b edge) int {
		return a.distance - b.distance
	})

	uf := graphs.NewUnionFind(len(points))
	numComponents := len(points)

	for _, e := range edges {
		if uf.Union(e.i, e.j) {
			numComponents--
			if numComponents == 1 {
				return points[e.i].X * points[e.j].X, nil
			}
		}
	}

	return 0, nil
}

func parsePoint(s string) (location.Point, error) {
	var X, Y, Z int

	values := strings.Split(s, ",")

	for i, v := range values {
		n, err := strconv.Atoi(v)
		if err != nil {
			return location.Point{}, err
		}

		switch i {
		case 0:
			X = n
		case 1:
			Y = n
		case 2:
			Z = n
		}
	}

	return location.Point{X: X, Y: Y, Z: Z}, nil
}
