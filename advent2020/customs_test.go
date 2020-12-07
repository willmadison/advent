package advent2020

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGroups(t *testing.T) {
	given := strings.NewReader(`abc

a
b
c

ab
ac

a
a
a
a

b`)

	expected := []Group{
		{
			map[int]map[rune]struct{}{
				0: {
					'a': {},
					'b': {},
					'c': {},
				},
			},
		},
		{
			map[int]map[rune]struct{}{
				0: {
					'a': {},
				},
				1: {
					'b': {},
				},
				2: {
					'c': {},
				},
			},
		},
		{
			map[int]map[rune]struct{}{
				0: {
					'a': {},
					'b': {},
				},
				1: {
					'a': {},
					'c': {},
				},
			},
		},
		{
			map[int]map[rune]struct{}{
				0: {
					'a': {},
				},
				1: {
					'a': {},
				},
				2: {
					'a': {},
				},
				3: {
					'a': {},
				},
			},
		},
		{
			map[int]map[rune]struct{}{
				0: {
					'b': {},
				},
			},
		},
	}

	actual := ParseGroups(given)
	assert.Equal(t, expected, actual)
}

func TestCountAffirmatives(t *testing.T) {
	given := []Group{
		{
			map[int]map[rune]struct{}{
				0: {
					'a': {},
					'b': {},
					'c': {},
				},
			},
		},
		{
			map[int]map[rune]struct{}{
				0: {
					'a': {},
				},
				1: {
					'b': {},
				},
				2: {
					'c': {},
				},
			},
		},
		{
			map[int]map[rune]struct{}{
				0: {
					'a': {},
					'b': {},
				},
				1: {
					'a': {},
					'c': {},
				},
			},
		},
		{
			map[int]map[rune]struct{}{
				0: {
					'a': {},
				},
				1: {
					'a': {},
				},
				2: {
					'a': {},
				},
				3: {
					'a': {},
				},
			},
		},
		{
			map[int]map[rune]struct{}{
				0: {
					'b': {},
				},
			},
		},
	}

	expected := 6

	actual := CountAffirmatives(given)
	assert.Equal(t, expected, actual)
}
