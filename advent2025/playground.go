package advent2025

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/willmadison/advent/internal/location"
)

type Circuit []location.Point

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{parent, size}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX != rootY {
		if uf.size[rootX] < uf.size[rootY] {
			rootX, rootY = rootY, rootX
		}
		uf.parent[rootY] = rootX
		uf.size[rootX] += uf.size[rootY]
		return true
	}
	return false
}

func (uf *UnionFind) GetCircuits(points []location.Point) []Circuit {
	groups := make(map[int][]int)
	for i := range uf.parent {
		root := uf.Find(i)
		groups[root] = append(groups[root], i)
	}

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

	uf := NewUnionFind(len(points))

	for i := 0; i < numConnections && i < len(edges); i++ {
		e := edges[i]
		uf.Union(e.i, e.j)
	}

	circuits := uf.GetCircuits(points)

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

	uf := NewUnionFind(len(points))
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
