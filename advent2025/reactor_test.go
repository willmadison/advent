package advent2025_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2025"
)

func TestCountAllDevicePaths(t *testing.T) {
	given := strings.NewReader(`aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`)

	numExpectedPaths := 5

	paths, err := advent2025.CountAllDevicePaths(given)

	assert.Nil(t, err)

	assert.Equal(t, numExpectedPaths, paths)
}

func TestCountAllDevicePathsPassingThrough(t *testing.T) {
	given := strings.NewReader(`svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`)

	numExpectedPaths := 2

	paths, err := advent2025.CountAllDevicePathsPassingThrough(given, advent2025.Dac, advent2025.Fft)

	assert.Nil(t, err)

	assert.Equal(t, numExpectedPaths, paths)
}
