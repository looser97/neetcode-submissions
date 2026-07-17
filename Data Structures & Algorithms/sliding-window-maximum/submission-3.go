func maxSlidingWindow(nums []int, k int) []int {
    output := []int{}
	left, right := 0, 0
	q := []int{}

	for right < len(nums) {
		for len(q) > 0 && nums[q[len(q)-1]] < nums[right] {
			q = q[:len(q)-1]
		}
		q = append(q, right)

		if left > q[0] {
			q = q[1:]
		}

		if right+1 >= k {
			output = append(output, nums[q[0]])
			left++
		}
		right++
	}
	return output
}
