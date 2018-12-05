package advent2018_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2018"
)

func TestReduction(t *testing.T) {
	cases := []struct {
		given    string
		expected string
	}{
		{
			"dabAcCaCBAcCcaDA",
			"dabCBAcaDA",
		},
		{
			"Aa",
			"",
		},
		{
			"abBA",
			"",
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("Reduce(%s)", tc.given), func(t *testing.T) {
			actual := advent2018.Reduce(tc.given)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestOptimalReduction(t *testing.T) {
	cases := []struct {
		given    string
		expected string
	}{
		{
			"dabAcCaCBAcCcaDA",
			"daDA",
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("OptimalReudction(%s)", tc.given), func(t *testing.T) {
			actual := advent2018.OptimalReduction(tc.given)
			assert.Equal(t, tc.expected, actual)
		})
	}

}
