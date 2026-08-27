func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := map[int][]int{}
	inDegree := make([]int, numCourses)

	for _, item := range prerequisites {
		graph[item[1]] = append(graph[item[1]], item[0])
		inDegree[item[0]]++
	}

	q := []int{}

	for i := 0 ; i < numCourses; i++ {
		if (inDegree[i] == 0) {
			q = append(q, i)
		}
	}

	ans := []int{}

	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		ans = append(ans, top)

		for _, neigh := range graph[top] {
			inDegree[neigh]--
			if inDegree[neigh] == 0 {
				q = append(q, neigh)
			}
		}
	}

	return len(ans) == numCourses
}
