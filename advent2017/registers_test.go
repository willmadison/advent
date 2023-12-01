package advent2017

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInstruction(t *testing.T) {
	cases := []struct {
		given    string
		expected Instruction
	}{
		{
			`b inc 15`,
			Instruction{
				register:  "b",
				operation: Increment,
				operand:   15,
				condition: Condition{},
			},
		}, {
			`b inc 5 if a > 1`,
			Instruction{
				register:  "b",
				operation: Increment,
				operand:   5,
				condition: Condition{
					register:   "a",
					comparator: GT,
					value:      1,
				},
			},
		},
		{
			`c inc -20 if c == 10`,
			Instruction{
				register:  "c",
				operation: Increment,
				operand:   -20,
				condition: Condition{
					register:   "c",
					comparator: EQ,
					value:      10,
				},
			},
		},
		{
			`a inc 1 if b < 5`,
			Instruction{
				register:  "a",
				operation: Increment,
				operand:   1,
				condition: Condition{
					register:   "b",
					comparator: LT,
					value:      5,
				},
			},
		},
		{
			`c dec -10 if a >= 1`,
			Instruction{
				register:  "c",
				operation: Decrement,
				operand:   -10,
				condition: Condition{
					register:   "a",
					comparator: GTE,
					value:      1,
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.given, func(t *testing.T) {
			actual, err := ParseInstruction(c.given)
			assert.Nil(t, err)
			assert.Equal(t, c.expected, actual)
		})
	}
}

func TestIsZeroCondition(t *testing.T) {
	var condition Condition

	assert.True(t, condition.IsZero())
}
