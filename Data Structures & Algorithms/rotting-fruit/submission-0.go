func orangesRotting(grid [][]int) int {
	M := len(grid)
	N := len(grid[0])
    freshOranges := 0
	q := [][]int{}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if (grid[i][j] == 1) {
				freshOranges++
			}

			if (grid[i][j] == 2) {
				q = append(q, []int{i, j})
			}
		}
	}

	q = append(q, []int{-1, -1})

	timeUsed := -1
	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for len(q) > 0 {
		top := q[0]
		q = q[1:]

		if (top[0] == -1) {
			timeUsed++
			if len(q) > 0 {
				q = append(q, []int{-1, -1})
			}
		}

		for _, dir := range directions {
			row := top[0] + dir[0]
			col := top[1] + dir[1]
			if (row >=0 && col >= 0 && row <= M-1 && col <= N-1) {
				if (grid[row][col] == 1) {
					grid[row][col] = 2
					freshOranges--
					q = append(q, []int{row, col})
				}
			}
		}
	}

	if freshOranges == 0 {
		return timeUsed
	}

	return -1
}