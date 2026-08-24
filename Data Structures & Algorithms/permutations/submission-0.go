func permute(nums []int) [][]int {
	ans := [][]int{}

	var backTrack func(current []int)

	used := make([]bool, len(nums))

	backTrack = func(current []int) {
		if (len(current) == len(nums)) {
			temp := append([]int{}, current...)
			ans = append(ans, temp)
			return
		}

		for i := 0; i < len(nums) ; i++ {
			if (used[i]) {
				continue
			}
			used[i] = true
			current = append(current, nums[i])
			backTrack(current)
			current = current[:len(current) - 1]
			used[i] = false
			
		}
	}

	backTrack([]int{})

	return ans
}
