func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left <= right {
		if !isAlphaNumeric(s[left]) {
			left++
		} else if !isAlphaNumeric(s[right]) {
			right--
		} else {
			if toLower(s[left]) != toLower(s[right]) {
				return false
			}
			left++
			right--
		}
	}

	return true
}
