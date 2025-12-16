package advent2022_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2022"
)

func TestFindTerminalTopCrates(t *testing.T) {
	given := strings.NewReader(`    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`)

	expected := "CMZ"

	result, err := advent2022.FindTerminalTopCrates(given, advent2022.MoveStrategyOneByOne)

	assert.Nil(t, err)

	assert.Equal(t, expected, result)

	given = strings.NewReader(`    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`)

	expected = "MCD"

	result, err = advent2022.FindTerminalTopCrates(given, advent2022.MoveStrategyMultipleAtOnce)

	assert.Nil(t, err)

	assert.Equal(t, expected, result)
}
