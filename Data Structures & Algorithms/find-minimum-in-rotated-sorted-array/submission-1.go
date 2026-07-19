func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	ans := nums[0]
	for l <= r {
		if nums[l] < nums[r] {
			ans = min(ans, nums[l])
			break
		}
		mid := l + (r-l)/2
		ans = min(ans, nums[mid])
		if nums[mid] >= nums[l] {
			// left side sorted

			l = mid + 1

		} else {
			// right side sorted
			r = mid - 1
		}
	}
	return ans
}
