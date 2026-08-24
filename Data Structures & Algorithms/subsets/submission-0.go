func subsets(nums []int) [][]int {
	ans := [][]int{}

	var backTrack func (current []int, index int)

	backTrack = func (current []int, index int) {
		temp := append([]int{}, current...)
		ans = append(ans, temp)
		if (len(current) == len(nums)) {
			return
		}

		for i := index; i < len(nums); i++ {
			current = append(current, nums[i])
			backTrack(current, i + 1)
			current = current[:len(current) - 1]
		}
	}
	backTrack([]int{}, 0)
	return ans
}
