func trap(height []int) int {
	left := 0
	leftMax := height[left]
	right := len(height) - 1
	rightMax := height[right]
	count := 0
	for left < right {
		if leftMax < rightMax {
			left++
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				count += leftMax - height[left]
			}
		} else {
			right--
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				count += rightMax - height[right]
			}
		}
	}

	return count
}
