package graphs

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{parent, size}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX != rootY {
		if uf.size[rootX] < uf.size[rootY] {
			rootX, rootY = rootY, rootX
		}
		uf.parent[rootY] = rootX
		uf.size[rootX] += uf.size[rootY]
		return true
	}
	return false
}

func (uf *UnionFind) GetGroups() map[int][]int {
	groups := make(map[int][]int)
	for i := range uf.parent {
		root := uf.Find(i)
		groups[root] = append(groups[root], i)
	}
	return groups
}
