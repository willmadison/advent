package advent2024_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2024"
)

func TestFindInstructions(t *testing.T) {
	given := strings.NewReader(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`)

	instructions, err := advent2024.FindInstructions(given)

	assert.Nil(t, err)
	assert.Equal(t, 4, len(instructions))
	assert.Equal(t, 2, len(instructions[0].Operands))

	expectedProduct := 161

	var actualProduct int

	for _, i := range instructions {
		actualProduct += i.Operands[0] * i.Operands[1]
	}

	assert.Equal(t, expectedProduct, actualProduct)

	given = strings.NewReader(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`)

	instructions, err = advent2024.FindInstructions(given)

	assert.Nil(t, err)
	assert.Equal(t, 6, len(instructions))
	assert.Equal(t, 2, len(instructions[0].Operands))

	expectedProduct = 48

	actualProduct = 0

	doableInstructions := advent2024.FindDoables(instructions)

	for _, i := range doableInstructions {
		actualProduct += i.Operands[0] * i.Operands[1]
	}

	assert.Equal(t, expectedProduct, actualProduct)
}
