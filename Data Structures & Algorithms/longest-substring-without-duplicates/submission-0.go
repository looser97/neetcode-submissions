func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0

	mapping := map[byte]bool{}
	maxLength := 0

	for right < len(s) {
		if mapping[s[right]] {
			mapping[s[left]] = false
			left++
		} else {
			maxLength = max(maxLength, right-left+1)
			mapping[s[right]] = true
			right++
		}

	}
	return maxLength
}
