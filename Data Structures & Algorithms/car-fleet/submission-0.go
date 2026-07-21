func carFleet(target int, position []int, speed []int) int {
	pairs := make([][2]int, len(position))
	for i, item := range position {
		pairs[i] = [2]int{item, speed[i]}
	}

	sort.Slice(pairs, func(i int, j int) bool {
		return pairs[i][0] > pairs[j][0]
	})

	stack := []float64{}

	for _, pair := range pairs {
		pos := pair[0]
		spe := pair[1]

		t := float64(target-pos) / float64(spe)

		if len(stack) == 0 || t > stack[len(stack)-1] {
			stack = append(stack, t)
		}
	}

	return len(stack)
}
