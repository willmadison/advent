package heaps

import "cmp"

type Heap[C cmp.Ordered] interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Push(x any)
	Pop() any
}

type MinHeap[C cmp.Ordered] struct {
	data []C
}

func NewMinHeap[C cmp.Ordered]() Heap[C] {
	return &MinHeap[C]{data: []C{}}
}

func (m *MinHeap[C]) Len() int {
	return len(m.data)
}

func (m *MinHeap[C]) Less(i, j int) bool {
	return m.data[i] < m.data[j]
}

func (m *MinHeap[C]) Swap(i, j int) {
	m.data[i], m.data[j] = m.data[j], m.data[i]
}

func (m *MinHeap[C]) Push(x any) {
	(*m).data = append((*m).data, x.(C))
}

func (m *MinHeap[C]) Pop() any {
	old := (*m).data
	n := len(old)
	x := old[n-1]
	(*m).data = old[0 : n-1]

	return x
}

type MaxHeap[C cmp.Ordered] struct {
	data []C
}

func NewMaxHeap[C cmp.Ordered]() Heap[C] {
	return &MaxHeap[C]{data: []C{}}
}

func (m *MaxHeap[C]) Len() int {
	return len(m.data)
}

func (m *MaxHeap[C]) Less(i, j int) bool {
	return m.data[i] > m.data[j]
}

func (m *MaxHeap[C]) Swap(i, j int) {
	m.data[i], m.data[j] = m.data[j], m.data[i]
}

func (m *MaxHeap[C]) Push(x any) {
	(*m).data = append((*m).data, x.(C))
}

func (m *MaxHeap[C]) Pop() any {
	old := (*m).data
	n := len(old)
	x := old[n-1]
	(*m).data = old[0 : n-1]

	return x
}
