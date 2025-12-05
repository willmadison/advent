package intervals

import (
	"sort"
	"strconv"
	"strings"
)

type RangeGroup []Range

func (r RangeGroup) Contains(value int64) bool {
	for _, rng := range r {
		if rng.Contains(value) {
			return true
		}
	}

	return false
}

func (g RangeGroup) normalize() RangeGroup {
	if len(g) == 0 {
		return g
	}

	ranges := make(RangeGroup, len(g))
	copy(ranges, g)

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	merged := ranges[:0]
	for _, r := range ranges {
		if r.End < r.Start {
			continue
		}

		n := len(merged)
		if n == 0 {
			merged = append(merged, r)
			continue
		}

		last := &merged[n-1]
		if r.Start <= last.End {
			if r.End > last.End {
				last.End = r.End
			}
		} else {
			merged = append(merged, r)
		}
	}

	return merged
}

func (r RangeGroup) Size() int {
	normalized := r.normalize()

	var total int
	for _, r := range normalized {
		total += r.Size()
	}
	return total
}

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

func (r Range) Contains(value int64) bool {
	return value >= r.Start && value <= r.End
}

func (r Range) Size() int {
	if r.End < r.Start {
		return 0
	}
	return int(r.End - r.Start + 1)
}

func ParseRange(s string) Range {
	parts := strings.Split(s, "-")
	start, _ := strconv.ParseInt(parts[0], 10, 64)
	end, _ := strconv.ParseInt(parts[1], 10, 64)

	return Range{Start: start, End: end}
}
