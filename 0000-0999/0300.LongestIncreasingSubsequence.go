func lengthOfLIS(nums []int) int {
    // DP, O(n^2) + O(n)
    n := len(nums)
    dp := make([]int, n)
    ans := 0
    for i := 0; i < n; i++ {
        dp[i] = 1
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
        ans = max(ans, dp[i])
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}