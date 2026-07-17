import "slices"
func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	s1Slice := make([]int, 26)
	s2Slice := make([]int, 26)
	for _, ch := range s1 {
		s1Slice[ch-'a']++
	}

	left := 0
	right := 0

	for right < len(s1)-1 {
		s2Slice[s2[right]-'a']++
		right++
	}

	for right < len(s2) {
		s2Slice[s2[right]-'a']++
		if slices.Equal(s1Slice, s2Slice) {
			return true
		}
		s2Slice[s2[left]-'a']--
		left++
		right++
	}

	return false
}
