func maxAreaOfIsland(grid [][]int) int {
	maxi := 0
	globalMaxi := 0
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	instlands := 0

	var dfs func(i, j int)

	dfs = func(i, j int) {
		visited[i][j] = true
		maxi++
		if i > 0 && grid[i-1][j] == 1 && !visited[i-1][j] {
			dfs(i-1, j)
		}
		if i+1 < len(grid) && grid[i+1][j] == 1 && !visited[i+1][j] {
			dfs(i+1, j)
		}
		if j > 0 && grid[i][j-1] == 1 && !visited[i][j-1] {
			dfs(i, j-1)
		}
		if j+1 < len(grid[0]) && grid[i][j+1] == 1 && !visited[i][j+1] {
			dfs(i, j+1)
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				instlands++
				maxi = 0
				dfs(i, j)
				globalMaxi = max(globalMaxi, maxi)
			}
		}
	}

	return globalMaxi
}
