func characterReplacement(s string, k int) int {
	countMapping := map[byte]int{}
	maxLength := 0
	maxFrequescy := 0

	start, end := 0, 0

	for end < len(s) {
		countMapping[s[end]]++
		if countMapping[s[end]] > maxFrequescy {
			maxFrequescy = countMapping[s[end]]
		}

		for (end-start+1)-maxFrequescy > k {
			countMapping[s[start]]--
			start++
		}

		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
		end++
	}
	return maxLength
}
