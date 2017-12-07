package advent2017

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNumUniqueDistributions(t *testing.T) {
	assert.Equal(t, 5, NumUniqueDistributions(Memory{0, 2, 7, 0}))
}

func TestSerialization(t *testing.T) {
	assert.Equal(t, "0270", Memory{0,2,7,0}.Serialize())
}

func TestRedistribution(t *testing.T) {
	memory := Memory([]int{0,2,7,0})
	memory.Redistribute()
	assert.Equal(t, Memory{2, 4, 1, 2}, memory)
}
