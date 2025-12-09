package containers_test

import (
	"testing"

	"github.com/willmadison/advent/internal/containers"
)

func TestNewStack(t *testing.T) {
	t.Run("creates empty int stack", func(t *testing.T) {
		s := containers.NewStack[int]()
		if s.Size() != 0 {
			t.Errorf("expected size 0, got %d", s.Size())
		}
	})

	t.Run("creates empty string stack", func(t *testing.T) {
		s := containers.NewStack[string]()
		if s.Size() != 0 {
			t.Errorf("expected size 0, got %d", s.Size())
		}
	})

	t.Run("creates empty float64 stack", func(t *testing.T) {
		s := containers.NewStack[float64]()
		if s.Size() != 0 {
			t.Errorf("expected size 0, got %d", s.Size())
		}
	})
}

func TestPush(t *testing.T) {
	t.Run("pushes single element", func(t *testing.T) {
		s := containers.NewStack[int]()
		s.Push(42)

		if s.Size() != 1 {
			t.Errorf("expected size 1, got %d", s.Size())
		}

		val, err := s.Peek()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val != 42 {
			t.Errorf("expected 42, got %d", val)
		}
	})

	t.Run("pushes multiple elements", func(t *testing.T) {
		s := containers.NewStack[int]()
		values := []int{1, 2, 3, 4, 5}

		for _, v := range values {
			s.Push(v)
		}

		if s.Size() != len(values) {
			t.Errorf("expected size %d, got %d", len(values), s.Size())
		}

		val, err := s.Peek()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val != 5 {
			t.Errorf("expected top element 5, got %d", val)
		}
	})

	t.Run("pushes strings", func(t *testing.T) {
		s := containers.NewStack[string]()
		s.Push("hello")
		s.Push("world")

		if s.Size() != 2 {
			t.Errorf("expected size 2, got %d", s.Size())
		}

		val, err := s.Peek()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val != "world" {
			t.Errorf("expected 'world', got %s", val)
		}
	})
}

func TestPop(t *testing.T) {
	t.Run("pops single element", func(t *testing.T) {
		s := containers.NewStack[int]()
		s.Push(42)

		val, err := s.Pop()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val != 42 {
			t.Errorf("expected 42, got %d", val)
		}
		if s.Size() != 0 {
			t.Errorf("expected size 0, got %d", s.Size())
		}
	})

	t.Run("pops in LIFO order", func(t *testing.T) {
		s := containers.NewStack[int]()
		values := []int{1, 2, 3, 4, 5}

		for _, v := range values {
			s.Push(v)
		}

		// Pop in reverse order
		for i := len(values) - 1; i >= 0; i-- {
			val, err := s.Pop()
			if err != nil {
				t.Fatalf("unexpected error at index %d: %v", i, err)
			}
			if val != values[i] {
				t.Errorf("expected %d, got %d", values[i], val)
			}
		}

		if s.Size() != 0 {
			t.Errorf("expected size 0, got %d", s.Size())
		}
	})

	t.Run("returns error when popping from empty stack", func(t *testing.T) {
		s := containers.NewStack[int]()

		val, err := s.Pop()
		if err == nil {
			t.Error("expected error when popping from empty stack, got nil")
		}
		if err.Error() != "no such element" {
			t.Errorf("expected error message 'no such element', got '%s'", err.Error())
		}
		if val != 0 {
			t.Errorf("expected zero value 0, got %d", val)
		}
	})

	t.Run("returns error when popping after stack is emptied", func(t *testing.T) {
		s := containers.NewStack[string]()
		s.Push("test")

		// First pop succeeds
		_, err := s.Pop()
		if err != nil {
			t.Fatalf("unexpected error on first pop: %v", err)
		}

		// Second pop should fail
		val, err := s.Pop()
		if err == nil {
			t.Error("expected error when popping from empty stack, got nil")
		}
		if val != "" {
			t.Errorf("expected zero value empty string, got '%s'", val)
		}
	})
}

func TestPeek(t *testing.T) {
	t.Run("peeks without removing element", func(t *testing.T) {
		s := containers.NewStack[int]()
		s.Push(42)

		val1, err := s.Peek()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val1 != 42 {
			t.Errorf("expected 42, got %d", val1)
		}

		// Peek again - should be same value
		val2, err := s.Peek()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val2 != 42 {
			t.Errorf("expected 42, got %d", val2)
		}

		// Size should still be 1
		if s.Size() != 1 {
			t.Errorf("expected size 1, got %d", s.Size())
		}
	})

	t.Run("returns error when peeking empty stack", func(t *testing.T) {
		s := containers.NewStack[int]()

		val, err := s.Peek()
		if err == nil {
			t.Error("expected error when peeking empty stack, got nil")
		}
		if err.Error() != "no such element" {
			t.Errorf("expected error message 'no such element', got '%s'", err.Error())
		}
		if val != 0 {
			t.Errorf("expected zero value 0, got %d", val)
		}
	})

	t.Run("peeks correct element after multiple pushes", func(t *testing.T) {
		s := containers.NewStack[string]()
		s.Push("first")
		s.Push("second")
		s.Push("third")

		val, err := s.Peek()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val != "third" {
			t.Errorf("expected 'third', got '%s'", val)
		}
	})
}

func TestSize(t *testing.T) {
	t.Run("returns 0 for new stack", func(t *testing.T) {
		s := containers.NewStack[int]()
		if s.Size() != 0 {
			t.Errorf("expected size 0, got %d", s.Size())
		}
	})

	t.Run("increments on push", func(t *testing.T) {
		s := containers.NewStack[int]()

		for i := 1; i <= 10; i++ {
			s.Push(i)
			if s.Size() != i {
				t.Errorf("after %d pushes, expected size %d, got %d", i, i, s.Size())
			}
		}
	})

	t.Run("decrements on pop", func(t *testing.T) {
		s := containers.NewStack[int]()
		for i := 1; i <= 10; i++ {
			s.Push(i)
		}

		for i := 9; i >= 0; i-- {
			s.Pop()
			if s.Size() != i {
				t.Errorf("after pop, expected size %d, got %d", i, s.Size())
			}
		}
	})

	t.Run("size unchanged by peek", func(t *testing.T) {
		s := containers.NewStack[int]()
		s.Push(1)
		s.Push(2)

		initialSize := s.Size()
		s.Peek()
		s.Peek()
		s.Peek()

		if s.Size() != initialSize {
			t.Errorf("expected size to remain %d, got %d", initialSize, s.Size())
		}
	})
}

func TestStackOperations(t *testing.T) {
	t.Run("push and pop sequence", func(t *testing.T) {
		s := containers.NewStack[int]()

		// Push 1, 2, 3
		s.Push(1)
		s.Push(2)
		s.Push(3)

		// Pop 3
		val, _ := s.Pop()
		if val != 3 {
			t.Errorf("expected 3, got %d", val)
		}

		// Push 4
		s.Push(4)

		// Pop 4, 2, 1
		expected := []int{4, 2, 1}
		for i, exp := range expected {
			val, err := s.Pop()
			if err != nil {
				t.Fatalf("unexpected error at position %d: %v", i, err)
			}
			if val != exp {
				t.Errorf("at position %d, expected %d, got %d", i, exp, val)
			}
		}
	})

	t.Run("interleaved push, pop, and peek", func(t *testing.T) {
		s := containers.NewStack[string]()

		s.Push("a")
		val, _ := s.Peek()
		if val != "a" {
			t.Errorf("expected 'a', got '%s'", val)
		}

		s.Push("b")
		val, _ = s.Peek()
		if val != "b" {
			t.Errorf("expected 'b', got '%s'", val)
		}

		val, _ = s.Pop()
		if val != "b" {
			t.Errorf("expected 'b', got '%s'", val)
		}

		val, _ = s.Peek()
		if val != "a" {
			t.Errorf("expected 'a', got '%s'", val)
		}

		s.Push("c")
		val, _ = s.Pop()
		if val != "c" {
			t.Errorf("expected 'c', got '%s'", val)
		}

		val, _ = s.Pop()
		if val != "a" {
			t.Errorf("expected 'a', got '%s'", val)
		}
	})

	t.Run("handles large number of elements", func(t *testing.T) {
		s := containers.NewStack[int]()
		n := 1000

		// Push n elements
		for i := 0; i < n; i++ {
			s.Push(i)
		}

		if s.Size() != n {
			t.Errorf("expected size %d, got %d", n, s.Size())
		}

		// Pop n elements
		for i := n - 1; i >= 0; i-- {
			val, err := s.Pop()
			if err != nil {
				t.Fatalf("unexpected error at %d: %v", i, err)
			}
			if val != i {
				t.Errorf("expected %d, got %d", i, val)
			}
		}

		if s.Size() != 0 {
			t.Errorf("expected size 0, got %d", s.Size())
		}
	})
}

func TestStackWithDifferentTypes(t *testing.T) {
	t.Run("works with float64", func(t *testing.T) {
		s := containers.NewStack[float64]()
		s.Push(3.14)
		s.Push(2.71)

		val, _ := s.Pop()
		if val != 2.71 {
			t.Errorf("expected 2.71, got %f", val)
		}

		val, _ = s.Peek()
		if val != 3.14 {
			t.Errorf("expected 3.14, got %f", val)
		}
	})

	t.Run("works with runes", func(t *testing.T) {
		s := containers.NewStack[rune]()
		s.Push('a')
		s.Push('b')
		s.Push('c')

		val, _ := s.Pop()
		if val != 'c' {
			t.Errorf("expected 'c', got '%c'", val)
		}
	})
}

func TestStackEdgeCases(t *testing.T) {
	t.Run("multiple pops on empty stack", func(t *testing.T) {
		s := containers.NewStack[int]()

		for i := 0; i < 5; i++ {
			_, err := s.Pop()
			if err == nil {
				t.Errorf("iteration %d: expected error, got nil", i)
			}
		}
	})

	t.Run("multiple peeks on empty stack", func(t *testing.T) {
		s := containers.NewStack[int]()

		for i := 0; i < 5; i++ {
			_, err := s.Peek()
			if err == nil {
				t.Errorf("iteration %d: expected error, got nil", i)
			}
		}
	})

	t.Run("push after emptying stack", func(t *testing.T) {
		s := containers.NewStack[int]()

		// Push and pop
		s.Push(1)
		s.Pop()

		// Push again
		s.Push(2)

		val, err := s.Pop()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val != 2 {
			t.Errorf("expected 2, got %d", val)
		}
	})

	t.Run("handles zero values", func(t *testing.T) {
		s := containers.NewStack[int]()
		s.Push(0)
		s.Push(0)
		s.Push(0)

		if s.Size() != 3 {
			t.Errorf("expected size 3, got %d", s.Size())
		}

		val, err := s.Pop()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val != 0 {
			t.Errorf("expected 0, got %d", val)
		}
	})

	t.Run("handles empty strings", func(t *testing.T) {
		s := containers.NewStack[string]()
		s.Push("")
		s.Push("non-empty")
		s.Push("")

		val, _ := s.Pop()
		if val != "" {
			t.Errorf("expected empty string, got '%s'", val)
		}

		val, _ = s.Pop()
		if val != "non-empty" {
			t.Errorf("expected 'non-empty', got '%s'", val)
		}

		val, _ = s.Pop()
		if val != "" {
			t.Errorf("expected empty string, got '%s'", val)
		}
	})
}
