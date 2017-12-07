package advent2017

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseJumpInstructions(t *testing.T) {
	assert.Equal(t, []int{0, 3, 0, 1, -3}, parseJumpInstructions(`0
3
0
1
-3`))
}

func TestJumpIterations(t *testing.T) {
	assert.Equal(t, 10, JumpIterations([]int{0, 3, 0, 1, -3}))
}

