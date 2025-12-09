package advent2024

import (
	"bufio"
	"container/heap"
	"io"
	"strconv"
	"strings"

	"github.com/willmadison/advent/internal/containers"
)

func DoFindTotalDistance(r io.Reader) int {
	return FindTotalDistance(MatchPairs(r))
}

func MatchPairs(r io.Reader) [][]int {
	left := containers.NewMinHeap[int]()
	right := containers.NewMinHeap[int]()

	heap.Init(left)
	heap.Init(right)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		values := strings.Fields(line)

		i, _ := strconv.Atoi(values[0])
		j, _ := strconv.Atoi(values[1])

		heap.Push(left, i)
		heap.Push(right, j)
	}

	var pairs [][]int

	for left.Len() > 0 && right.Len() > 0 {
		l := heap.Pop(left).(int)
		r := heap.Pop(right).(int)

		pairs = append(pairs, []int{l, r})
	}

	return pairs
}

func ParseLists(r io.Reader) ([]int, []int) {
	var left []int
	var right []int

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		values := strings.Fields(line)

		i, _ := strconv.Atoi(values[0])
		j, _ := strconv.Atoi(values[1])

		left = append(left, i)
		right = append(right, j)
	}

	return left, right
}

func SimilarityScore(r io.Reader) int {
	left, right := ParseLists(r)

	ocurrencesByValue := map[int]int{}

	for _, v := range right {
		ocurrencesByValue[v]++
	}

	var score int

	for _, v := range left {
		score += ocurrencesByValue[v] * v
	}

	return score
}

func FindTotalDistance(pairs [][]int) int {
	var distance int

	for _, pair := range pairs {
		distance += abs(pair[0] - pair[1])
	}

	return distance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
