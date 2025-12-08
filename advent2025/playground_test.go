package advent2025_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2025"
)

func TestFindNLargestCircuits(t *testing.T) {
	given := strings.NewReader(`162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`)

	expectedSizeProduct := 40

	circuits, err := advent2025.FindNLargestCircuits(given, 3, 10)

	assert.Nil(t, err)

	actualProduct := 1

	for _, circuit := range circuits {
		actualProduct *= len(circuit)
	}

	assert.Equal(t, expectedSizeProduct, actualProduct)
}

func TestFindLastConnectionToUnify(t *testing.T) {
	given := strings.NewReader(`162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`)

	expectedProduct := 25272 // 216 * 117

	actualProduct, err := advent2025.FindLastConnectionToUnify(given)

	assert.Nil(t, err)
	assert.Equal(t, expectedProduct, actualProduct)
}
