func maxProfit(prices []int) int {
	left := 0
	right := 1

	maxi := 0

	for right < len(prices) {
		if prices[right] < prices[left] {
			left = right
		} else {
			if prices[right]-prices[left] > maxi {
				maxi = prices[right] - prices[left]
			}
		}
		right++
	}
	return maxi
}
