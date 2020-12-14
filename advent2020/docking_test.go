package advent2020

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInitializationProgram(t *testing.T) {
	given := strings.NewReader(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`)

	expected := Program([]Instruction{
		{Operation: SetMask, Mask: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"},
		{Operation: Write, Location: 8, Value: 11},
		{Operation: Write, Location: 7, Value: 101},
		{Operation: Write, Location: 8, Value: 0},
	})

	actual := ParseInitializationProgram(given)

	assert.Equal(t, expected, actual)
}

func TestRunProgramV1(t *testing.T) {
	given := Program([]Instruction{
		{Operation: SetMask, Mask: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"},
		{Operation: Write, Location: 8, Value: 11},
		{Operation: Write, Location: 7, Value: 101},
		{Operation: Write, Location: 8, Value: 0},
	})

	expectedTotalMemoryValue := 165
	var c Computer

	given.Run(&c)

	var actualTotalMemoryValue int

	for _, v := range c.Memory {
		actualTotalMemoryValue += v
	}

	assert.Equal(t, expectedTotalMemoryValue, actualTotalMemoryValue)
}

func TestRunProgramV2(t *testing.T) {
	given := Program([]Instruction{
		{Operation: SetMask, Mask: "000000000000000000000000000000X1001X"},
		{Operation: Write, Location: 42, Value: 100},
		{Operation: SetMask, Mask: "00000000000000000000000000000000X0XX"},
		{Operation: Write, Location: 26, Value: 1},
	})

	expectedTotalMemoryValue := 208

	var c Computer
	c.Version = Version2

	given.Run(&c)

	var actualTotalMemoryValue int

	for _, v := range c.Memory {
		actualTotalMemoryValue += v
	}

	assert.Equal(t, expectedTotalMemoryValue, actualTotalMemoryValue)
}
