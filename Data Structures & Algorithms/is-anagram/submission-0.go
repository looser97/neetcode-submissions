import "slices"
func isAnagram(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	if sLen != tLen {
		return false
	}

	sSlice := make([]int, 26)
	tSlice := make([]int, 26)

	for i := 0; i < sLen; i++ {
		sSlice[s[i]-'a']++
		tSlice[t[i]-'a']++
	}

	return slices.Equal(sSlice, tSlice)
}
