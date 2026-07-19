func searchMatrix(matrix [][]int, target int) bool {
	noOfRows := len(matrix)
	noOfColums := len(matrix[0])

	t := 0
	b := noOfRows - 1

	selectedRow := -1

	for t <= b {
		mid := t + (b-t)/2
		if matrix[mid][0] <= target && matrix[mid][noOfColums-1] >= target {
			selectedRow = mid
			break
		}

		if matrix[mid][0] > target {
			b = mid - 1
		} else {
			t = mid + 1
		}
	}

	if selectedRow == -1 {
		return false
	}

	l := 0
	r := noOfColums - 1

	for l <= r {
		mid := l + (r-l)/2

		if matrix[selectedRow][mid] == target {
			return true
		}

		if matrix[selectedRow][mid] > target {
			r = mid - 1
		} else {
			l = l + 1
		}
	}

	return false
}
