package advent2019_test

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2019"
)

func TestDeriveFuelRequirementsByMass(t *testing.T) {
	cases := []struct {
		mass     int
		expected int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("DeriveFuelRequirementsByMass(%d)", tc.mass), func(t *testing.T) {
			actual := advent2019.DeriveFuelRequirementsByMass(tc.mass)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestDeriveTotalFuelRequirement(t *testing.T) {
	cases := []struct {
		name     string
		given    io.Reader
		expected int
	}{
		{
			"Basic case...",
			strings.NewReader(`12
14
1969
100756`),
			34241,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := advent2019.DeriveTotalFuelRequirement(tc.given)
			assert.Equal(t, tc.expected, actual)
		})

	}

}

func TestDeriveTotalFuelRequirementIncludingFuelMass(t *testing.T) {
	cases := []struct {
		name     string
		given    io.Reader
		expected int
	}{
		{
			"Basic case...",
			strings.NewReader(`12
14
1969
100756`),
			51316,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := advent2019.DeriveTotalFuelRequirementIncludingFuelMass(tc.given)
			assert.Equal(t, tc.expected, actual)
		})

	}

}
