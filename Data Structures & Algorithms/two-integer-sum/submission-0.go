func twoSum(nums []int, target int) []int {
    visited := map[int]int{}

	for i, item := range nums {
		remainingSum := target - item
		found, ok := visited[remainingSum]
		if ok {
			return []int{found, i}
		}
		visited[item] = i
	}
	return []int{}
}
