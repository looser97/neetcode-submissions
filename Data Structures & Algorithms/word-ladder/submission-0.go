func ladderLength(beginWord string, endWord string, wordList []string) int {
    graph := map[string][]string{}

	for _, word := range wordList {
		for i := 0; i < len(word); i++ {
			tempWord := word[:i] + "*" + word[i+1:]
			graph[tempWord] = append(graph[tempWord], word)
		} 
	}

	level := 0

	q := []string{beginWord}

	visited := map[string]bool{}

	for len(q) > 0 {
		level++
		n := len(q)
		for range n {
			top := q[0]
			q = q[1:]
			if (top == endWord) {
				return level
			}
			for i := 0; i < len(top); i++ {
				tempWord := top[:i] + "*" + top[i+1:]
				for _, w := range graph[tempWord] {
					if (!visited[w]) {
						visited[w] = true
						q = append(q, w)
					}
				}
			} 

		}
	}

	return 0

}
