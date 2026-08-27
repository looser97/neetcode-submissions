func islandsAndTreasure(grid [][]int) {
	M := len(grid)
	N := len(grid[0])
    q := [][]int{}

	for i:=0; i < M; i++ {
		for j := 0; j < N; j++ {
			if (grid[i][j] == 0) {
				q = append(q, []int{i, j})
			}
		}
	}

	q = append(q, []int{-1, -1})
	level := 1
	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		if (top[0] == -1) {
			level++
			if len(q) > 0 {
				q = append(q, []int{-1, -1})
			}
		}
		for _, dir := range directions {
			row := top[0] + dir[0]
			col := top[1] + dir[1]

			if (row >=0 && col >= 0 && row < M && col < N) {
				if (grid[row][col] == 2147483647) {
					grid[row][col] = level
					q = append(q, []int{row, col})
				}
			}
		}
	}
}
