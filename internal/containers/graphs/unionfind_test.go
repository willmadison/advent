package graphs

import (
	"testing"
)

func TestNewUnionFind(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{"single element", 1},
		{"small set", 5},
		{"larger set", 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := NewUnionFind(tt.n)
			if uf == nil {
				t.Fatal("NewUnionFind returned nil")
			}
			if len(uf.parent) != tt.n {
				t.Errorf("expected parent length %d, got %d", tt.n, len(uf.parent))
			}
			if len(uf.size) != tt.n {
				t.Errorf("expected size length %d, got %d", tt.n, len(uf.size))
			}

			for i := 0; i < tt.n; i++ {
				if uf.parent[i] != i {
					t.Errorf("expected parent[%d] = %d, got %d", i, i, uf.parent[i])
				}
				if uf.size[i] != 1 {
					t.Errorf("expected size[%d] = 1, got %d", i, uf.size[i])
				}
			}
		})
	}
}

func TestFind(t *testing.T) {
	uf := NewUnionFind(5)

	for i := 0; i < 5; i++ {
		if root := uf.Find(i); root != i {
			t.Errorf("Find(%d) = %d, expected %d", i, root, i)
		}
	}

	uf.Union(0, 1)
	root0 := uf.Find(0)
	root1 := uf.Find(1)
	if root0 != root1 {
		t.Errorf("after Union(0,1), Find(0)=%d and Find(1)=%d should be equal", root0, root1)
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		unions   [][2]int
		expected map[int][]int
	}{
		{
			name:   "simple union",
			n:      5,
			unions: [][2]int{{0, 1}, {2, 3}},
			expected: map[int][]int{
				0: {0, 1},
				2: {2, 3},
				4: {4},
			},
		},
		{
			name:   "chain unions",
			n:      5,
			unions: [][2]int{{0, 1}, {1, 2}, {2, 3}},
			expected: map[int][]int{
				0: {0, 1, 2, 3},
				4: {4},
			},
		},
		{
			name:   "all connected",
			n:      4,
			unions: [][2]int{{0, 1}, {1, 2}, {2, 3}},
			expected: map[int][]int{
				0: {0, 1, 2, 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := NewUnionFind(tt.n)

			for _, u := range tt.unions {
				uf.Union(u[0], u[1])
			}

			groups := uf.GetGroups()
			if len(groups) != len(tt.expected) {
				t.Errorf("expected %d groups, got %d", len(tt.expected), len(groups))
			}

			groupsByRoot := make(map[int][]int)
			for i := 0; i < tt.n; i++ {
				root := uf.Find(i)
				groupsByRoot[root] = append(groupsByRoot[root], i)
			}

			for root, members := range groupsByRoot {
				if len(members) != len(tt.expected[root]) {
					t.Errorf("group rooted at %d has %d members, expected %d", root, len(members), len(tt.expected[root]))
				}
			}
		})
	}
}

func TestUnionReturnValue(t *testing.T) {
	uf := NewUnionFind(5)

	if !uf.Union(0, 1) {
		t.Error("Union(0, 1) should return true when connecting different components")
	}

	if uf.Union(0, 1) {
		t.Error("Union(0, 1) should return false when elements are already connected")
	}

	if !uf.Union(2, 3) {
		t.Error("Union(2, 3) should return true when connecting different components")
	}

	if uf.Union(3, 2) {
		t.Error("Union(3, 2) should return false when elements are already connected")
	}
}

func TestPathCompression(t *testing.T) {
	uf := NewUnionFind(10)

	uf.Union(0, 1)
	uf.Union(1, 2)
	uf.Union(2, 3)
	uf.Union(3, 4)

	root := uf.Find(4)

	uf.Find(4)
	if uf.parent[4] != root {
		t.Error("path compression should update parent of 4 to point directly to root")
	}
}

func TestUnionBySize(t *testing.T) {
	uf := NewUnionFind(6)

	uf.Union(0, 1)
	uf.Union(2, 3)
	uf.Union(3, 4)

	root01 := uf.Find(0)
	root234 := uf.Find(2)

	size01 := uf.size[root01]
	size234 := uf.size[root234]

	if size01 != 2 {
		t.Errorf("size of component {0,1} should be 2, got %d", size01)
	}
	if size234 != 3 {
		t.Errorf("size of component {2,3,4} should be 3, got %d", size234)
	}

	uf.Union(0, 2)

	finalRoot := uf.Find(0)
	finalSize := uf.size[finalRoot]

	if finalSize != 5 {
		t.Errorf("after merging, size should be 5, got %d", finalSize)
	}

	if uf.Find(1) != finalRoot {
		t.Error("all elements should have same root after union")
	}
	if uf.Find(3) != finalRoot {
		t.Error("all elements should have same root after union")
	}
}

func TestGetGroups(t *testing.T) {
	uf := NewUnionFind(8)

	uf.Union(0, 1)
	uf.Union(1, 2)
	uf.Union(3, 4)
	uf.Union(6, 7)

	groups := uf.GetGroups()

	expectedGroupCount := 4
	if len(groups) != expectedGroupCount {
		t.Errorf("expected %d groups, got %d", expectedGroupCount, len(groups))
	}

	totalElements := 0
	for _, group := range groups {
		totalElements += len(group)
	}

	if totalElements != 8 {
		t.Errorf("expected 8 total elements across all groups, got %d", totalElements)
	}

	for root, members := range groups {
		for _, member := range members {
			if uf.Find(member) != root {
				t.Errorf("member %d in group with root %d has different root %d", member, root, uf.Find(member))
			}
		}
	}
}

func TestGetGroupsSingleElement(t *testing.T) {
	uf := NewUnionFind(1)
	groups := uf.GetGroups()

	if len(groups) != 1 {
		t.Errorf("expected 1 group for single element, got %d", len(groups))
	}

	for _, group := range groups {
		if len(group) != 1 {
			t.Errorf("expected group size 1, got %d", len(group))
		}
		if group[0] != 0 {
			t.Errorf("expected group to contain element 0, got %d", group[0])
		}
	}
}

func TestGetGroupsNoUnions(t *testing.T) {
	uf := NewUnionFind(5)
	groups := uf.GetGroups()

	if len(groups) != 5 {
		t.Errorf("expected 5 groups when no unions performed, got %d", len(groups))
	}

	for _, group := range groups {
		if len(group) != 1 {
			t.Errorf("expected each group to have size 1, got %d", len(group))
		}
	}
}

func TestComplexScenario(t *testing.T) {
	uf := NewUnionFind(10)

	operations := []struct {
		x, y     int
		expected bool
	}{
		{0, 1, true},
		{2, 3, true},
		{4, 5, true},
		{6, 7, true},
		{8, 9, true},
		{0, 2, true},
		{4, 6, true},
		{1, 3, false},
		{5, 7, false},
		{0, 4, true},
		{8, 0, true},
		{9, 7, false},
	}

	for i, op := range operations {
		result := uf.Union(op.x, op.y)
		if result != op.expected {
			t.Errorf("operation %d: Union(%d, %d) = %v, expected %v", i, op.x, op.y, result, op.expected)
		}
	}

	groups := uf.GetGroups()
	if len(groups) != 1 {
		t.Errorf("expected all elements to be in 1 group, got %d groups", len(groups))
	}

	for i := 0; i < 10; i++ {
		root := uf.Find(i)
		if root != uf.Find(0) {
			t.Errorf("element %d has root %d, expected all to have same root as element 0 (%d)", i, root, uf.Find(0))
		}
	}
}
