package intervals

import (
	"strconv"
	"strings"
)

type Range struct {
	Start, End int64
}

func (r Range) Iter() <-chan int64 {
	ch := make(chan int64)
	go func() {
		for i := r.Start; i <= r.End; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func ParseRange(s string) Range {
	parts := strings.Split(s, "-")
	start, _ := strconv.ParseInt(parts[0], 10, 64)
	end, _ := strconv.ParseInt(parts[1], 10, 64)

	return Range{Start: start, End: end}
}
