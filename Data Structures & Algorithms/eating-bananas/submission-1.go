import "slices"
func minEatingSpeed(piles []int, h int) int {
	maxi := slices.Max(piles)

	if h == len(piles) {
		return maxi
	}

	l := 0
	r := maxi

	minimumCuts := math.MaxFloat64

	for l <= r {
		mid := l + (r-l)/2
		if mid == 0 {
			break
		}
		hoursNeeded := 0.0
		for _, item := range piles {
			if mid > 0 {
				val := math.Ceil(float64(item) / float64(mid))
				hoursNeeded += val
			}
		}

		fmt.Println(l, r, mid, hoursNeeded)

		if int(hoursNeeded) <= h {
			minimumCuts = min(minimumCuts, float64(mid))
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return int(minimumCuts)
}
