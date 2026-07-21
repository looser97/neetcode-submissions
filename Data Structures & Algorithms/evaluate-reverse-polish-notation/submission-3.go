func evalRPN(tokens []string) int {
	if len(tokens) == 1 {
		convert, _ := strconv.Atoi(tokens[0])
		return convert
	}
	val := 0
	stack := []int{}
	for _, item := range tokens {
		convert, error := strconv.Atoi(item)
		if error == nil {
			stack = append(stack, convert)
		} else {
			firstElement := stack[len(stack)-2]
			secondElement := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch item {
			case "+":
				val = firstElement + secondElement
			case "-":
				val = firstElement - secondElement
			case "*":
				val = firstElement * secondElement
			case "/":
				val = firstElement / secondElement
			}
			stack = append(stack, val)
		}
	}
	return val
}