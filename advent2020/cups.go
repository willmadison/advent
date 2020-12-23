package advent2020

import (
	"bytes"
	"container/ring"
	"io"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Circle struct {
	Cups      *ring.Ring
	CupsByID  map[int]*ring.Ring
	Current   *ring.Ring
	Low, High int
}

func (c *Circle) Move() {
	removed := c.Current.Unlink(3)

	inHand := map[int]struct{}{}

	for n := 0; n < 3; n++ {
		inHand[removed.Move(n).Value.(int)] = struct{}{}
	}

	destination := c.Current.Value.(int) - 1

	var found bool

	for !found {
		if destination < c.Low {
			destination = c.High
		}

		if _, holding := inHand[destination]; !holding {
			found = true
			break
		}

		destination--
	}

	c.CupsByID[destination].Link(removed)
	c.Current = c.Current.Next()
}

func (c Circle) Serialize(start ...int) string {
	startingPoint := c.Current

	if len(start) > 0 {
		startingPoint = c.CupsByID[start[0]]
	}

	var buf bytes.Buffer

	current := startingPoint

	for i := 0; i < c.Cups.Len(); i++ {
		buf.WriteString(strconv.Itoa(current.Value.(int)))
		current = current.Next()
	}

	return buf.String()
}

func (c *Circle) AddAdditionalCups(n int) {
	additionalCups := ring.New(n)

	c.Current.Prev().Link(additionalCups)

	for nextCupID := c.High + 1; nextCupID <= n+c.High; nextCupID++ {
		additionalCups.Value = nextCupID
		c.CupsByID[nextCupID] = additionalCups
		additionalCups = additionalCups.Next()

	}

	c.High += n
}

func ParseCups(r io.Reader) Circle {
	rawValue, _ := ioutil.ReadAll(r)
	value := strings.Trim(string(rawValue), "\n")

	circle := Circle{}

	circle.Low = math.MaxInt64
	circle.High = math.MinInt64

	var first *ring.Ring

	cups := ring.New(len(value))
	cupsByID := map[int]*ring.Ring{}

	for _, v := range value {
		cupID, _ := strconv.Atoi(string(v))

		cups.Value = cupID
		cupsByID[cupID] = cups

		if first == nil {
			first = cups
		}

		cups = cups.Next()

		if cupID > circle.High {
			circle.High = cupID
		}

		if cupID < circle.Low {
			circle.Low = cupID
		}
	}

	circle.Cups = cups
	circle.CupsByID = cupsByID
	circle.Current = first

	return circle
}
