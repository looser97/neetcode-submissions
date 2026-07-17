func maxArea(height []int) int {
	left := 0
	right := len(height) - 1

	area := 0

	for left < right {
		calculatedArea := (right - left) * int(math.Min(float64(height[left]), float64(height[right])))
		if calculatedArea > area {
			area = calculatedArea
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return area
}
