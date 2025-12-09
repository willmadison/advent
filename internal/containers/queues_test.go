package containers

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()

	if q == nil {
		t.Fatal("NewQueue() returned nil")
	}

	if q.Size() != 0 {
		t.Errorf("NewQueue() size = %d, want 0", q.Size())
	}
}

func TestEnqueue(t *testing.T) {
	q := NewQueue[int]()

	q.Enqueue(1)
	if q.Size() != 1 {
		t.Errorf("After Enqueue(1), size = %d, want 1", q.Size())
	}

	q.Enqueue(2)
	q.Enqueue(3)
	if q.Size() != 3 {
		t.Errorf("After enqueueing 3 items, size = %d, want 3", q.Size())
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue[int]()

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	val, err := q.Dequeue()
	if err != nil {
		t.Fatalf("Dequeue() error = %v, want nil", err)
	}
	if val != 10 {
		t.Errorf("Dequeue() = %d, want 10", val)
	}
	if q.Size() != 2 {
		t.Errorf("After dequeue, size = %d, want 2", q.Size())
	}

	val, err = q.Dequeue()
	if err != nil {
		t.Fatalf("Dequeue() error = %v, want nil", err)
	}
	if val != 20 {
		t.Errorf("Dequeue() = %d, want 20", val)
	}

	val, err = q.Dequeue()
	if err != nil {
		t.Fatalf("Dequeue() error = %v, want nil", err)
	}
	if val != 30 {
		t.Errorf("Dequeue() = %d, want 30", val)
	}
	if q.Size() != 0 {
		t.Errorf("After dequeuing all items, size = %d, want 0", q.Size())
	}
}

func TestDequeueEmpty(t *testing.T) {
	q := NewQueue[int]()

	val, err := q.Dequeue()
	if err == nil {
		t.Error("Dequeue() on empty queue should return error")
	}
	if val != 0 {
		t.Errorf("Dequeue() on empty queue = %d, want 0 (zero value)", val)
	}
	if err.Error() != "no such element" {
		t.Errorf("Dequeue() error message = %q, want %q", err.Error(), "no such element")
	}
}

func TestPeek(t *testing.T) {
	q := NewQueue[string]()

	q.Enqueue("first")
	q.Enqueue("second")
	q.Enqueue("third")

	val, err := q.Peek()
	if err != nil {
		t.Fatalf("Peek() error = %v, want nil", err)
	}
	if val != "first" {
		t.Errorf("Peek() = %q, want %q", val, "first")
	}
	if q.Size() != 3 {
		t.Errorf("After Peek(), size = %d, want 3 (size should not change)", q.Size())
	}

	// Peek again to ensure it doesn't remove the element
	val, err = q.Peek()
	if err != nil {
		t.Fatalf("Second Peek() error = %v, want nil", err)
	}
	if val != "first" {
		t.Errorf("Second Peek() = %q, want %q", val, "first")
	}
}

func TestPeekEmpty(t *testing.T) {
	q := NewQueue[string]()

	val, err := q.Peek()
	if err == nil {
		t.Error("Peek() on empty queue should return error")
	}
	if val != "" {
		t.Errorf("Peek() on empty queue = %q, want empty string (zero value)", val)
	}
	if err.Error() != "no such element" {
		t.Errorf("Peek() error message = %q, want %q", err.Error(), "no such element")
	}
}

func TestSize(t *testing.T) {
	q := NewQueue[int]()

	sizes := []int{0, 1, 2, 3, 4, 5}

	for i, expectedSize := range sizes {
		if q.Size() != expectedSize {
			t.Errorf("At step %d, size = %d, want %d", i, q.Size(), expectedSize)
		}
		if i < len(sizes)-1 {
			q.Enqueue(i)
		}
	}

	// Dequeue all and verify size decreases
	for i := 5; i > 0; i-- {
		q.Dequeue()
		expectedSize := i - 1
		if q.Size() != expectedSize {
			t.Errorf("After dequeue, size = %d, want %d", q.Size(), expectedSize)
		}
	}
}

func TestFIFOOrder(t *testing.T) {
	q := NewQueue[int]()

	// Enqueue 1 through 100
	for i := 1; i <= 100; i++ {
		q.Enqueue(i)
	}

	// Dequeue and verify order
	for i := 1; i <= 100; i++ {
		val, err := q.Dequeue()
		if err != nil {
			t.Fatalf("Dequeue() at iteration %d error = %v, want nil", i, err)
		}
		if val != i {
			t.Errorf("Dequeue() at iteration %d = %d, want %d", i, val, i)
		}
	}
}

func TestQueueString(t *testing.T) {
	q := NewQueue[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	str := q.(*queue[int]).String()
	expected := "[1 2 3]"
	if str != expected {
		t.Errorf("String() = %q, want %q", str, expected)
	}
}

func TestQueueWithStructs(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	q := NewQueue[Person]()

	alice := Person{"Alice", 30}
	bob := Person{"Bob", 25}
	charlie := Person{"Charlie", 35}

	q.Enqueue(alice)
	q.Enqueue(bob)
	q.Enqueue(charlie)

	if q.Size() != 3 {
		t.Errorf("Size = %d, want 3", q.Size())
	}

	val, err := q.Dequeue()
	if err != nil {
		t.Fatalf("Dequeue() error = %v, want nil", err)
	}
	if val != alice {
		t.Errorf("Dequeue() = %v, want %v", val, alice)
	}

	val, err = q.Peek()
	if err != nil {
		t.Fatalf("Peek() error = %v, want nil", err)
	}
	if val != bob {
		t.Errorf("Peek() = %v, want %v", val, bob)
	}
}

func TestMixedOperations(t *testing.T) {
	q := NewQueue[int]()

	// Enqueue some items
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Dequeue one
	val, _ := q.Dequeue()
	if val != 1 {
		t.Errorf("First dequeue = %d, want 1", val)
	}

	// Enqueue more
	q.Enqueue(4)
	q.Enqueue(5)

	// Peek
	val, _ = q.Peek()
	if val != 2 {
		t.Errorf("Peek = %d, want 2", val)
	}

	// Dequeue remaining in order
	expected := []int{2, 3, 4, 5}
	for i, exp := range expected {
		val, err := q.Dequeue()
		if err != nil {
			t.Fatalf("Dequeue at position %d error = %v, want nil", i, err)
		}
		if val != exp {
			t.Errorf("Dequeue at position %d = %d, want %d", i, val, exp)
		}
	}

	// Verify empty
	if q.Size() != 0 {
		t.Errorf("Final size = %d, want 0", q.Size())
	}

	_, err := q.Dequeue()
	if err == nil {
		t.Error("Dequeue on empty queue should return error")
	}
}

func TestQueueWithPointers(t *testing.T) {
	q := NewQueue[*int]()

	a := 10
	b := 20
	c := 30

	q.Enqueue(&a)
	q.Enqueue(&b)
	q.Enqueue(&c)

	val, err := q.Dequeue()
	if err != nil {
		t.Fatalf("Dequeue() error = %v, want nil", err)
	}
	if val != &a || *val != 10 {
		t.Errorf("Dequeue() = %v (value %d), want pointer to a (value 10)", val, *val)
	}
}

func TestEmptyQueueState(t *testing.T) {
	q := NewQueue[int]()

	// Enqueue and dequeue to empty the queue
	q.Enqueue(1)
	q.Dequeue()

	// Verify it's properly empty
	if q.Size() != 0 {
		t.Errorf("Size after dequeuing all = %d, want 0", q.Size())
	}

	_, err := q.Peek()
	if err == nil {
		t.Error("Peek() on emptied queue should return error")
	}

	_, err = q.Dequeue()
	if err == nil {
		t.Error("Dequeue() on emptied queue should return error")
	}

	// Should be able to use again
	q.Enqueue(2)
	val, err := q.Dequeue()
	if err != nil {
		t.Fatalf("Dequeue() after re-use error = %v, want nil", err)
	}
	if val != 2 {
		t.Errorf("Dequeue() after re-use = %d, want 2", val)
	}
}
