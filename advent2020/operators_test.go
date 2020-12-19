package advent2020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseExpression(t *testing.T) {
	cases := []struct {
		scenario string
		given    string
		expected Expression
	}{
		{
			"no grouping",
			"1 + 2 * 3 + 4 * 5 + 6",
			Expression([]Token{
				Operand(1),
				Operand(2),
				Operator("+"),
				Operand(3),
				Operator("*"),
				Operand(4),
				Operator("+"),
				Operand(5),
				Operator("*"),
				Operand(6),
				Operator("+"),
			}),
		},
		{
			"grouping",
			"1 + (2 * 3) + (4 * (5 + 6))",
			Expression([]Token{
				Operand(1),
				Operand(2),
				Operand(3),
				Operator("*"),
				Operator("+"),
				Operand(4),
				Operand(5),
				Operand(6),
				Operator("+"),
				Operator("*"),
				Operator("+"),
			}),
		},
		{
			"complex grouping",
			"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			Expression([]Token{
				Operand(2),
				Operand(4),
				Operator("+"),
				Operand(9),
				Operator("*"),
				Operand(6),
				Operand(9),
				Operator("+"),
				Operand(8),
				Operator("*"),
				Operand(6),
				Operator("+"),
				Operator("*"),
				Operand(6),
				Operator("+"),
				Operand(2),
				Operator("+"),
				Operand(4),
				Operator("+"),
				Operand(2),
				Operator("*"),
			}),
		},
		{
			"much more complex grouping",
			"6 * ((5 * 3 * 2 + 9 * 4) * (8 * 8 + 2 * 3) * 5 * 8) * 2 + (4 + 9 * 5 * 5 + 8) * 4",
			Expression([]Token{
				Operand(6),
				Operand(5),
				Operand(3),
				Operator("*"),
				Operand(2),
				Operator("*"),
				Operand(9),
				Operator("+"),
				Operand(4),
				Operator("*"),
				Operand(8),
				Operand(8),
				Operator("*"),
				Operand(2),
				Operator("+"),
				Operand(3),
				Operator("*"),
				Operator("*"),
				Operand(5),
				Operator("*"),
				Operand(8),
				Operator("*"),
				Operator("*"),
				Operand(2),
				Operator("*"),
				Operand(4),
				Operand(9),
				Operator("+"),
				Operand(5),
				Operator("*"),
				Operand(5),
				Operator("*"),
				Operand(8),
				Operator("+"),
				Operator("+"),
				Operand(4),
				Operator("*"),
			}),
		},
	}

	for _, tc := range cases {
		t.Run(tc.scenario, func(t *testing.T) {
			actual := ParseExpression(tc.given)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestEvaluateExpression(t *testing.T) {
	cases := []struct {
		given    Expression
		expected int
	}{
		{
			[]Token{
				Operand(1),
				Operand(2),
				Operator("+"),
				Operand(3),
				Operator("*"),
				Operand(4),
				Operator("+"),
				Operand(5),
				Operator("*"),
				Operand(6),
				Operator("+"),
			},
			71,
		},
		{
			[]Token{
				Operand(1),
				Operand(2),
				Operand(3),
				Operator("*"),
				Operator("+"),
				Operand(4),
				Operand(5),
				Operand(6),
				Operator("+"),
				Operator("*"),
				Operator("+"),
			},
			51,
		},
		{
			[]Token{
				Operand(2),
				Operand(4),
				Operator("+"),
				Operand(9),
				Operator("*"),
				Operand(6),
				Operand(9),
				Operator("+"),
				Operand(8),
				Operator("*"),
				Operand(6),
				Operator("+"),
				Operator("*"),
				Operand(6),
				Operator("+"),
				Operand(2),
				Operator("+"),
				Operand(4),
				Operator("+"),
				Operand(2),
				Operator("*"),
			},
			13632,
		},
		{
			[]Token{
				Operand(6),
				Operand(5),
				Operand(3),
				Operator("*"),
				Operand(2),
				Operator("*"),
				Operand(9),
				Operator("+"),
				Operand(4),
				Operator("*"),
				Operand(8),
				Operand(8),
				Operator("*"),
				Operand(2),
				Operator("+"),
				Operand(3),
				Operator("*"),
				Operator("*"),
				Operand(5),
				Operator("*"),
				Operand(8),
				Operator("*"),
				Operator("*"),
				Operand(2),
				Operator("*"),
				Operand(4),
				Operand(9),
				Operator("+"),
				Operand(5),
				Operator("*"),
				Operand(5),
				Operator("*"),
				Operand(8),
				Operator("+"),
				Operator("+"),
				Operand(4),
				Operator("*"),
			},
			59306292,
		},
	}

	for _, tc := range cases {
		actual := EvaluateExpression(tc.given)
		assert.Equal(t, tc.expected, actual)
	}
}
