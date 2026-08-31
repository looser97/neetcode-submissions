func rob(nums []int) int {
	if len(nums) == 1 {
        return nums[0]
    }
    return max(help(nums[:len(nums)-1]), help(nums[1:]))
}

func help(nums []int) int {
	n := len(nums)
    if n == 0 {
        return 0
    }
    if n == 1 {
        return nums[0]
    }

    dp := make([]int, n)
    dp[0] = nums[0]
    dp[1] = max(nums[0], nums[1])

    for i := 2; i < n; i++ {
        dp[i] = max(dp[i-1], nums[i] + dp[i-2])
    }

    return dp[n-1]
}
