package advent2017

import (
	"strconv"
	"strings"
)

type Memory []int

func (m Memory) Serialize() string {
	var values []string

	for _, i := range m {
		values = append(values, strconv.Itoa(i))
	}

	return strings.Join(values, "-")
}

func (m Memory) Redistribute() {
	maxBlocks := m[0]
	maxAt := 0

	for i, b := range m {
		if b > maxBlocks {
			maxBlocks = b
			maxAt = i
		}
	}

	blocks := maxBlocks
	m[maxAt] = 0

	currentIndex := maxAt + 1

	for blocks > 0 {
		if currentIndex > len(m)-1 {
			currentIndex = 0
		}

		m[currentIndex]++
		blocks--
		currentIndex++
	}
}

func NumUniqueDistributions(memory Memory) (int, int) {
	distributions := map[string]int{}

	var redistributions int

	var seen bool

	for !seen {
		memory.Redistribute()
		redistributions++

		key := memory.Serialize()
		_, seen = distributions[key]

		if !seen {
			distributions[key] = redistributions
		}
	}

	return redistributions, redistributions - distributions[memory.Serialize()]
}
