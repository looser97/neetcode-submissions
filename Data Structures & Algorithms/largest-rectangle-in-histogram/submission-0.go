func largestRectangleArea(heights []int) int {
	stack := [][2]int{}
	maxArea := 0

	heights = append(heights, 0)

	for i, item := range heights {
		if len(stack) == 0 || stack[len(stack)-1][1] <= item {
			stack = append(stack, [2]int{i, item})
		} else {
			indexToUse := i
			for len(stack) > 0 && stack[len(stack)-1][1] > item {
				topElement := stack[len(stack)-1]
				area := topElement[1] * (i - topElement[0])
				maxArea = max(area, maxArea)
				indexToUse = topElement[0]
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, [2]int{indexToUse, item})
		}
	}
	return maxArea
}
