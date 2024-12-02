package heaps_test

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/internal/heaps"
)

func TestMinHeap(t *testing.T) {
	minHeap := heaps.NewMinHeap[int]()
	heap.Init(minHeap)
	heap.Push(minHeap, 3)
	heap.Push(minHeap, 1)
	heap.Push(minHeap, 2)

	assert.Equal(t, 1, heap.Pop(minHeap))
	assert.Equal(t, 2, heap.Pop(minHeap))
	assert.Equal(t, 3, heap.Pop(minHeap))

}

func TestMaxHeap(t *testing.T) {
	maxHeap := heaps.NewMaxHeap[int]()
	heap.Init(maxHeap)
	heap.Push(maxHeap, 1)
	heap.Push(maxHeap, 3)
	heap.Push(maxHeap, 2)

	assert.Equal(t, 3, heap.Pop(maxHeap))
	assert.Equal(t, 2, heap.Pop(maxHeap))
	assert.Equal(t, 1, heap.Pop(maxHeap))
}
