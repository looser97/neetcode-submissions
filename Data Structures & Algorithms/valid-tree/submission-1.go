type DSU struct {
	parent []int
	rank []int
}

func NewDSU(n int) DSU {
	parent := make([]int, n)
	rank := make([]int, n)

	for i := 0 ; i<n; i++ {
		parent[i] = i
		rank[i] = i
	}

	return DSU{
		parent,
		rank,
	}
}

func (dsu *DSU) Find(node int) int {
	curr := node
	for curr != dsu.parent[curr] {
		dsu.parent[curr] = dsu.parent[dsu.parent[curr]]
		curr = dsu.parent[curr]
	}
	return curr
}

func (dsu *DSU) Union(u int, v int) bool {
	pu := dsu.Find(u)
	pv := dsu.Find(v)

	if (pu == pv) {
		return false
	}

	if (dsu.rank[pv] > dsu.rank[pu]) {
		pu, pv = pv, pu
	}

	dsu.parent[pv] = pu
	dsu.rank[pu] += dsu.rank[pv]
	return true
}


func validTree(n int, edges [][]int) bool {
    if len(edges) > n - 1 {
		return false;
	}

	dsu := NewDSU(n)

	totalComponents := n

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if (!dsu.Union(u, v)) {
			return false
		}
		totalComponents--
	}

	return totalComponents == 1
	 
}
