func isValid(s string) bool {
    stack := []rune{}

	for _, ch := range s {
		if ch == '{' || ch == '[' || ch == '(' {
			stack = append(stack, ch)
		} else {
			l := len(stack)
			if l == 0 {
				return false
			}
			if ch == '}' && stack[l-1] != '{' {
				return false
			}
			if ch == ']' && stack[l-1] != '[' {
				return false
			}
			if ch == ')' && stack[l-1] != '(' {
				return false
			}
			stack = stack[:l-1]
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
