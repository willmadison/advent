package advent2025

import (
	"bufio"
	"io"
	"strings"
)

type Device string

const (
	You    Device = "you"
	Out    Device = "out"
	Server Device = "svr"
	Dac    Device = "dac"
	Fft    Device = "fft"
)

func CountAllDevicePaths(r io.Reader) (int, error) {
	adjacencyList := map[Device]map[Device]struct{}{}

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		from, to := parseConnection(line)

		if _, exists := adjacencyList[from]; !exists {
			adjacencyList[from] = map[Device]struct{}{}
		}

		for _, entry := range to {
			adjacencyList[from][entry] = struct{}{}
		}
	}

	visited := map[Device]struct{}{}
	pathCount := dfs(You, Out, adjacencyList, visited)

	return pathCount, nil
}

func CountAllDevicePathsPassingThrough(r io.Reader, mustPassThrough ...Device) (int, error) {
	adjacencyList := map[Device]map[Device]struct{}{}

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		from, to := parseConnection(line)

		if _, exists := adjacencyList[from]; !exists {
			adjacencyList[from] = map[Device]struct{}{}
		}

		for _, entry := range to {
			adjacencyList[from][entry] = struct{}{}
		}
	}

	requiredBits := make(map[Device]int)
	for i, device := range mustPassThrough {
		requiredBits[device] = i
	}

	targetMask := (1 << len(mustPassThrough)) - 1
	memo := make(map[memoKey]int)
	visited := map[Device]struct{}{}

	pathCount := dfsMemo(Server, Out, 0, targetMask, adjacencyList, requiredBits, visited, memo)

	return pathCount, nil
}

func parseConnection(line string) (Device, []Device) {
	parts := strings.Split(line, " ")

	root := Device(parts[0][0 : len(parts[0])-1])

	var connections []Device

	for _, conn := range parts[1:] {
		connections = append(connections, Device(conn))
	}

	return root, connections
}

type memoKey struct {
	node Device
	mask int
}

func dfsMemo(current, target Device, currentMask, targetMask int, adjacencyList map[Device]map[Device]struct{}, requiredBits map[Device]int, visited map[Device]struct{}, memo map[memoKey]int) int {
	if _, inPath := visited[current]; inPath {
		return 0
	}

	if bitPos, isRequired := requiredBits[current]; isRequired {
		currentMask |= (1 << bitPos)
	}

	if current == target {
		if currentMask == targetMask {
			return 1
		}
		return 0
	}

	key := memoKey{current, currentMask}
	if count, found := memo[key]; found {
		return count
	}

	visited[current] = struct{}{}
	defer delete(visited, current)

	totalPaths := 0
	if neighbors, hasNeighbors := adjacencyList[current]; hasNeighbors {
		for neighbor := range neighbors {
			totalPaths += dfsMemo(neighbor, target, currentMask, targetMask, adjacencyList, requiredBits, visited, memo)
		}
	}

	memo[key] = totalPaths
	return totalPaths
}

func dfs(current, target Device, adjacencyList map[Device]map[Device]struct{}, visited map[Device]struct{}) int {
	if _, inPath := visited[current]; inPath {
		return 0
	}

	if current == target {
		return 1
	}

	visited[current] = struct{}{}
	defer delete(visited, current)

	totalPaths := 0
	if neighbors, hasNeighbors := adjacencyList[current]; hasNeighbors {
		for neighbor := range neighbors {
			totalPaths += dfs(neighbor, target, adjacencyList, visited)
		}
	}

	return totalPaths
}
