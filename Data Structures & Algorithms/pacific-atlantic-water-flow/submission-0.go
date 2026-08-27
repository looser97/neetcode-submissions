func pacificAtlantic(heights [][]int) [][]int {
    M := len(heights)
	N := len(heights[0])

	pc := map[[2]int]bool{}
	at := map[[2]int]bool{}

	var dfs func(i, j int, visited map[[2]int]bool, prevHeight int) 
	dfs = func(i, j int, visited map[[2]int]bool, prevHeight int) {
		coo := [2]int{i, j}

		if visited[coo] || i < 0 || j < 0 || i == M || j == N || heights[i][j] < prevHeight {
			return
		}

		visited[coo] = true
		dfs(i+1, j, visited, heights[i][j])
		dfs(i-1, j, visited, heights[i][j])
		dfs(i, j+1, visited, heights[i][j])
		dfs(i, j-1, visited, heights[i][j])
	}

	for r := 0 ; r < M; r++ {
		dfs(r, 0, pc, heights[r][0])
		dfs(r, N-1, at, heights[r][N-1])
	}

	for c := 0 ; c < N; c++ {
		dfs(0, c, pc, heights[0][c])
		dfs(M - 1, c, at, heights[M-1][c])
	}

	result := make([][]int, 0)
    for r := 0; r < M; r++ {
        for c := 0; c < N; c++ {
            coord := [2]int{r, c}
            if pc[coord] && at[coord] {
                result = append(result, []int{r, c})
            }
        }
    }

    return result
}
