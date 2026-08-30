type DSU struct {
	rank []int
	parent []int
}

func generateDSU(n int) DSU {
	dsu := DSU {
		rank: make([]int, n),
		parent: make([]int, n),
	}

	for i := 0; i<n; i++ {
		dsu.parent[i] = i
		dsu.rank[i] = i
	}

	return dsu
}

func (d *DSU) Find(node int) int {
	curr := node
	for curr != d.parent[curr] {
		d.parent[curr] = d.parent[d.parent[curr]]
		curr = d.parent[curr]
	}
	return curr
}

func (d *DSU) Union(u, v int) bool {
	pu := d.Find(u)
	pv := d.Find(v)

	if (pu == pv) {
		return false
	}

	if (d.rank[v] > d.rank[u]) {
		pu, pv = pv, pu
	}

	d.parent[pv] = pu
	d.rank[pu] += d.rank[pv]
	return true
}


func countComponents(n int, edges [][]int) int {
    res := n
	dsu := generateDSU(n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if (dsu.Union(u, v)) {
			res--
		}
	}
	return res
}
