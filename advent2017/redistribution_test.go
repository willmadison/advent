package advent2017

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumUniqueDistributions(t *testing.T) {
	distributions, cycleSize := NumUniqueDistributions(Memory{0, 2, 7, 0})
	assert.Equal(t, 5, distributions, "incorrect number of unique distributions")
	assert.Equal(t, 4, cycleSize, "incorrect cycle size")
}

func TestSerialization(t *testing.T) {
	assert.Equal(t, "0-2-7-0", Memory{0, 2, 7, 0}.Serialize())
}

func TestRedistribution(t *testing.T) {
	memory := Memory([]int{0, 2, 7, 0})
	memory.Redistribute()
	assert.Equal(t, Memory{2, 4, 1, 2}, memory)
}
