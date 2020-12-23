package advent2020

import (
	"container/ring"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCups(t *testing.T) {
	given := strings.NewReader(`32415`)

	cupsByID := map[int]*ring.Ring{}
	cupRing := ring.New(5)

	cupRing.Value = 3
	cupsByID[3] = cupRing
	cupRing = cupRing.Next()
	cupRing.Value = 2
	cupsByID[2] = cupRing
	cupRing = cupRing.Next()
	cupRing.Value = 4
	cupsByID[4] = cupRing
	cupRing = cupRing.Next()
	cupRing.Value = 1
	cupsByID[1] = cupRing
	cupRing = cupRing.Next()
	cupRing.Value = 5
	cupsByID[5] = cupRing
	cupRing = cupRing.Next()

	expected := Circle{
		Cups:     cupRing,
		CupsByID: cupsByID,
		Current:  cupsByID[3],
		Low:      1,
		High:     5,
	}

	actual := ParseCups(given)

	assert.Equal(t, expected, actual)
}

func TestMove(t *testing.T) {
	circle := ParseCups(strings.NewReader("389125467"))

	for i := 0; i < 100; i++ {
		circle.Move()
	}

	expected := "67384529"

	assert.Equal(t, expected, circle.Serialize(1)[1:])
}

func TestAdditionalMoves(t *testing.T) {
	circle := ParseCups(strings.NewReader("389125467"))

	circle.AddAdditionalCups(1000000 - circle.Cups.Len())

	for i := 0; i < 10_000_000; i++ {
		circle.Move()
	}

	first := circle.CupsByID[1]

	a := first.Move(1).Value.(int)
	b := first.Move(2).Value.(int)

	expectedNext := 934001
	expectedTwoAway := 159792

	assert.Equal(t, expectedNext, a)
	assert.Equal(t, expectedTwoAway, b)
}
