func maxCoins(nums []int) int {
    // O(n^3)
    nums = append([]int{1}, nums...)
    nums = append(nums, 1)
    n := len(nums)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    for k := 2; k < n; k++ {
        for left := 0; left < n-k; left++ {
            right := left + k
            for i := left+1; i < right; i++ {
                dp[left][right] = max(dp[left][right], dp[left][i] + nums[left]*nums[i]*nums[right] + dp[i][right])
            }
        }
    }
    return dp[0][n-1]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}