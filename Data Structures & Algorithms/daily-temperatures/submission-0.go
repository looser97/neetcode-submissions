func dailyTemperatures(temperatures []int) []int {
	results := make([]int, len(temperatures))

	stack := []int{}

	for i, item := range temperatures {
		if len(stack) > 0 {
			for len(stack) > 0 && item > temperatures[stack[len(stack)-1]] {
				results[stack[len(stack)-1]] = i - stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
		} else {
			stack = append(stack, i)
		}
	}

	return results
}
