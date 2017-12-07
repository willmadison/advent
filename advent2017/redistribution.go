package advent2017

import (
	"bytes"
	"strconv"
)

type Memory []int

func (m Memory) Serialize() string {
	var buffer bytes.Buffer

	for _, i := range m {
		buffer.WriteString(strconv.Itoa(i))
	}

	return buffer.String()
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
		if currentIndex > len(m) - 1 {
			currentIndex = 0
		}

		m[currentIndex] += 1
		blocks--
		currentIndex++
	}
}

func NumUniqueDistributions(memory Memory) int {
	return 0
}
